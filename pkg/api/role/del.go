/*
 * @Author: changge <changge1519@gmail.com>
 * @Date: 2022-10-31 09:46:27
 * @LastEditTime: 2022-10-31 10:54:40
 * @Description: Do not edit
 */
package role

import (
	"context"

	"github.com/chenke1115/ismart-permission/internal/pkg/response"
	"github.com/cloudwego/hertz/pkg/app"
)

/**
 * @description: del
 * @param {context.Context} ctx
 * @param {*app.RequestContext} c
 * @return {*}
 */
func DelHandler(ctx context.Context, c *app.RequestContext) {
	var (
		err error
	)

	defer func() {
		response.HandleResponse(c, err, nil)
	}()
}
