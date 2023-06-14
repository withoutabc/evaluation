// Code generated by goctl. DO NOT EDIT.
// Source: file.proto

package file

import (
	"context"

	"rpc/app/service/file/rpc/pb"

	"github.com/zeromicro/go-zero/zrpc"
	"google.golang.org/grpc"
)

type (
	DownloadReq  = pb.DownloadReq
	DownloadResp = pb.DownloadResp
	File         = pb.File
	GetFileReq   = pb.GetFileReq
	GetFileResp  = pb.GetFileResp
	UploadReq    = pb.UploadReq
	UploadResp   = pb.UploadResp

	FileZrpcClient interface {
		Upload(ctx context.Context, in *UploadReq, opts ...grpc.CallOption) (*UploadResp, error)
		Download(ctx context.Context, in *DownloadReq, opts ...grpc.CallOption) (*DownloadResp, error)
		GetFiles(ctx context.Context, in *GetFileReq, opts ...grpc.CallOption) (*GetFileResp, error)
	}

	defaultFileZrpcClient struct {
		cli zrpc.Client
	}
)

func NewFileZrpcClient(cli zrpc.Client) FileZrpcClient {
	return &defaultFileZrpcClient{
		cli: cli,
	}
}

func (m *defaultFileZrpcClient) Upload(ctx context.Context, in *UploadReq, opts ...grpc.CallOption) (*UploadResp, error) {
	client := pb.NewFileClient(m.cli.Conn())
	return client.Upload(ctx, in, opts...)
}

func (m *defaultFileZrpcClient) Download(ctx context.Context, in *DownloadReq, opts ...grpc.CallOption) (*DownloadResp, error) {
	client := pb.NewFileClient(m.cli.Conn())
	return client.Download(ctx, in, opts...)
}

func (m *defaultFileZrpcClient) GetFiles(ctx context.Context, in *GetFileReq, opts ...grpc.CallOption) (*GetFileResp, error) {
	client := pb.NewFileClient(m.cli.Conn())
	return client.GetFiles(ctx, in, opts...)
}