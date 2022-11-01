/*
 * @Author: changge <changge1519@gmail.com>
 * @Date: 2022-10-27 10:35:51
 * @LastEditTime: 2022-11-01 15:11:27
 * @Description: Do not edit
 */
package role

import (
	"context"

	"github.com/chenke1115/ismart-permission/internal/constant/status"
	"github.com/chenke1115/ismart-permission/internal/pkg/errors"
	"github.com/chenke1115/ismart-permission/internal/pkg/query"
	"github.com/chenke1115/ismart-permission/internal/pkg/response"
	"github.com/chenke1115/ismart-permission/pkg/model"

	"github.com/cloudwego/hertz/pkg/app"
)

type RespList struct {
	Total int64         `json:"total"`
	Roles *[]model.Role `json:"roles"`
	query.PaginationQuery
}

// ListHandler goDoc
// @Summary     角色列表
// @Description This is a api of role list
// @Tags        RoleList
// @Accept      json
// @Produce     json
// @Success     200 {object} response.BaseResponse{data=role.RespList{}}
// @Failure     400 {object} response.BaseResponse{data=interface{}}
// @Router      /api/role/list [get]
func ListHandler(ctx context.Context, c *app.RequestContext) {
	var (
		err   error
		query *model.RoleQuery
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
	query = &model.RoleQuery{}
	err = c.BindAndValidate(query)
	if err != nil {
		err = errors.WrapCode(err, status.RoleParamBindingErrorCode)
		return
	}

	// Bind query param to resp
	resp.PaginationQuery = query.PaginationQuery

	resp.Roles, resp.Total, err = query.Search()
}
