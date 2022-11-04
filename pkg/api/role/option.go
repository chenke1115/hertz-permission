/*
 * @Author: changge <changge1519@gmail.com>
 * @Date: 2022-11-01 17:29:31
 * @LastEditTime: 2022-11-03 16:36:25
 * @Description: Do not edit
 */
package role

import (
	"context"

	"github.com/cloudwego/hertz/pkg/app"
)

// OptionHandler goDoc
// @Summary     角色下拉选项
// @Description This is a api of role option
// @Tags        RoleOption
// @Accept      json
// @Produce     json
// @Success     200 {object} response.BaseResponse{data=[]string{}}
// @Failure     400 {object} response.BaseResponse{data=interface{}}
// @Router      /api/role/option [get]
func OptionHandler(ctx context.Context, c *app.RequestContext) {

}
