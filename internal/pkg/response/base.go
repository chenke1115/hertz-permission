/*
 * @Author: changge <changge1519@gmail.com>
 * @Date: 2022-08-31 15:20:40
 * @LastEditTime: 2022-10-28 09:45:08
 * @Description: Do not edit
 */
package response

import (
	"runtime"

	"github.com/chenke1115/ismart-permission/internal/pkg/conver"
	"github.com/chenke1115/ismart-permission/internal/pkg/errors"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/protocol/consts"
)

type BaseResponse struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data,omitempty"`
}

/**
 * @description: Get result
 * @param {error} err
 * @param {interface{}} resp
 * @return {*}
 */
func getResponse(err error, resp interface{}) *BaseResponse {
	// nit code and message
	baseResult := BaseResponse{consts.StatusOK, consts.StatusMessage(consts.StatusOK), resp}

	// Set code and message
	if err != nil {
		switch v := err.(type) {
		case errors.Error:
			baseResult.Code = v.Code.Status
			baseResult.Message = v.Code.Reason
		case runtime.Error:
			baseResult.Code = errors.InternalServerErr
			baseResult.Message = errors.GetCodeReason(errors.InternalServerErr)
		default:
			baseResult.Code = errors.InternalServerErr
			baseResult.Message = conver.Strval(err.Error())
		}
	}

	return &baseResult
}

/**
 * @description: Handle of Response
 * @param {*app.RequestContext} ctx
 * @param {error} err
 * @param {interface{}} resp
 * @return {*}
 */
func HandleResponse(ctx *app.RequestContext, err error, resp interface{}) {
	HandleResponseWithStatus(ctx, consts.StatusOK, err, resp)
}

/**
 * @description: Return json with code
 * @param {*app.RequestContext} ctx
 * @param {int} code
 * @param {error} err
 * @param {interface{}} resp
 * @return {*}
 */
func HandleResponseWithStatus(ctx *app.RequestContext, code int, err error, resp interface{}) {
	baseResult := getResponse(err, resp)

	// If http Status, set http code
	if consts.StatusMessage(baseResult.Code) != "Unknown Status Code" {
		code = baseResult.Code
	}

	ctx.JSON(code, baseResult)
}
