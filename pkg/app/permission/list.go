/*
 * @Author: changge <changge1519@gmail.com>
 * @Date: 2022-10-27 10:35:51
 * @LastEditTime: 2022-10-28 09:43:01
 * @Description: Do not edit
 */
package permission

import (
	"context"

	"github.com/chenke1115/ismart-permission/internal/pkg/response"
	"github.com/cloudwego/hertz/pkg/app"
)

/**
 * @description: list
 * @param {context.Context} ctx
 * @param {*app.RequestContext} c
 * @return {*}
 */
func PermissionList(ctx context.Context, c *app.RequestContext) {
	var (
		err error
	)

	defer func() {
		response.HandleResponse(c, err, nil)
	}()
}
