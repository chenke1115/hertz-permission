/*
 * @Author: changge <changge1519@gmail.com>
 * @Date: 2022-09-19 13:59:12
 * @LastEditTime: 2022-11-03 14:06:35
 * @Description: Do not edit
 */
package middleware

import (
	"context"

	"github.com/chenke1115/hertz-permission/internal/pkg/response"

	"github.com/cloudwego/hertz/pkg/app"
)

/**
 * @description: Middleware of painc catch
 * @return {*}
 */
func Recovery() app.HandlerFunc {
	return func(ctx context.Context, c *app.RequestContext) {
		defer func() {
			// Catch painc error
			if e := recover(); e != nil {
				response.HandleResponse(c, e.(error), nil)
			}
		}()

		c.Next(ctx)
	}
}
