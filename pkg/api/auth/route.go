/*
 * @Author: changge <changge1519@gmail.com>
 * @Date: 2022-11-11 09:56:19
 * @LastEditTime: 2022-11-14 10:37:12
 * @Description: Do not edit
 */
package auth

import (
	"github.com/chenke1115/hertz-permission/pkg/middleware"
	"github.com/cloudwego/hertz/pkg/route"
)

/**
 * @description: Registe router and func of auth
 * @param {*route.RouterGroup} g
 * @return {*}
 */
func RegisterAuthRouter(g *route.RouterGroup) {
	// Group
	auth := g.Group("/auth")
	// Login
	auth.POST("/login", LoginHandler)

	auth.Use(middleware.JwtMiddlewareFunc, middleware.PermissionMiddleware)
	{
		// Refresh time can be longer than token timeout
		auth.POST("/refresh_token", RefreshHandler)
		auth.POST("/logout", LogoutHandler)
	}
}
