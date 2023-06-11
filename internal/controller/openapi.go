package controller

import (
	"context"
	v1 "login-demo/api/v1"
	"login-demo/internal/logic"
)

type OpenapiController struct{}

func (OpenapiController) List(ctx context.Context, req *v1.ProjectListReq) (res *v1.ProjectListRes, err error) {
	str, err := logic.OpenApi.SendOpenAPI()
	res = &v1.ProjectListRes{
		Data: str,
	}
	return res, err
}
