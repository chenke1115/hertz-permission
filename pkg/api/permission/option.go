/*
 * @Author: changge <changge1519@gmail.com>
 * @Date: 2022-11-01 17:53:08
 * @LastEditTime: 2022-11-04 18:14:03
 * @Description: Do not edit
 */
package permission

import (
	"context"

	"github.com/chenke1115/hertz-permission/internal/pkg/response"
	"github.com/chenke1115/hertz-permission/pkg/model"

	"github.com/cloudwego/hertz/pkg/app"
)

// OptionHandler goDoc
// @Summary     权限下拉选项
// @Description This is a api of permission option
// @Tags        PermissionOption
// @Accept      json
// @Produce     json
// @Success     200 {object} response.BaseResponse{data=[]model.PermissionOption{}}
// @Failure     400 {object} response.BaseResponse{data=interface{}}
// @Router      /api/permission/option [get]
func OptionHandler(ctx context.Context, c *app.RequestContext) {
	var (
		err  error
		resp []model.PermissionOption
	)

	// Response
	defer func() {
		if err != nil {
			resp = []model.PermissionOption{}
		}

		response.HandleResponse(c, err, &resp)
	}()

	option := model.Permission{}
	resp, err = option.Option()
}
