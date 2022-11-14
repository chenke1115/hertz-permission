/*
 * @Author: changge <changge1519@gmail.com>
 * @Date: 2022-08-24 17:39:59
 * @LastEditTime: 2022-08-25 11:43:03
 * @Description: Do not edit
 */
package middleware

import (
	"context"
	"time"

	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/common/hlog"
)

/**
 * @description: AccessLog middleware
 * @return {*}
 */
func AccessLog() app.HandlerFunc {
	return func(c context.Context, ctx *app.RequestContext) {
		start := time.Now()
		ctx.Next(c)
		end := time.Now()
		latency := end.Sub(start).Microseconds
		hlog.CtxTracef(c, "status=%d cost=%d method=%s full_path=%s client_ip=%s host=%s",
			ctx.Response.StatusCode(), latency,
			ctx.Request.Header.Method(), ctx.Request.URI().PathOriginal(), ctx.ClientIP(), ctx.Request.Host())
	}
}
