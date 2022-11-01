/*
 * @Author: changge <changge1519@gmail.com>
 * @Date: 2022-10-27 10:35:51
 * @LastEditTime: 2022-11-01 15:14:36
 * @Description: Do not edit
 */
package permission

import (
	"context"

	"github.com/chenke1115/ismart-permission/internal/pkg/response"

	"github.com/cloudwego/hertz/pkg/app"
)

// ListHandler goDoc
// @Summary     权限列表
// @Description This is a api of permission list
// @Tags        PermissionList
// @Accept      json
// @Produce     json
// @Success     200 {object} response.BaseResponse{data=[]model.Permission{}}
// @Failure     400 {object} response.BaseResponse{data=interface{}}
// @Router      /api/permission/list [get]
func ListHandler(ctx context.Context, c *app.RequestContext) {
	var (
		err error
	)

	defer func() {
		response.HandleResponse(c, err, nil)
	}()
}
