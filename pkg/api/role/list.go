/*
 * @Author: changge <changge1519@gmail.com>
 * @Date: 2022-10-27 10:35:51
 * @LastEditTime: 2022-11-14 17:12:47
 * @Description: Do not edit
 */
package role

import (
	"context"
	"strconv"

	"github.com/chenke1115/hertz-permission/internal/constant/status"
	"github.com/chenke1115/hertz-permission/internal/pkg/errors"
	"github.com/chenke1115/hertz-permission/internal/pkg/query"
	"github.com/chenke1115/hertz-permission/internal/pkg/response"
	"github.com/chenke1115/hertz-permission/pkg/model"

	"github.com/cloudwego/hertz/pkg/app"
)

type RespList struct {
	Total int64             `json:"total"`
	Roles *[]model.RoleShow `json:"data"`
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

// BindListHandler goDoc
// @Summary     绑定权限数组
// @Description This is a api of bind permission for role
// @Tags        RoleBindPermissionList
// @Accept      json
// @Produce     json
// @Param       id  query    int true "角色ID" example(1)
// @Success     200 {object} response.BaseResponse{data=[]int{}}
// @Failure     400 {object} response.BaseResponse{data=interface{}}
// @Router      /api/role/{id}/bind [get]
func BindListHandler(ctx context.Context, c *app.RequestContext) {
	var (
		err  error
		ID   int
		resp []int
		role model.Role
	)

	// Response
	defer func() {
		if err != nil {
			resp = []int{}
		}

		response.HandleResponse(c, err, &resp)
	}()

	// ID
	if ID, err = strconv.Atoi(c.Param("id")); err != nil {
		err = errors.Wrap(err, status.RoleIdMissCode)
		return
	}

	// Find
	if role, err = model.GetRoleByID(ID); err != nil {
		return
	}

	// Get
	resp, err = role.BindList()
}
