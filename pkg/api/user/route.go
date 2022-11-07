/*
 * @Author: changge <changge1519@gmail.com>
 * @Date: 2022-11-01 17:39:26
 * @LastEditTime: 2022-11-07 15:28:52
 * @Description: Do not edit
 */
package user

import "github.com/cloudwego/hertz/pkg/route"

func RegisterUserRouter(g *route.RouterGroup) {
	// Group
	userGroup := g.Group("/user")

	userGroup.POST("/:id/assign", AssignHandler)
}
