package logic

import (
	"context"

	"github.com/gogf/gf/v2/net/ghttp"
)

type LogicContext struct {
}

var Ctx LogicContext = LogicContext{}

const (
	HTTPContextKey = "paramsCtx"
)

type HTTPContextParams struct {
	Username string
}

func (LogicContext) SetHTTPContextParams(r *ghttp.Request, username string) {
	r.SetCtxVar(HTTPContextKey, HTTPContextParams{
		Username: username,
	})
}

func (LogicContext) GetHTTPContextParams(ctx context.Context) *HTTPContextParams {
	v, ok := ctx.Value(HTTPContextKey).(HTTPContextParams)
	if ok {
		return &v
	} else {
		return nil
	}
}
