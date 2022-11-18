/*
 * @Author: changge <changge1519@gmail.com>
 * @Date: 2022-10-27 09:53:22
 * @LastEditTime: 2022-11-18 17:20:11
 * @Description: Do not edit
 */
package test

import (
	"context"
	"time"

	"github.com/chenke1115/go-common/configs"
	"github.com/chenke1115/hertz-common/global"
	_ "github.com/chenke1115/hertz-permission/docs"
	"github.com/chenke1115/hertz-permission/pkg/middleware"
	"github.com/chenke1115/hertz-permission/pkg/route"

	"github.com/cloudwego/hertz/pkg/app/server"
	"github.com/cloudwego/hertz/pkg/common/hlog"
	"github.com/hertz-contrib/swagger"
	swaggerFiles "github.com/swaggo/files"
)

func RegisterRoute(h *server.Hertz) {
	// start swagger
	h.GET("/swagger/*any", swagger.WrapHandler(swaggerFiles.Handler))

	// use middleware
	h.Use(
		middleware.Cors(),          // CORS
		middleware.AccessLog(),     // AccessLog
		middleware.Recovery(),      // Recovery
		middleware.Session(),       // Session
		middleware.GlobalSession(), // Session
	)

	// NoRoute
	h.NoRoute(middleware.JwtMiddlewareFunc(), middleware.JwtNoRoute())

	// Group of api
	apiGroup := h.Group("api")
	route.LoadModules(apiGroup)

	// global routers
	global.RouteInfo = h.Routes()
}

func HttpServer(conf *configs.Options) {
	// Writing to log file and defer close
	// f := logs.WriteLog(conf)
	// defer f.Close()

	// server.Default() creates a Hertz with recovery middleware.
	// Maximum wait time before exit, if not specified the default is 5s
	h := server.Default(
		server.WithHostPorts(conf.Server.Http.Addr),
		server.WithExitWaitTime(0*time.Second),
	)

	// Register route
	RegisterRoute(h)

	// Graceful exit
	h.OnShutdown = append(h.OnShutdown, func(ctx context.Context) {
		<-ctx.Done()
		hlog.Warn("exit timeout!")
	})

	// run
	h.Spin()
}
