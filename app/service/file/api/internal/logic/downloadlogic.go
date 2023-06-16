package logic

import (
	"bufio"
	"context"
	"rpc/app/common/consts/errs"
	"rpc/app/common/consts/maps"
	"rpc/app/service/file/rpc/pb"
	"rpc/utils/joinstring"
	"strconv"

	"rpc/app/service/file/api/internal/svc"
	"rpc/app/service/file/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

const (
	prefix = "/evaluation/test/"
)

type DownloadLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewDownloadLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DownloadLogic {
	return &DownloadLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *DownloadLogic) Download(req *types.DownloadReq) (resp *types.DownloadResp, body []byte, filename string, err error) {
	//判读是否存在要下载的内容
	response, err := l.svcCtx.FileRpc.Download(l.ctx, &pb.DownloadReq{
		Type: req.Type,
		Id:   req.Id,
	})
	//判断response
	if response.StatusCode != errs.No {
		return &types.DownloadResp{
			Status: response.StatusCode,
			Msg:    response.StatusMsg,
		}, nil, "", nil
	}
	//没有问题的话就去下载文件
	str := joinstring.Join(req.Year, req.Month, req.Set, req.Level)
	newFileName := joinstring.JoinOrigin(str) + ".pdf"
	mkdirPath := prefix + strconv.Itoa(int(req.Year)) + "/" + maps.LevelMap[req.Level] + "/" + str
	file, err := l.svcCtx.HdfsCli.Open(mkdirPath + "/" + newFileName)
	reader := bufio.NewReader(file)
	body = make([]byte, maxFileSize)
	//读取文件
	_, err = reader.Read(body)
	if err != nil {
		return &types.DownloadResp{
			Status: errs.InternalServer,
			Msg:    errs.ErrorsMap[errs.InternalServer].Error(),
		}, nil, "", nil
	}
	defer file.Close()
	return &types.DownloadResp{
		Status: errs.No,
		Msg:    errs.ErrorsMap[errs.No].Error(),
	}, body, newFileName, nil
}
