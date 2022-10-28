/*
 * @Author: changge <changge1519@gmail.com>
 * @Date: 2022-09-01 17:55:26
 * @LastEditTime: 2022-09-28 16:56:27
 * @Description: code of errors
 */
package errors

const (
	OK                = 200
	BadRequest        = 400
	Forbidden         = 403
	NotFound          = 404
	APIVersion        = 410
	InternalServerErr = 500
	BadGateway        = 502
)

type Code struct {
	Status int    `json:"status"`
	Reason string `json:"reason"`
}

var codeMap = map[int]string{
	OK:                "成功",
	BadRequest:        "错误请求",
	Forbidden:         "权限不足，无法访问",
	NotFound:          "找不到资源",
	APIVersion:        "API版本不兼容",
	InternalServerErr: "服务器内部异常",
	BadGateway:        "服务器请求异常",
}

/**
 * @description: func of set new code
 * @param {int} status
 * @param {string} reason
 * @return {*}
 */
func NewCode(status int, reason string) {
	if _, ok := codeMap[status]; ok {
		panic("status existed!")
	}
	codeMap[status] = reason
}

/**
 * @description: func of get reason for code
 * @param {int} status
 * @return {*}
 */
func GetCodeReason(status int) string {
	return codeMap[status]
}

/**
 * @description: func of get code
 * @param {int} status
 * @return {*}
 */
func GetCode(status int) Code {
	return Code{
		Status: status,
		Reason: GetCodeReason(status),
	}
}
