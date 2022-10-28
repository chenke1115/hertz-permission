/*
 * @Author: changge <changge1519@gmail.com>
 * @Date: 2022-09-06 16:11:57
 * @LastEditTime: 2022-09-15 16:17:11
 * @Description: Do not edit
 */
package validate

import "github.com/cloudwego/hertz/pkg/app/server/binding"

type BindError struct {
	ErrType, FailField, Msg string
}

// Error implements error interface.
func (e *BindError) Error() string {
	if e.Msg != "" {
		return e.Msg
	}
	return e.FailField + " is invalid"
}

type ValidateError struct {
	ErrType, FailField, Msg string
}

// Error implements error interface.
func (e *ValidateError) Error() string {
	if e.Msg != "" {
		return e.Msg
	}
	return e.FailField + " is validate fail"
}

/**
 * @description: init
 * @return {*}
 */
func init() {
	CustomBindErrFunc := func(failField, msg string) error {
		err := BindError{
			ErrType:   "bindErr",
			FailField: failField,
			Msg:       msg,
		}

		return &err
	}

	CustomValidateErrFunc := func(failField, msg string) error {
		err := ValidateError{
			ErrType:   "validateErr",
			FailField: failField,
			Msg:       msg,
		}

		return &err
	}

	binding.SetErrorFactory(CustomBindErrFunc, CustomValidateErrFunc)
}
