/*
 * @Author: changge <changge1519@gmail.com>
 * @Date: 2022-11-11 14:13:19
 * @LastEditTime: 2022-11-11 14:29:38
 * @Description: Do not edit
 */
package validate

import (
	"fmt"
	"regexp"

	"github.com/chenke1115/hertz-permission/internal/pkg/conver"
	"github.com/chenke1115/hertz-permission/internal/pkg/match"

	"github.com/cloudwego/hertz/pkg/app/server/binding"
)

func init() {
	/**
	 * @description: Is nickname
	 * @param {...interface{}} args
	 * @return {*}
	 */
	binding.MustRegValidateFunc("checkNickname", func(args ...interface{}) error {
		if !match.IsNicknameString(conver.Strval(args[0])) {
			return fmt.Errorf("命名不符合规范")
		}

		return nil
	})

	/**
	 * @description: check length and make up of password
	 * @param {...interface{}} args
	 * @return {*}
	 */
	binding.MustRegValidateFunc("checkPassword", func(args ...interface{}) error {
		ps := conver.Strval(args)
		if len(ps) < 9 {
			return fmt.Errorf("密码长度小于9")
		}

		num := `[0-9]{1}`
		a_z := `[a-z]{1}`
		A_Z := `[A-Z]{1}`
		symbol := `[!@#~$%^&*()+|_]{1}`
		if b, err := regexp.MatchString(num, ps); !b || err != nil {
			return fmt.Errorf("密码必须包含数字")
		}
		if b, err := regexp.MatchString(a_z, ps); !b || err != nil {
			return fmt.Errorf("密码必须包含小写字母")
		}
		if b, err := regexp.MatchString(A_Z, ps); !b || err != nil {
			return fmt.Errorf("密码必须包含大写字母")
		}
		if b, err := regexp.MatchString(symbol, ps); !b || err != nil {
			return fmt.Errorf("密码必须包含特殊字符[!@#~$%%^&*()+|_]")
		}

		return nil
	})

	/**
	 * @description: check confirm password
	 * @param {...interface{}} args
	 * @return {*}
	 */
	binding.MustRegValidateFunc("confirmPassword", func(args ...interface{}) error {
		if conver.Strval(args[0]) != conver.Strval(args[1]) {
			return fmt.Errorf("确认密码不等于新密码")
		}

		return nil
	})
}
