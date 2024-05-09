package logic_test

import (
	"context"
	"fmt"
	"testing"

	"github.com/lius-new/blog-backend/rpc/user/internal/logic"
	"github.com/lius-new/blog-backend/rpc/user/tests"
	"github.com/lius-new/blog-backend/rpc/user/user"
)

func TestModify(t *testing.T) {
	logicClient := logic.NewModifyLogic(context.Background(), tests.SVC_CONTEXT)

	resp, err := logicClient.Modify(&user.ModifyUserRequest{
		Id:       "663b9e602a23854f42fc6df8",
		Username: "lius6666",
		Password: "13",
		SecretId: "661e33b895fb0ef92c541012",
	})
	if err != nil {
		panic(err)
	}
	fmt.Println(resp)
}
