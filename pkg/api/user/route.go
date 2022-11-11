/*
 * @Author: changge <changge1519@gmail.com>
 * @Date: 2022-11-01 17:39:26
 * @LastEditTime: 2022-11-11 16:12:32
 * @Description: Do not edit
 */
package user

import "github.com/cloudwego/hertz/pkg/route"

func RegisterUserRouter(g *route.RouterGroup) {
	// Group
	userGroup := g.Group("/user")
	userGroup.GET("/current", CurrentUserHandler)
	userGroup.GET("/list", ListHandler)
	userGroup.POST("/add", AddHandler)
	userGroup.PUT("/:id/edit", EditHandler)
	userGroup.POST("/:id/assign", AssignHandler)
	userGroup.PUT("/:id/state", StateHandler)
	userGroup.PUT("/:id/reset", ResetHandler)
}
