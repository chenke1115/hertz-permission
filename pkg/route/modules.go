/*
 * @Author: changge <changge1519@gmail.com>
 * @Date: 2022-10-27 10:20:25
 * @LastEditTime: 2022-11-07 17:40:17
 * @Description: Do not edit
 */
package route

import (
	"github.com/chenke1115/hertz-permission/pkg/api/permission"
	"github.com/chenke1115/hertz-permission/pkg/api/role"
	"github.com/chenke1115/hertz-permission/pkg/api/user"

	"github.com/cloudwego/hertz/pkg/route"
)

/**
 * @description: load route modules
 * @param {*route.RouterGroup} g
 * @return {*}
 */
func LoadModules(g *route.RouterGroup) {
	permission.RegisterPermissionRouter(g)
	role.RegisterRoleRouter(g)
	user.RegisterUserRouter(g)

}
