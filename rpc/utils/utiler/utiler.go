// Code generated by goctl. DO NOT EDIT.
// Source: utils.proto

package utiler

import (
	"context"

	"github.com/lius-new/blog-backend/rpc/utils/utils"

	"github.com/zeromicro/go-zero/zrpc"
	"google.golang.org/grpc"
)

type (
	MD5Reponse = utils.MD5Reponse
	MD5Reqeust = utils.MD5Reqeust

	Utiler interface {
		MD5(ctx context.Context, in *MD5Reqeust, opts ...grpc.CallOption) (*MD5Reponse, error)
	}

	defaultUtiler struct {
		cli zrpc.Client
	}
)

func NewUtiler(cli zrpc.Client) Utiler {
	return &defaultUtiler{
		cli: cli,
	}
}

func (m *defaultUtiler) MD5(ctx context.Context, in *MD5Reqeust, opts ...grpc.CallOption) (*MD5Reponse, error) {
	client := utils.NewUtilerClient(m.cli.Conn())
	return client.MD5(ctx, in, opts...)
}
