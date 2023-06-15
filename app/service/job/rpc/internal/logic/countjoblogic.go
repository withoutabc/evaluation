package logic

import (
	"context"
	"rpc/app/common/consts/errs"
	"rpc/app/common/consts/maps"
	"rpc/utils/joinstring"

	"rpc/app/service/job/rpc/internal/svc"
	"rpc/app/service/job/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

const (
	prefix = "/evaluation/test/"
)

type CountJobLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewCountJobLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CountJobLogic {
	return &CountJobLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *CountJobLogic) CountJob(in *pb.CountJobReq) (*pb.CountJobResp, error) {
	//生成job名称
	str := joinstring.Join(in.Year, in.Month, in.Set, in.Level)
	jobName := joinstring.JoinJobName(str, "wordcount")
	//input path
	newFileName := joinstring.JoinOrigin(str) + ".txt"
	inputPath := prefix + string(in.Year) + "/" + maps.LevelMap[in.Level] + "/" + str + "/" + newFileName
	//output path
	outputPath := prefix + string(in.Year) + "/" + maps.LevelMap[in.Level] + "/" + str + "/" + jobName
	code := l.svcCtx.JobModel.WordCount(jobName, inputPath, outputPath)
	return &pb.CountJobResp{
		StatusCode: code,
		StatusMsg:  errs.ErrorsMap[code].Error(),
		JobName:    jobName,
	}, nil
}
