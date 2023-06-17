package logic

import (
	"context"
	"fmt"
	"github.com/colinmarc/hdfs/v2"
	"io"
	"log"
	"rpc/app/common/consts/errs"
	"rpc/app/common/consts/maps"
	"strconv"

	"rpc/app/service/job/rpc/internal/svc"
	"rpc/app/service/job/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type JoinData2Logic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewJoinData2Logic(ctx context.Context, svcCtx *svc.ServiceContext) *JoinData2Logic {
	return &JoinData2Logic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *JoinData2Logic) JoinData2(in *pb.JoinDataReq2) (*pb.JoinDataResp2, error) {
	//找到输出文件夹
	jobName := fmt.Sprintf("%s-%s-%s*%s", strconv.Itoa(int(in.Begin)), strconv.Itoa(int(in.End)), maps.LevelMap[in.Level], maps.WordlistMap[in.Except])
	outputPath := outputPathPrefix + jobName

	// 读取输出目录中的所有文件
	fileInfos, err := l.svcCtx.HdfsCli.ReadDir(outputPath)
	if err != nil {
		log.Fatal(err)
	}
	//找到聚合后的位置
	targetPath := outputPathPrefix + "result/" + jobName + ".txt"
	//创建文件
	targetFile, err := l.svcCtx.HdfsCli.Create(targetPath)
	if err != nil {
		log.Println(err)
		return &pb.JoinDataResp2{
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
			return &pb.JoinDataResp2{
				StatusCode: errs.FileSystemInternal,
				StatusMsg:  errs.ErrorsMap[errs.FileSystemInternal].Error(),
			}, nil
		}
		// 将文件内容写入目标文件
		_, err = io.Copy(targetFile, inFile)
		if err != nil {
			log.Println(err)
			return &pb.JoinDataResp2{
				StatusCode: errs.FileSystemInternal,
				StatusMsg:  errs.ErrorsMap[errs.FileSystemInternal].Error(),
			}, nil
		}
		fmt.Printf("文件 '%s' 被复制到输出文件中\n", inFilePath)
		inFile.Close() //关闭
	}
	return &pb.JoinDataResp2{
		StatusCode: errs.No,
		StatusMsg:  errs.ErrorsMap[errs.No].Error(),
	}, nil

	return &pb.JoinDataResp2{}, nil
}
