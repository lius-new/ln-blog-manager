// Code generated by goctl. DO NOT EDIT.
// Source: content.proto

package contentclient

import (
	"context"

	"github.com/lius-new/blog-backend/rpc/content/content"

	"github.com/zeromicro/go-zero/zrpc"
	"google.golang.org/grpc"
)

type (
	CreateArticleRequest               = content.CreateArticleRequest
	CreateArticleResponse              = content.CreateArticleResponse
	CreateTagRequest                   = content.CreateTagRequest
	CreateTagResponse                  = content.CreateTagResponse
	DeleteArticleRequest               = content.DeleteArticleRequest
	DeleteArticleResponse              = content.DeleteArticleResponse
	DeleteTagRequest                   = content.DeleteTagRequest
	DeleteTagResponse                  = content.DeleteTagResponse
	ModifyArticleContentRequest        = content.ModifyArticleContentRequest
	ModifyArticleContentResponse       = content.ModifyArticleContentResponse
	ModifyArticleCoverRequest          = content.ModifyArticleCoverRequest
	ModifyArticleCoverResponse         = content.ModifyArticleCoverResponse
	ModifyArticleDescRequest           = content.ModifyArticleDescRequest
	ModifyArticleDescResponse          = content.ModifyArticleDescResponse
	ModifyArticleTagRequest            = content.ModifyArticleTagRequest
	ModifyArticleTagResponse           = content.ModifyArticleTagResponse
	ModifyArticleTitleRequest          = content.ModifyArticleTitleRequest
	ModifyArticleTitleResponse         = content.ModifyArticleTitleResponse
	ModifyArticleVisiableByTagRequest  = content.ModifyArticleVisiableByTagRequest
	ModifyArticleVisiableByTagResponse = content.ModifyArticleVisiableByTagResponse
	ModifyArticleVisiableRequest       = content.ModifyArticleVisiableRequest
	ModifyArticleVisiableResponse      = content.ModifyArticleVisiableResponse
	ModifyTagNameRequest               = content.ModifyTagNameRequest
	ModifyTagNameResponse              = content.ModifyTagNameResponse
	ModifyTagPushArticleRequest        = content.ModifyTagPushArticleRequest
	ModifyTagPushArticleResponse       = content.ModifyTagPushArticleResponse
	ModifyTagRemoveArticleRequest      = content.ModifyTagRemoveArticleRequest
	ModifyTagRemoveArticleResponse     = content.ModifyTagRemoveArticleResponse
	ModifyTagVisiableRequest           = content.ModifyTagVisiableRequest
	ModifyTagVisiableResponse          = content.ModifyTagVisiableResponse
	SearchArtilce                      = content.SearchArtilce
	SearchArtilceRequest               = content.SearchArtilceRequest
	SearchArtilceResponse              = content.SearchArtilceResponse
	SelectArticle                      = content.SelectArticle
	SelectArticleByIdRequest           = content.SelectArticleByIdRequest
	SelectArticleByPageRequest         = content.SelectArticleByPageRequest
	SelectArticleByPageResponse        = content.SelectArticleByPageResponse
	SelectArticleByTagRequest          = content.SelectArticleByTagRequest
	SelectArticleByTagResponse         = content.SelectArticleByTagResponse
	SelectArticles                     = content.SelectArticles
	SelectTagByIdRequest               = content.SelectTagByIdRequest
	SelectTagByIdResponse              = content.SelectTagByIdResponse
	SelectTagByPageRequest             = content.SelectTagByPageRequest
	SelectTagByPageResponse            = content.SelectTagByPageResponse

	Content interface {
		// * article create *
		CreateArtilce(ctx context.Context, in *CreateArticleRequest, opts ...grpc.CallOption) (*CreateArticleResponse, error)
		// * article select *
		SelectArtilceById(ctx context.Context, in *SelectArticleByIdRequest, opts ...grpc.CallOption) (*SelectArticle, error)
		// 分页获取文章
		SelectArtilceByPage(ctx context.Context, in *SelectArticleByPageRequest, opts ...grpc.CallOption) (*SelectArticleByPageResponse, error)
		// 根据tag获取文章
		SelectArtilceByTag(ctx context.Context, in *SelectArticleByTagRequest, opts ...grpc.CallOption) (*SelectArticleByTagResponse, error)
		// * article search *
		SearchArtilce(ctx context.Context, in *SearchArtilceRequest, opts ...grpc.CallOption) (*SearchArtilceResponse, error)
		// * article modify *
		ModifyArtilceTitle(ctx context.Context, in *ModifyArticleTitleRequest, opts ...grpc.CallOption) (*ModifyArticleTitleResponse, error)
		// 修改文章描述
		ModifyArtilceDesc(ctx context.Context, in *ModifyArticleDescRequest, opts ...grpc.CallOption) (*ModifyArticleDescResponse, error)
		// 修改文章内容
		ModifyArtilceContent(ctx context.Context, in *ModifyArticleContentRequest, opts ...grpc.CallOption) (*ModifyArticleContentResponse, error)
		// 修改文章标签
		ModifyArtilceTag(ctx context.Context, in *ModifyArticleTagRequest, opts ...grpc.CallOption) (*ModifyArticleTagResponse, error)
		// 修改文章Cover
		ModifyArtilceCover(ctx context.Context, in *ModifyArticleCoverRequest, opts ...grpc.CallOption) (*ModifyArticleCoverResponse, error)
		// 修改文章的可见性
		ModifyArtilceVisiable(ctx context.Context, in *ModifyArticleVisiableRequest, opts ...grpc.CallOption) (*ModifyArticleVisiableResponse, error)
		// 根据tag修改文章的可见性
		ModifyArtilceVisiableByTag(ctx context.Context, in *ModifyArticleVisiableByTagRequest, opts ...grpc.CallOption) (*ModifyArticleVisiableByTagResponse, error)
		// 根据删除文章
		DeleteArtilceById(ctx context.Context, in *DeleteArticleRequest, opts ...grpc.CallOption) (*DeleteArticleResponse, error)
		// ** tag **
		CreateTag(ctx context.Context, in *CreateTagRequest, opts ...grpc.CallOption) (*CreateTagResponse, error)
		// 修改tag name
		ModifyTagName(ctx context.Context, in *ModifyTagNameRequest, opts ...grpc.CallOption) (*ModifyTagNameResponse, error)
		// 修改tag可见性(visiable)
		ModifyTagVisiable(ctx context.Context, in *ModifyTagVisiableRequest, opts ...grpc.CallOption) (*ModifyTagVisiableResponse, error)
		// 添加article到tag
		ModifyTagPushArticle(ctx context.Context, in *ModifyTagPushArticleRequest, opts ...grpc.CallOption) (*ModifyTagPushArticleResponse, error)
		// 从tag中移除article
		ModifyTagRemoveArticle(ctx context.Context, in *ModifyTagRemoveArticleRequest, opts ...grpc.CallOption) (*ModifyTagRemoveArticleResponse, error)
		// 根据分页获取tag
		SelectTagByPage(ctx context.Context, in *SelectTagByPageRequest, opts ...grpc.CallOption) (*SelectTagByPageResponse, error)
		// 根据id获取tag
		SelectTagById(ctx context.Context, in *SelectTagByIdRequest, opts ...grpc.CallOption) (*SelectTagByIdResponse, error)
		// 删除tag
		DeleteTag(ctx context.Context, in *DeleteTagRequest, opts ...grpc.CallOption) (*DeleteTagResponse, error)
	}

	defaultContent struct {
		cli zrpc.Client
	}
)

func NewContent(cli zrpc.Client) Content {
	return &defaultContent{
		cli: cli,
	}
}

// * article create *
func (m *defaultContent) CreateArtilce(ctx context.Context, in *CreateArticleRequest, opts ...grpc.CallOption) (*CreateArticleResponse, error) {
	client := content.NewContentClient(m.cli.Conn())
	return client.CreateArtilce(ctx, in, opts...)
}

// * article select *
func (m *defaultContent) SelectArtilceById(ctx context.Context, in *SelectArticleByIdRequest, opts ...grpc.CallOption) (*SelectArticle, error) {
	client := content.NewContentClient(m.cli.Conn())
	return client.SelectArtilceById(ctx, in, opts...)
}

// 分页获取文章
func (m *defaultContent) SelectArtilceByPage(ctx context.Context, in *SelectArticleByPageRequest, opts ...grpc.CallOption) (*SelectArticleByPageResponse, error) {
	client := content.NewContentClient(m.cli.Conn())
	return client.SelectArtilceByPage(ctx, in, opts...)
}

// 根据tag获取文章
func (m *defaultContent) SelectArtilceByTag(ctx context.Context, in *SelectArticleByTagRequest, opts ...grpc.CallOption) (*SelectArticleByTagResponse, error) {
	client := content.NewContentClient(m.cli.Conn())
	return client.SelectArtilceByTag(ctx, in, opts...)
}

// * article search *
func (m *defaultContent) SearchArtilce(ctx context.Context, in *SearchArtilceRequest, opts ...grpc.CallOption) (*SearchArtilceResponse, error) {
	client := content.NewContentClient(m.cli.Conn())
	return client.SearchArtilce(ctx, in, opts...)
}

// * article modify *
func (m *defaultContent) ModifyArtilceTitle(ctx context.Context, in *ModifyArticleTitleRequest, opts ...grpc.CallOption) (*ModifyArticleTitleResponse, error) {
	client := content.NewContentClient(m.cli.Conn())
	return client.ModifyArtilceTitle(ctx, in, opts...)
}

// 修改文章描述
func (m *defaultContent) ModifyArtilceDesc(ctx context.Context, in *ModifyArticleDescRequest, opts ...grpc.CallOption) (*ModifyArticleDescResponse, error) {
	client := content.NewContentClient(m.cli.Conn())
	return client.ModifyArtilceDesc(ctx, in, opts...)
}

// 修改文章内容
func (m *defaultContent) ModifyArtilceContent(ctx context.Context, in *ModifyArticleContentRequest, opts ...grpc.CallOption) (*ModifyArticleContentResponse, error) {
	client := content.NewContentClient(m.cli.Conn())
	return client.ModifyArtilceContent(ctx, in, opts...)
}

// 修改文章标签
func (m *defaultContent) ModifyArtilceTag(ctx context.Context, in *ModifyArticleTagRequest, opts ...grpc.CallOption) (*ModifyArticleTagResponse, error) {
	client := content.NewContentClient(m.cli.Conn())
	return client.ModifyArtilceTag(ctx, in, opts...)
}

// 修改文章Cover
func (m *defaultContent) ModifyArtilceCover(ctx context.Context, in *ModifyArticleCoverRequest, opts ...grpc.CallOption) (*ModifyArticleCoverResponse, error) {
	client := content.NewContentClient(m.cli.Conn())
	return client.ModifyArtilceCover(ctx, in, opts...)
}

// 修改文章的可见性
func (m *defaultContent) ModifyArtilceVisiable(ctx context.Context, in *ModifyArticleVisiableRequest, opts ...grpc.CallOption) (*ModifyArticleVisiableResponse, error) {
	client := content.NewContentClient(m.cli.Conn())
	return client.ModifyArtilceVisiable(ctx, in, opts...)
}

// 根据tag修改文章的可见性
func (m *defaultContent) ModifyArtilceVisiableByTag(ctx context.Context, in *ModifyArticleVisiableByTagRequest, opts ...grpc.CallOption) (*ModifyArticleVisiableByTagResponse, error) {
	client := content.NewContentClient(m.cli.Conn())
	return client.ModifyArtilceVisiableByTag(ctx, in, opts...)
}

// 根据删除文章
func (m *defaultContent) DeleteArtilceById(ctx context.Context, in *DeleteArticleRequest, opts ...grpc.CallOption) (*DeleteArticleResponse, error) {
	client := content.NewContentClient(m.cli.Conn())
	return client.DeleteArtilceById(ctx, in, opts...)
}

// ** tag **
func (m *defaultContent) CreateTag(ctx context.Context, in *CreateTagRequest, opts ...grpc.CallOption) (*CreateTagResponse, error) {
	client := content.NewContentClient(m.cli.Conn())
	return client.CreateTag(ctx, in, opts...)
}

// 修改tag name
func (m *defaultContent) ModifyTagName(ctx context.Context, in *ModifyTagNameRequest, opts ...grpc.CallOption) (*ModifyTagNameResponse, error) {
	client := content.NewContentClient(m.cli.Conn())
	return client.ModifyTagName(ctx, in, opts...)
}

// 修改tag可见性(visiable)
func (m *defaultContent) ModifyTagVisiable(ctx context.Context, in *ModifyTagVisiableRequest, opts ...grpc.CallOption) (*ModifyTagVisiableResponse, error) {
	client := content.NewContentClient(m.cli.Conn())
	return client.ModifyTagVisiable(ctx, in, opts...)
}

// 添加article到tag
func (m *defaultContent) ModifyTagPushArticle(ctx context.Context, in *ModifyTagPushArticleRequest, opts ...grpc.CallOption) (*ModifyTagPushArticleResponse, error) {
	client := content.NewContentClient(m.cli.Conn())
	return client.ModifyTagPushArticle(ctx, in, opts...)
}

// 从tag中移除article
func (m *defaultContent) ModifyTagRemoveArticle(ctx context.Context, in *ModifyTagRemoveArticleRequest, opts ...grpc.CallOption) (*ModifyTagRemoveArticleResponse, error) {
	client := content.NewContentClient(m.cli.Conn())
	return client.ModifyTagRemoveArticle(ctx, in, opts...)
}

// 根据分页获取tag
func (m *defaultContent) SelectTagByPage(ctx context.Context, in *SelectTagByPageRequest, opts ...grpc.CallOption) (*SelectTagByPageResponse, error) {
	client := content.NewContentClient(m.cli.Conn())
	return client.SelectTagByPage(ctx, in, opts...)
}

// 根据id获取tag
func (m *defaultContent) SelectTagById(ctx context.Context, in *SelectTagByIdRequest, opts ...grpc.CallOption) (*SelectTagByIdResponse, error) {
	client := content.NewContentClient(m.cli.Conn())
	return client.SelectTagById(ctx, in, opts...)
}

// 删除tag
func (m *defaultContent) DeleteTag(ctx context.Context, in *DeleteTagRequest, opts ...grpc.CallOption) (*DeleteTagResponse, error) {
	client := content.NewContentClient(m.cli.Conn())
	return client.DeleteTag(ctx, in, opts...)
}
