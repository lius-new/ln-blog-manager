package logic_test

import (
	"context"
	"fmt"
	"testing"

	"github.com/lius-new/blog-backend/rpc/content/content"
	"github.com/lius-new/blog-backend/rpc/content/internal/logic"
	"github.com/lius-new/blog-backend/rpc/content/tests"
)

func TestCreateCovers(t *testing.T) {
	ctx := context.Background()

	createCoversLogic := logic.NewCreateCoversLogic(ctx, tests.SVC_CONTEXT)

	resp, err := createCoversLogic.CreateCovers(&content.CreateCoversRequest{})

	if err != nil {
		fmt.Println("error: ", err)
	}

	fmt.Println(resp)
}
