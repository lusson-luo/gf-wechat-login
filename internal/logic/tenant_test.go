package logic_test

import (
	"context"
	"login-demo/internal/consts"
	"login-demo/internal/logic"
	"testing"

	"github.com/gogf/gf/v2/test/gtest"
)

func TestSetCurrentTenantId(t *testing.T) {
	ctx = context.WithValue(ctx, consts.TenantIDKey, 1)
	// 1. 存在 tenantId
	gtest.C(t, func(t *gtest.T) {
		var tenantId interface{}
		logic.SetCurrentTenantId(ctx, &tenantId)
		t.Assert(tenantId, 1)
	})
	// 2. 不存在 tenantId，异常
	gtest.C(t, func(t *gtest.T) {
		var tenantId interface{}
		err := logic.SetCurrentTenantId(context.Background(), &tenantId)
		t.AssertNE(err, nil)
	})
}
