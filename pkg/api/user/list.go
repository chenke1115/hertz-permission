/*
 * @Author: changge <changge1519@gmail.com>
 * @Date: 2022-11-09 16:42:44
 * @LastEditTime: 2022-11-14 17:11:58
 * @Description: Do not edit
 */
package user

import (
	"context"

	"github.com/chenke1115/hertz-permission/internal/pkg/errors"
	"github.com/chenke1115/hertz-permission/internal/pkg/query"
	"github.com/chenke1115/hertz-permission/internal/pkg/response"
	"github.com/chenke1115/hertz-permission/pkg/model"

	"github.com/cloudwego/hertz/pkg/app"
)

type RespList struct {
	Total int64            `json:"total"`
	Users *[]model.APIUser `json:"data"`
	query.PaginationQuery
}

// ListHandler goDoc
// @Summary     用户列表
// @Description This is a api for user list
// @Tags        UserList
// @Accept      json
// @Produce     json
// @Success     200 {object} response.BaseResponse{data=user.RespList{}}
// @Failure     400 {object} response.BaseResponse{data=interface{}}
// @Router      /api/user/list [get]
func ListHandler(ctx context.Context, c *app.RequestContext) {
	var (
		err   error
		query *model.UserQuery
		resp  RespList
	)

	defer func() {
		if err != nil {
			resp = RespList{}
		}

		response.HandleResponse(c, err, &resp)
	}()

	query = &model.UserQuery{}
	err = c.BindAndValidate(query)
	if err != nil {
		err = errors.WrapCode(err, errors.BadRequest)
		return
	}

	// bind query param to resp
	resp.PaginationQuery = query.PaginationQuery

	resp.Users, resp.Total, err = query.Search()
}
