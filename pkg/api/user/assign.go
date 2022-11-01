/*
 * @Author: changge <changge1519@gmail.com>
 * @Date: 2022-11-01 17:38:22
 * @LastEditTime: 2022-11-01 18:03:43
 * @Description: Do not edit
 */
package user

import (
	"context"

	"github.com/cloudwego/hertz/pkg/app"
)

// AssignHandler goDoc
// @Summary     角色分配
// @Description This is a api to assign role for user
// @Tags        UserRoleAssign
// @Accept      json
// @Produce     json
// @Success     200 {object} response.BaseResponse{data=interface{}}
// @Failure     400 {object} response.BaseResponse{data=interface{}}
// @Router      /api/user/{id}/assign [post]
func AssignHandler(ctx context.Context, c *app.RequestContext) {

}
