/*
 * @Author: changge <changge1519@gmail.com>
 * @Date: 2022-10-21 10:13:31
 * @LastEditTime: 2022-10-21 10:18:18
 * @Description: Do not edit
 */
package file

import "os"

/**
 * @description: create dir
 * @param {string} dir
 * @return {*}
 */
func MakeDir(dir string) (err error) {
	if dir == "" {
		dir, _ = os.Getwd()
	}

	err = os.MkdirAll(dir, os.ModePerm)

	return
}
