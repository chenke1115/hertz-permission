/*
 * @Author: changge <changge1519@gmail.com>
 * @Date: 2022-10-31 09:34:57
 * @LastEditTime: 2022-10-31 10:55:14
 * @Description: Do not edit
 */
package role

import "github.com/cloudwego/hertz/pkg/route"

/**
 * @description: Register permission router
 * @param {*route.RouterGroup} g
 * @return {*}
 */
func RegisterRoleRouter(g *route.RouterGroup) {
	// Group
	roleGroup := g.Group("/role")
	roleGroup.GET("/list", ListHandler)
	roleGroup.POST("/add", AddHandler)
	roleGroup.POST("/:id/edit", EditHandler)
	roleGroup.POST("/:id/del", DelHandler)
}
