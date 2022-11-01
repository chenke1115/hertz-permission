/*
 * @Author: changge <changge1519@gmail.com>
 * @Date: 2022-10-31 09:34:57
 * @LastEditTime: 2022-11-01 17:33:06
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
	roleGroup.GET("/option", OptionHandler)
	roleGroup.POST("/add", AddHandler)

	roleGroup.GET("/:id/info", InfoHandler)
	roleGroup.PUT("/:id/edit", EditHandler)
	roleGroup.DELETE("/:id/del", DelHandler)
	roleGroup.POST("/:id/bind", BindHandler)
}
