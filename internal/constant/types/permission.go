/*
 * @Author: changge <changge1519@gmail.com>
 * @Date: 2022-11-03 10:41:55
 * @LastEditTime: 2022-11-03 10:46:59
 * @Description: Do not edit
 */
package types

type PermissionType struct {
	Dir, Menu, Button string
}

// 权限类型
var PermissionTypeArr = PermissionType{
	Dir:    "D",
	Menu:   "M",
	Button: "B",
}
