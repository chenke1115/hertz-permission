/*
 * @Author: changge <changge1519@gmail.com>
 * @Date: 2022-10-27 10:35:51
 * @LastEditTime: 2023-04-10 14:30:17
 * @Description: Do not edit
 */
package permission

import (
	"context"

	"github.com/chenke1115/hertz-common/pkg/errors"
	_ "github.com/chenke1115/hertz-common/pkg/errors/validate"
	"github.com/chenke1115/hertz-common/pkg/query"
	"github.com/chenke1115/hertz-common/pkg/response"
	"github.com/chenke1115/hertz-permission/internal/constant/status"
	"github.com/chenke1115/hertz-permission/pkg/model"
	"github.com/cloudwego/hertz/pkg/app"
)

type RespList struct {
	Total       int64                   `json:"total"`
	Permissions *[]model.PermissionShow `json:"data"`
	query.PaginationQuery
}

// ListHandler goDoc
// @Summary     权限列表
// @Description This is a api of permission list
// @Tags        Permission【权限】
// @Accept      json
// @Produce     json
// @Security    authorization
// @Param       query query    model.PermissionQuery true "请求数据"
// @Success     200   {object} response.BaseResponse{data=permission.RespList{}}
// @Failure     400   {object} response.BaseResponse{data=interface{}}
// @Router      /api/permission/list [get]
func ListHandler(ctx context.Context, c *app.RequestContext) {
	var (
		err   error
		query *model.PermissionQuery
		resp  RespList
	)

	// Response
	defer func() {
		if err != nil {
			resp = RespList{}
		}

		response.HandleResponse(c, err, &resp)
	}()

	// Bind and validate
	query = &model.PermissionQuery{}
	err = c.BindAndValidate(query)
	if err != nil {
		err = errors.WrapCode(err, status.RoleParamBindingErrorCode)
		return
	}

	// Bind query param to resp
	resp.PaginationQuery = query.PaginationQuery
	resp.Permissions, resp.Total, err = query.Search()
}
