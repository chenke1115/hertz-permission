/*
 * @Author: changge <changge1519@gmail.com>
 * @Date: 2022-09-22 11:35:57
 * @LastEditTime: 2023-04-10 14:37:53
 * @Description: Do not edit
 */
package middleware

import (
	"context"
	"fmt"
	"regexp"

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
			url, _ = ReplaceStringByRegex(url, "/[0-9]+/", "/:id/")

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

/**
 * @description: ReplaceStringByRegex
 * @param {*} str
 * @param {*} rule
 * @param {string} replace
 * @return {*}
 */
func ReplaceStringByRegex(str, rule, replace string) (string, error) {
	reg, err := regexp.Compile(rule)
	if reg == nil || err != nil {
		return "", fmt.Errorf("regexp fail: %v", err)
	}
	return reg.ReplaceAllString(str, replace), nil
}
