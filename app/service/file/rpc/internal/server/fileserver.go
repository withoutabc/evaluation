// Code generated by goctl. DO NOT EDIT.
// Source: file.proto

package server

import (
	"context"

	"rpc/app/service/file/rpc/internal/logic"
	"rpc/app/service/file/rpc/internal/svc"
	"rpc/app/service/file/rpc/pb"
)

type FileServer struct {
	svcCtx *svc.ServiceContext
	pb.UnimplementedFileServer
}

func NewFileServer(svcCtx *svc.ServiceContext) *FileServer {
	return &FileServer{
		svcCtx: svcCtx,
	}
}

func (s *FileServer) Upload(ctx context.Context, in *pb.UploadReq) (*pb.UploadResp, error) {
	l := logic.NewUploadLogic(ctx, s.svcCtx)
	return l.Upload(in)
}

func (s *FileServer) Download(ctx context.Context, in *pb.DownloadReq) (*pb.DownloadResp, error) {
	l := logic.NewDownloadLogic(ctx, s.svcCtx)
	return l.Download(in)
}

func (s *FileServer) GetFiles(ctx context.Context, in *pb.GetFileReq) (*pb.GetFileResp, error) {
	l := logic.NewGetFilesLogic(ctx, s.svcCtx)
	return l.GetFiles(in)
}

func (s *FileServer) RemoveFile(ctx context.Context, in *pb.RemoveFileReq) (*pb.RemoveFileResp, error) {
	l := logic.NewRemoveFileLogic(ctx, s.svcCtx)
	return l.RemoveFile(in)
}
