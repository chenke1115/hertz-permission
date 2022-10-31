/*
 * @Author: changge <changge1519@gmail.com>
 * @Date: 2022-10-27 10:29:31
 * @LastEditTime: 2022-10-31 10:36:20
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
	permissionGroup.POST("/add", AddHandler)
	permissionGroup.POST("/:id/edit", EditHandler)
	permissionGroup.POST("/:id/del", DelHandler)
}
