/*
 * @Author: changge <changge1519@gmail.com>
 * @Date: 2022-09-22 11:35:57
 * @LastEditTime: 2022-11-17 18:15:49
 * @Description: Do not edit
 */
package middleware

import (
	"context"

	"github.com/chenke1115/go-common/functions/conver"
	"github.com/chenke1115/hertz-common/pkg/errors"
	"github.com/chenke1115/hertz-common/pkg/response"
	"github.com/cloudwego/hertz/pkg/app"
)

var PermissionMiddleware = permissionCheck()

/**
 * @description: Check permission
 * @return {*}
 */
func permissionCheck() app.HandlerFunc {
	return func(ctx context.Context, c *app.RequestContext) {
		if currentUser, _ := GetCurrentUser(ctx, c); currentUser != nil {
			// Url string
			url := conver.ToString(c.Request.RequestURI())

			// Check
			if !currentUser.IsSuperUser() {
				if !currentUser.CheckPermission(url) {
					response.HandleResponse(c, errors.New(errors.Forbidden), nil)
					c.Abort()
				}
			}

			c.Next(ctx)
		}
	}
}
