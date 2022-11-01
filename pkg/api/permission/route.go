/*
 * @Author: changge <changge1519@gmail.com>
 * @Date: 2022-10-27 10:29:31
 * @LastEditTime: 2022-11-01 17:54:56
 * @Description: Do not edit
 */
package permission

import "github.com/cloudwego/hertz/pkg/route"

/**
 * @description: Register permission router
 * @param {*route.RouterGroup} g
 * @return {*}
 */
func RegisterPermissionRouter(g *route.RouterGroup) {
	// Group
	permissionGroup := g.Group("/permission")
	permissionGroup.GET("/list", ListHandler)
	permissionGroup.GET("/option", OptionHandler)
	permissionGroup.POST("/add", AddHandler)

	permissionGroup.GET("/:id/info", InfoHandler)
	permissionGroup.PUT("/:id/edit", EditHandler)
	permissionGroup.DELETE("/:id/del", DelHandler)
}
