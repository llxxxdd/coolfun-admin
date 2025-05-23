package middleware

import (
	"context"

	"coolfun-admin/api/model/base"
	"coolfun-admin/configs"

	"github.com/cloudwego/hertz/pkg/app"
)

func ForbidOperation(config configs.Config) app.HandlerFunc {
	return func(ctx context.Context, c *app.RequestContext) {
		// pre-handle
		if config.IsDemo && (string(c.Request.Method()) == "POST" || string(c.Request.Method()) == "PUT" || string(c.Request.Method()) == "DELETE") {
			resp := new(base.BaseResp)
			resp.ErrCode = base.ErrCode_Fail
			resp.ErrMsg = "forbidden operation in demo mode"
			c.JSON(403, resp)
			c.Abort()
			return
		}
		c.Next(ctx)
	}
}
