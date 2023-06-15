package logic

import (
	"context"
	"fmt"
	"io"
	"log"
	"net/http"
	"path/filepath"
	"rpc/app/common/consts/errs"
	"rpc/app/common/consts/maps"
	"rpc/app/service/file/api/internal/svc"
	"rpc/app/service/file/api/internal/types"
	"rpc/app/service/file/rpc/pb"
	"rpc/utils/joinstring"
	"strconv"

	"github.com/zeromicro/go-zero/core/logx"
)

const maxFileSize = 1024 * (1 << 20) // 1GB

type UploadLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUploadLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UploadLogic {
	return &UploadLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UploadLogic) Upload(req *types.UploadReq, r *http.Request) (resp *types.UploadResp, err error) {
	err = r.ParseMultipartForm(maxFileSize)
	if err != nil {
		log.Printf("max file siza err:%v\n", err)
		return &types.UploadResp{
			Status: errs.FileTooBig,
			Msg:    errs.ErrorsMap[errs.FileTooBig].Error(),
		}, nil
	}
	//获取文件
	file, handler, err := r.FormFile("file")
	if err != nil {
		fmt.Println(err)
		return &types.UploadResp{
			Status: errs.FileGetWrong,
			Msg:    errs.ErrorsMap[errs.FileGetWrong].Error(),
		}, nil
	}
	defer file.Close() //关闭文件
	//判断是否为46级,1-3套，年份4位数<2024，月份1-12
	if !((req.Level == 4 || req.Level == 6) && (req.Month <= 12 && req.Month >= 1) && (req.Year < 2024 && len(strconv.Itoa(int(req.Year))) == 4)) {
		return &types.UploadResp{
			Status: errs.FileNameWrong,
			Msg:    errs.ErrorsMap[errs.FileNameWrong].Error(),
		}, nil
	}

	// 重命名文件
	str := joinstring.Join(req.Year, req.Month, req.Set, req.Level)
	newFileName := joinstring.JoinOrigin(str) + filepath.Ext(handler.Filename)
	// 创建一个目录，如果目录已存在则无操作
	mkdirPath := prefix + string(req.Year) + "/" + maps.LevelMap[req.Level] + "/" + str
	err = l.svcCtx.HdfsCli.MkdirAll(mkdirPath, 0755)
	if err != nil {
		log.Printf("mkdirall err: %v", err)
		return &types.UploadResp{
			Status: errs.FileSystemInternal,
			Msg:    errs.ErrorsMap[errs.FileSystemInternal].Error(),
		}, nil
	}
	// 传输路径
	hdfsPath := mkdirPath + "/" + newFileName
	hdfsFile, err := l.svcCtx.HdfsCli.Create(hdfsPath)
	if err != nil {
		log.Printf("create err: %v", err)
		return &types.UploadResp{
			Status: errs.FileSystemInternal,
			Msg:    errs.ErrorsMap[errs.FileSystemInternal].Error(),
		}, nil
	}
	defer hdfsFile.Close()
	//复制到hdfs系统上
	_, err = io.Copy(hdfsFile, file)
	if err != nil {
		log.Printf("Failed to copy file to HDFS: %v", err)
		return &types.UploadResp{
			Status: errs.FileSystemInternal,
			Msg:    errs.ErrorsMap[errs.FileSystemInternal].Error(),
		}, nil
	}
	defer file.Close()
	//如果不是pdf文件就直接结束
	if filepath.Ext(handler.Filename) != ".pdf" {
		return &types.UploadResp{
			Status: errs.No,
			Msg:    errs.ErrorsMap[errs.No].Error(),
		}, nil
	}
	//文件信息写入数据库
	//检查文件是否已存在,不存在就写入，因为前面上传也能拦截，所以差不多
	response, err := l.svcCtx.FileRpc.Upload(l.ctx, &pb.UploadReq{
		Level: req.Level,
		Year:  req.Year,
		Month: req.Month,
		Set:   req.Set,
	})
	return &types.UploadResp{
		Status: response.StatusCode,
		Msg:    response.StatusMsg,
	}, nil
}
