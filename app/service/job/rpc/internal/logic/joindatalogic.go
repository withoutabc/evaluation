package logic

import (
	"context"
	"fmt"
	"github.com/colinmarc/hdfs/v2"
	"io"
	"log"
	"rpc/app/common/consts/errs"
	"rpc/app/common/consts/maps"
	"rpc/utils/joinstring"
	"strconv"

	"rpc/app/service/job/rpc/internal/svc"
	"rpc/app/service/job/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type JoinDataLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewJoinDataLogic(ctx context.Context, svcCtx *svc.ServiceContext) *JoinDataLogic {
	return &JoinDataLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *JoinDataLogic) JoinData(in *pb.JoinDataReq) (*pb.JoinDataResp, error) {
	//找到输出的文件夹
	str := joinstring.Join(in.Year, in.Month, in.Set, in.Level)
	jobName := joinstring.JoinJobName(str, "wordcount")
	outputPath := prefix + strconv.Itoa(int(in.Year)) + "/" + maps.LevelMap[in.Level] + "/" + str + "/" + jobName

	// 读取输出目录中的所有文件
	fileInfos, err := l.svcCtx.HdfsCli.ReadDir(outputPath)
	if err != nil {
		log.Fatal(err)
	}

	//找到聚合后的位置
	resultFilename := joinstring.JoinResult(str)
	targetPath := prefix + strconv.Itoa(int(in.Year)) + "/" + maps.LevelMap[in.Level] + "/" + str + "/" + resultFilename //聚合后的目标路径
	//创建文件
	targetFile, err := l.svcCtx.HdfsCli.Create(targetPath)
	if err != nil {
		log.Println(err)
		return &pb.JoinDataResp{
			StatusCode: errs.FileIsExist,
			StatusMsg:  errs.ErrorsMap[errs.FileIsExist].Error(),
		}, nil
	}
	defer targetFile.Close()
	//遍历文件信息
	for _, fileInfo := range fileInfos {
		if fileInfo.IsDir() {
			continue
		}
		//找到每一个文件
		inFilePath := outputPath + "/" + fileInfo.Name()
		//打开
		var inFile *hdfs.FileReader
		inFile, err = l.svcCtx.HdfsCli.Open(inFilePath)
		if err != nil {
			log.Println(err)
			return &pb.JoinDataResp{
				StatusCode: errs.FileSystemInternal,
				StatusMsg:  errs.ErrorsMap[errs.FileSystemInternal].Error(),
			}, nil
		}
		// 将文件内容写入目标文件
		_, err = io.Copy(targetFile, inFile)
		if err != nil {
			log.Println(err)
			return &pb.JoinDataResp{
				StatusCode: errs.FileSystemInternal,
				StatusMsg:  errs.ErrorsMap[errs.FileSystemInternal].Error(),
			}, nil
		}
		fmt.Printf("文件 '%s' 被复制到输出文件中\n", inFilePath)
		inFile.Close() //关闭
	}
	return &pb.JoinDataResp{
		StatusCode: errs.No,
		StatusMsg:  errs.ErrorsMap[errs.No].Error(),
	}, nil
}
