package logic

import (
	"context"
	"fmt"
	"log"
	"os"
	"rpc/app/common/consts/errs"
	"rpc/app/common/consts/maps"
	"strconv"
	"strings"

	"rpc/app/service/job/rpc/internal/svc"
	"rpc/app/service/job/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

const (
	outputPathPrefix = "/evaluation/collect/"
	exceptPathPrefix = "/evaluation/wordlist/"
)

type CollectJobLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewCollectJobLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CollectJobLogic {
	return &CollectJobLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *CollectJobLogic) CollectJob(in *pb.CollectJobReq) (*pb.CollectJobResp, error) {
	jobName := fmt.Sprintf("%s-%s-%s/%s", strconv.Itoa(int(in.Begin)), strconv.Itoa(int(in.End)), maps.LevelMap[in.Level], maps.WordlistMap[in.Except])
	outputPath := outputPathPrefix + jobName
	exceptPath := exceptPathPrefix + maps.WordlistMap[in.Except]
	var inputFilePaths []string
	fileInfoList, err := l.svcCtx.HdfsCli.ReadDir(prefix)
	if err != nil {
		return &pb.CollectJobResp{
			StatusCode: errs.FileSystemInternal,
			StatusMsg:  errs.ErrorsMap[errs.FileSystemInternal].Error(),
			JobName:    jobName,
		}, nil
	}
	if len(fileInfoList) == 0 {
		return &pb.CollectJobResp{
			StatusCode: errs.NoRecord,
			StatusMsg:  errs.ErrorsMap[errs.NoRecord].Error(),
			JobName:    jobName,
		}, nil
	}
	//找到所有存在的文件
	for i := in.Begin; i <= in.End; i++ {
		for _, fileInfo := range fileInfoList {
			if fileInfo.Name() == strconv.Itoa(int(i)) {
				//如果有，就继续往下找文件
				var fileInfoList1 []os.FileInfo
				//读取对应等级的文件列表
				fileInfoList1, err = l.svcCtx.HdfsCli.ReadDir(prefix + fileInfo.Name() + "/" + maps.LevelMap[in.Level])
				if err != nil {
					return &pb.CollectJobResp{
						StatusCode: errs.FileSystemInternal,
						StatusMsg:  errs.ErrorsMap[errs.FileSystemInternal].Error(),
						JobName:    jobName,
					}, nil
				}
				if len(fileInfoList1) == 0 {
					continue
				}
				//遍历每个文件夹
				for _, fileInfo1 := range fileInfoList1 {
					//看这个文件夹里的所有文件里有没有result.txt结尾的文件

					var fileInfoList2 []os.FileInfo
					fileInfoList2, err = l.svcCtx.HdfsCli.ReadDir(prefix + fileInfo.Name() + "/" + maps.LevelMap[in.Level] + "/" + fileInfo1.Name())
					if err != nil {
						return &pb.CollectJobResp{
							StatusCode: errs.FileSystemInternal,
							StatusMsg:  errs.ErrorsMap[errs.FileSystemInternal].Error(),
							JobName:    jobName,
						}, nil
					}
					if len(fileInfoList2) == 0 {
						continue
					}
					//遍历这个文件夹下的所有文件
					for _, fileInfo2 := range fileInfoList2 {
						//有就组合一下路径
						if strings.HasSuffix(fileInfo2.Name(), "result.pdf") {
							inputFilePaths = append(inputFilePaths, prefix+fileInfo.Name()+"/"+maps.LevelMap[in.Level]+"/"+fileInfo1.Name()+"/"+fileInfo2.Name())
						} //没有就继续循环
					}

				}
			}
		}
	}
	log.Println(inputFilePaths)
	code := l.svcCtx.JobModel.Collect(jobName, outputPath, inputFilePaths, exceptPath)
	return &pb.CollectJobResp{
		StatusCode: code,
		StatusMsg:  errs.ErrorsMap[code].Error(),
		JobName:    jobName,
	}, nil
}
