package logic

import (
	"context"
	"rpc/app/common/consts/errs"
	"rpc/app/service/job/rpc/pb"
	"strconv"

	"rpc/app/service/job/api/internal/svc"
	"rpc/app/service/job/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type CountJobLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewCountJobLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CountJobLogic {
	return &CountJobLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *CountJobLogic) CountJob(req *types.CountJobReq) (resp *types.CountJobResp, err error) {
	//判断是否为46级,1-3套，年份4位数<2024，月份1-12
	if !((req.Level == 4 || req.Level == 6) && (req.Month <= 12 && req.Month >= 1) && (req.Year < 2024 && len(strconv.Itoa(int(req.Year))) == 4)) {
		return &types.CountJobResp{
			Status: errs.FileNameWrong,
			Msg:    errs.ErrorsMap[errs.FileNameWrong].Error(),
		}, nil
	}
	response, err := l.svcCtx.JobRpc.CountJob(l.ctx, &pb.CountJobReq{
		Level: req.Level,
		Year:  req.Year,
		Month: req.Month,
		Set:   req.Set,
	})
	return &types.CountJobResp{
		Status:  response.StatusCode,
		Msg:     response.StatusMsg,
		JobName: response.JobName,
	}, nil
}
