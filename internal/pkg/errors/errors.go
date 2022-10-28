/*
 * @Author: changge <changge1519@gmail.com>
 * @Date: 2022-09-01 16:55:57
 * @LastEditTime: 2022-09-15 17:21:48
 * @Description: Do not edit
 */
package errors

import (
	"fmt"
	"regexp"

	"github.com/pkg/errors"
)

type Error struct {
	Err  error
	Code Code
}

/**
 * @description: string of err
 * @return {*}
 */
func (e Error) Error() string {
	return e.Code.Reason
}

/**
 * @description: 获取错误调用栈(跳过New,Newf,Wrapf,Wrap调用栈)
 * @return {*}
 */
func (e Error) GetErrStack() errors.StackTrace {
	type stackTracer interface {
		StackTrace() errors.StackTrace
	}

	st := e.Err.(stackTracer)
	stackTrace := st.StackTrace()

	filterStack := errors.StackTrace{}
	filterFuncRegex, _ := regexp.Compile(`/utils/errors\.(New|Wrap)f?`)

	for _, f := range stackTrace[:2] {
		stackText, _ := f.MarshalText()
		if !filterFuncRegex.MatchString(string(stackText)) {
			filterStack = append(filterStack, f)
		}
	}

	filterStack = append(filterStack, stackTrace[2:]...)
	return filterStack
}

/**
 * @description: New
 * @param {int} status
 * @return {*}
 */
func New(status int) error {
	return Newf(status)
}

/**
 * @description: Newf
 * @param {int} status
 * @param {...interface{}} args
 * @return {*}
 */
func Newf(status int, args ...interface{}) error {

	code := GetCode(status)
	if len(args) != 0 {
		code.Reason = fmt.Sprintf(code.Reason, args...)
	}
	return Error{
		Err:  errors.New(code.Reason),
		Code: code,
	}
}

/**
 * @description: Wrap err with code
 * @param {error} err
 * @param {int} status
 * @return {*}
 */
func Wrap(err error, status int) error {
	return Wrapf(err, status)
}

/**
 * @description: Wrapf err with code (支持格式化错误原因)
 * @param {error} err
 * @param {int} status
 * @param {...interface{}} args
 * @return {*}
 */
func Wrapf(err error, status int, args ...interface{}) error {
	code := GetCode(status)
	switch v := err.(type) {
	case Error:
		err = v.Err
		code = v.Code
	default:
		err = errors.WithStack(err)
	}

	if len(args) != 0 {
		code.Reason = fmt.Sprintf(code.Reason, args...)
	}

	// 避免传nil error时日志没有任何调用栈
	if err == nil {
		err = errors.New(code.Reason)
	}

	return Error{
		Err:  err,
		Code: code,
	}
}

/**
 * @description: Wrap code and direct output err.Error() as message
 * @param {error} err
 * @param {int} status
 * @return {*}
 */
func WrapCode(err error, status int) error {
	var code Code
	if err != nil {
		code = Code{
			Status: status,
			Reason: err.Error(),
		}
	} else {
		code = Code{
			Status: InternalServerErr,
			Reason: GetCodeReason(InternalServerErr),
		}
	}

	return Error{
		Err:  err,
		Code: code,
	}
}

/**
 * @description: Cause(获取原始错误)
 * @param {error} err
 * @return {*}
 */
func Cause(err error) error {
	switch v := err.(type) {
	case Error:
		return errors.Cause(v.Err)
	default:
		return errors.Cause(err)
	}
}
