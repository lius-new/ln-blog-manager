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

type GetArticlesByPageWithBackendLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetArticlesByPageWithBackendLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetArticlesByPageWithBackendLogic {
	return &GetArticlesByPageWithBackendLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetArticlesByPageWithBackendLogic) GetArticlesByPageWithBackend(req *types.GetArticlesByPageWithBackendRequest) (resp *types.GetArticlesByPageWithBackendResponse, err error) {
	defer func() {
		if catchErr := recover(); catchErr != nil {
			var catchErr = catchErr.(error)
			switch {
			case strings.Contains(catchErr.Error(), api.ErrRequestParam.Error()):
				err = errors.New(api.ErrRequestParam.Error())
			}
		} else if err != nil {
			err = errors.New(strings.Replace(err.Error(), "rpc error: code = Unknown desc = ", "", 1))
		}
	}()

	// 调用rpc获取结果
	articlesResp, err := l.svcCtx.Content.SelectArtilceByPage(l.ctx, &content.SelectArticleByPageRequest{
		PageNum:  req.PageNum,
		PageSize: req.PageSize,
		HideShow: true,
	})
	if err != nil {
		panic(err)
	}

	// 封装数据
	forLen := len(articlesResp.Articles)
	data := make([]types.Data, forLen)
	for i := 0; i < forLen; i++ {
		current := articlesResp.Articles[i]
		data[i] = types.Data{
			Id:          current.Id,
			Title:       current.Title,
			Description: current.Desc,
			Tags:        current.Tags,
			Covers:      current.Covers,
			Visiable:    current.Visiable,
			UpdateAt:    current.Time,
		}
	}

	return &types.GetArticlesByPageWithBackendResponse{
		Data:  data,
		Total: articlesResp.Total,
	}, nil
}
