package logic

import (
	"context"
	"log"
	"rpc/app/common/consts/errs"
	"rpc/app/service/file/rpc/internal/svc"
	"rpc/app/service/file/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetFilesLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetFilesLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetFilesLogic {
	return &GetFilesLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetFilesLogic) GetFiles(in *pb.GetFileReq) (*pb.GetFileResp, error) {
	log.Println("2")
	files, code := l.svcCtx.FileModel.GetFileList(l.ctx, in.Level, in.Year, in.Month, in.Set)
	var data []*pb.File
	if files == nil {
		data = nil
	} else {
		for _, file := range files {
			pbFile := pb.File{
				Id:         file.ID,
				Set:        file.Set,
				Month:      file.Month,
				Year:       file.Year,
				Level:      file.Level,
				UpdateTime: file.UpdateTime.Format("2006-01-02 15:04:05"),
			}
			data = append(data, &pbFile)
		}
	}
	log.Println(data)
	return &pb.GetFileResp{
		StatusCode: code,
		StatusMsg:  errs.ErrorsMap[code].Error(),
		Data:       data,
	}, nil
}
