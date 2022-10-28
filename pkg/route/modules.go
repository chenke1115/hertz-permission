/*
 * @Author: changge <changge1519@gmail.com>
 * @Date: 2022-10-27 10:20:25
 * @LastEditTime: 2022-10-28 09:42:53
 * @Description: Do not edit
 */
package route

import (
	"github.com/chenke1115/ismart-permission/pkg/app/permission"
	"github.com/cloudwego/hertz/pkg/route"
)

/**
 * @description: load route modules
 * @param {*route.RouterGroup} g
 * @return {*}
 */
func LoadModules(g *route.RouterGroup) {
	permission.RegisterPermissionRouter(g)
}
