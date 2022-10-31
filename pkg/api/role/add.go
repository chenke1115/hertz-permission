/*
 * @Author: changge <changge1519@gmail.com>
 * @Date: 2022-10-31 09:44:07
 * @LastEditTime: 2022-10-31 09:50:37
 * @Description: Do not edit
 */
package role

import (
	"context"

	"github.com/chenke1115/ismart-permission/internal/pkg/response"
	"github.com/cloudwego/hertz/pkg/app"
)

/**
 * @description: add
 * @param {context.Context} ctx
 * @param {*app.RequestContext} c
 * @return {*}
 */
func AddHandler(ctx context.Context, c *app.RequestContext) {
	var (
		err error
	)

	defer func() {
		response.HandleResponse(c, err, nil)
	}()
}
