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

type SearchArticleLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewSearchArticleLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SearchArticleLogic {
	return &SearchArticleLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *SearchArticleLogic) SearchArticle(req *types.SearchArticleRequest) (resp *types.SearchArticleResponse, err error) {

	defer func() {
		if catchErr := recover(); catchErr != nil {
			var catchErr = catchErr.(error)
			switch {
			case strings.Contains(catchErr.Error(), api.ErrNotFound.Error()):
				err = errors.New(api.ErrInvalidNotFound.Error())
			}
		} else if err != nil {
			err = errors.New(strings.Replace(err.Error(), "rpc error: code = Unknown desc = ", "", 1))
		}
	}()
	searchResp, err := l.svcCtx.Content.SearchArtilce(l.ctx, &content.SearchArtilceRequest{
		Search: req.Search,
	})

	if err != nil {
		panic(err)
	}

	forLen := len(searchResp.Articles)
	data := make([]types.Article, forLen)
	for i := 0; i < forLen; i++ {
		current := searchResp.Articles[i]
		data[i] = types.Article{
			Id:          current.Id,
			Title:       current.Title,
			Description: current.Desc,
		}
	}

	return &types.SearchArticleResponse{
		Data:  data,
		Total: int64(forLen),
	}, nil
}
