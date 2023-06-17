package logic

import (
	"bufio"
	"context"
	"fmt"
	"log"
	"rpc/app/common/consts/errs"
	"rpc/app/common/consts/maps"
	"strconv"
	"strings"

	"rpc/app/service/job/rpc/internal/svc"
	"rpc/app/service/job/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

const (
	outputPathResultPrefix = "/evaluation/collect/result/"
)

type ViewWordLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewViewWordLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ViewWordLogic {
	return &ViewWordLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *ViewWordLogic) ViewWord(in *pb.ViewWordReq) (*pb.ViewWordResp, error) {
	//创建文件夹
	err := l.svcCtx.HdfsCli.MkdirAll("/evaluation/collect/result", 0755)
	if err != nil {
		return &pb.ViewWordResp{
			StatusCode: errs.FileSystemInternal,
			StatusMsg:  errs.ErrorsMap[errs.FileSystemInternal].Error(),
			Words:      nil,
		}, nil
	}
	//拼名字
	jobName := fmt.Sprintf("%s-%s-%s*%s", strconv.Itoa(int(in.Begin)), strconv.Itoa(int(in.End)), maps.LevelMap[in.Level], maps.WordlistMap[in.Except])
	fileName := fmt.Sprintf("%s.txt", jobName)
	//检查是否有这个文件
	fileInfoList, err := l.svcCtx.HdfsCli.ReadDir(outputPathResultPrefix)
	if err != nil {
		return &pb.ViewWordResp{
			StatusCode: errs.FileSystemInternal,
			StatusMsg:  errs.ErrorsMap[errs.FileSystemInternal].Error(),
			Words:      nil,
		}, nil
	}
	var count = 0
	for _, fileInfo := range fileInfoList {
		if fileInfo.Name() == jobName+".txt" {
			count++
		}
	}
	if count == 0 {
		log.Println(count)
		return &pb.ViewWordResp{
			StatusCode: errs.FileNotExist,
			StatusMsg:  errs.ErrorsMap[errs.FileNotExist].Error(),
			Words:      nil,
		}, nil
	}
	file, err := l.svcCtx.HdfsCli.Open(outputPathResultPrefix + fileName)
	log.Println(outputPathResultPrefix + fileName)
	if err != nil {
		return &pb.ViewWordResp{
			StatusCode: errs.FileSystemInternal,
			StatusMsg:  errs.ErrorsMap[errs.FileSystemInternal].Error(),
			Words:      nil,
		}, nil
	}
	var words []*pb.Words
	var x int32 = 0
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		parts := strings.Split(line, "\t") //将当前行按照制表符分割成两部分
		if len(parts) != 2 {               //如果分割后的长度不是2，说明格式不正确，跳过此行
			continue
		}
		word := parts[0] //获取单词
		var number int64
		number, err = strconv.ParseInt(parts[1], 10, 64) //将单词出现的次数转化为整型
		if err != nil {                                  //如果转化出错，说明格式不正确，跳过此行
			continue
		}
		words = append(words, &pb.Words{Word: word, Count: number})
		x++
		if x >= in.Count {
			break
		}
	}
	return &pb.ViewWordResp{
		StatusCode: errs.No,
		StatusMsg:  errs.ErrorsMap[errs.No].Error(),
		Words:      words,
	}, nil
}
