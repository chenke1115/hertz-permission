/*
 * @Author: changge <changge1519@gmail.com>
 * @Date: 2022-10-27 10:20:25
 * @LastEditTime: 2022-11-14 10:38:08
 * @Description: Do not edit
 */
package route

import (
	"github.com/chenke1115/hertz-permission/pkg/api/auth"
	"github.com/chenke1115/hertz-permission/pkg/api/permission"
	"github.com/chenke1115/hertz-permission/pkg/api/role"
	"github.com/chenke1115/hertz-permission/pkg/api/user"
	"github.com/chenke1115/hertz-permission/pkg/middleware"

	"github.com/cloudwego/hertz/pkg/route"
)

/**
 * @description: load route modules
 * @param {*route.RouterGroup} g
 * @return {*}
 */
func LoadModules(g *route.RouterGroup) {
	LoadModulesWithoutAuth(g)
	LoadModulesWithAuth(g)
}

/**
 * @description: without auth
 * @param {*route.RouterGroup} g
 * @return {*}
 */
func LoadModulesWithoutAuth(g *route.RouterGroup) {
	auth.RegisterAuthRouter(g)
}

/**
 * @description: with auth
 * @param {*route.RouterGroup} g
 * @return {*}
 */
func LoadModulesWithAuth(g *route.RouterGroup) {
	// Jwt and permission
	g.Use(middleware.JwtMiddlewareFunc, middleware.PermissionMiddleware)

	permission.RegisterPermissionRouter(g)
	role.RegisterRoleRouter(g)
	user.RegisterUserRouter(g)
}
