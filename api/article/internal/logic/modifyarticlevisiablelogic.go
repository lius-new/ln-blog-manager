package logic

import (
	"context"
	"errors"
	"strings"

	"github.com/lius-new/blog-backend/api"
	"github.com/lius-new/blog-backend/api/article/internal/svc"
	"github.com/lius-new/blog-backend/api/article/internal/types"
	"github.com/lius-new/blog-backend/rpc/content/content"

	"github.com/zeromicro/go-zero/core/logx"
)

type ModifyArticleVisiableLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewModifyArticleVisiableLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ModifyArticleVisiableLogic {
	return &ModifyArticleVisiableLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *ModifyArticleVisiableLogic) ModifyArticleVisiable(req *types.ModifyArticleVisiableRequest) (resp *types.ModifyArticleVisiableResponse, err error) {
	// defer recover错误信息
	defer func() {
		if catchErr := recover(); catchErr != nil {
			var catchErr = catchErr.(error)
			switch {
			case strings.Contains(catchErr.Error(), api.ErrRequestParam.Error()):
				err = errors.New(api.ErrRequestParam.Error())
			case strings.Contains(catchErr.Error(), api.ErrNotFound.Error()):
				err = errors.New(api.ErrInvalidNotFound.Error())
			}
		} else if err != nil {
			err = errors.New(strings.Replace(err.Error(), "rpc error: code = Unknown desc = ", "", 1))
		}
	}()

	_, err = l.svcCtx.Content.ModifyArtilceVisiable(l.ctx, &content.ModifyArticleVisiableRequest{
		Id:       req.Id,
		Visiable: req.Visiable,
	})
	if err != nil {
		panic(err)
	}

	return
}
