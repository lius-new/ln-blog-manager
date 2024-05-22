package logic

import (
	"context"

	"github.com/zeromicro/go-zero/core/logx"

	"github.com/lius-new/blog-backend/rpc"
	"github.com/lius-new/blog-backend/rpc/content/content"
	"github.com/lius-new/blog-backend/rpc/content/internal/svc"
)

type SelectArtilceByTagLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewSelectArtilceByTagLogic(
	ctx context.Context,
	svcCtx *svc.ServiceContext,
) *SelectArtilceByTagLogic {
	return &SelectArtilceByTagLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 根据tag name获取文章
func (l *SelectArtilceByTagLogic) SelectArtilceByTag(
	in *content.SelectArticleByTagRequest,
) (*content.SelectArticleByTagResponse, error) {
	// 判断分页要求的参数
	if in.PageSize == 0 || in.PageNum == 0 {
		return nil, rpc.ErrRequestParam
	}

	tag, err := l.svcCtx.ModelWithTag.FindByName(l.ctx, in.Tag)
	if err != nil {
		return nil, err
	}

	total := int64(len(tag.Articles))

	// 如: total= 99, 但是pageSize * PageNum = 100的情况就会超出slice范围导致异常.
	end := in.PageSize * in.PageNum
	if total < end {
		end = total
	}

	articleIds := tag.Articles[in.PageSize*(in.PageNum-1) : end]
	findArticles, err := l.svcCtx.ModelWithArticle.FindByIds(l.ctx, articleIds)
	if err != nil {
		return nil, err
	}

	forLen := len(findArticles)
	// 封装查询结果为[]*content.SelectArticles类型
	respArticles := make([]*content.SelectArticles, forLen)

	for i := 0; i < forLen; i++ {
		currentArticle := findArticles[i]

		// tags属性中原本包含的是tagid, 修改article中的tagid为tagname, 即根据id查询tag再获取tagName
		selectTagByIdLogic := NewSelectTagByIdLogic(l.ctx, l.svcCtx)
		for i := 0; i < len(currentArticle.Tags); i++ {
			tag, _ := selectTagByIdLogic.SelectTagById(&content.SelectTagByIdRequest{
				Id: currentArticle.Tags[i],
			})

			currentArticle.Tags[i] = tag.Name
		}

		respArticles[i] = &content.SelectArticles{
			Id:    currentArticle.ID.Hex(),
			Title: currentArticle.Title,
			Desc:  currentArticle.Desc,
			Tags:  currentArticle.Tags,
		}
	}

	return &content.SelectArticleByTagResponse{
		Articles: respArticles,
		Total:    total,
	}, nil
}
