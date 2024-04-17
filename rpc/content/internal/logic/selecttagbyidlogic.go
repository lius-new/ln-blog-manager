package logic

import (
	"context"

	"github.com/lius-new/blog-backend/rpc/content/content"
	"github.com/lius-new/blog-backend/rpc/content/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

type SelectTagByIdLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewSelectTagByIdLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SelectTagByIdLogic {
	return &SelectTagByIdLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 根据id获取tag
func (l *SelectTagByIdLogic) SelectTagById(in *content.SelectTagByIdRequest) (*content.SelectTagByIdResponse, error) {
	// todo: add your logic here and delete this line

	return &content.SelectTagByIdResponse{}, nil
}
