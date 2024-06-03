/*
 * @Author: changge <changge1519@gmail.com>
 * @Date: 2022-10-27 09:53:22
 * @LastEditTime: 2023-08-01 17:38:52
 * @Description: Do not edit
 */
package test

import (
	"context"
	"time"

	"github.com/chenke1115/go-common/configs"
	"github.com/chenke1115/hertz-common/global"
	myLog "github.com/chenke1115/hertz-common/pkg/logs/hlog"
	mw "github.com/chenke1115/hertz-common/pkg/middleware"
	"github.com/chenke1115/hertz-common/pkg/redis"
	"github.com/chenke1115/hertz-permission/docs"
	"github.com/chenke1115/hertz-permission/pkg/middleware"
	"github.com/chenke1115/hertz-permission/pkg/route"
	"github.com/cloudwego/hertz/pkg/app/middlewares/server/recovery"
	"github.com/cloudwego/hertz/pkg/app/server"
	"github.com/cloudwego/hertz/pkg/common/hlog"
)

func RegisterRoute(h *server.Hertz) {
	// use middleware
	h.Use(
		mw.Cors(),      // CORS
		mw.AccessLog(), // AccessLog
		// mw.Recovery(),  // Recovery
		recovery.Recovery(recovery.WithRecoveryHandler(mw.MyRecoveryHandler)),
		mw.Swagger(mw.WithSwaggerHandler(h, docs.SwaggerInfo)),
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
	myLog.WriteLog(conf)

	// Init redis
	err := redis.InitClient(conf)
	if err != nil {
		panic(err)
	}

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
