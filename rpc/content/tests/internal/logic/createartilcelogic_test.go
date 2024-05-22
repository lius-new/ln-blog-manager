package logic_test

import (
	"context"
	"fmt"
	"testing"

	"github.com/lius-new/blog-backend/rpc/content/content"
	"github.com/lius-new/blog-backend/rpc/content/internal/logic"
	"github.com/lius-new/blog-backend/rpc/content/tests"
)

func TestCreateArtilce(t *testing.T) {
	ctx := context.Background()

	createArtilceLogic := logic.NewCreateArtilceLogic(ctx, tests.SVC_CONTEXT)

	resp, err := createArtilceLogic.CreateArtilce(&content.CreateArticleRequest{})

	if err != nil {
		fmt.Println("error: ", err)
	}

	fmt.Println(resp)
}
