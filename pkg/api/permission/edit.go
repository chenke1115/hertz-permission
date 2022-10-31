/*
 * @Author: changge <changge1519@gmail.com>
 * @Date: 2022-10-31 09:45:52
 * @LastEditTime: 2022-10-31 10:36:45
 * @Description: Do not edit
 */
package permission

import (
	"context"

	"github.com/chenke1115/ismart-permission/internal/pkg/response"
	"github.com/cloudwego/hertz/pkg/app"
)

/**
 * @description: edit
 * @param {context.Context} ctx
 * @param {*app.RequestContext} c
 * @return {*}
 */
func EditHandler(ctx context.Context, c *app.RequestContext) {
	var (
		err error
	)

	defer func() {
		response.HandleResponse(c, err, nil)
	}()
}
