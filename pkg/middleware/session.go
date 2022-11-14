/*
 * @Author: changge <changge1519@gmail.com>
 * @Date: 2022-09-19 17:44:44
 * @LastEditTime: 2022-11-14 15:58:21
 * @Description: Do not edit
 */
package middleware

import (
	"context"
	"fmt"

	"github.com/chenke1115/hertz-permission/internal/configs"
	"github.com/chenke1115/hertz-permission/internal/constant/consts"
	"github.com/chenke1115/hertz-permission/internal/constant/global"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/hertz-contrib/sessions"
	"github.com/hertz-contrib/sessions/redis"
)

/**
 * @description: middleware of session
 * @return {*}
 */
func Session() app.HandlerFunc {
	// Get configs
	redisConf := configs.GetConf().Redis
	// New redis store
	store, err := redis.NewStore(
		redisConf.Size,
		redisConf.Network,
		redisConf.Addr,
		redisConf.Password,
		[]byte(consts.SessionSecret),
	)
	if err != nil {
		panic(fmt.Errorf("缓存初始化失败: %v", err.Error()))
	}

	return sessions.Sessions(consts.SessionName, store)
}

/**
 * @description: Set global session
 * @return {*}
 */
func GlobalSession() app.HandlerFunc {
	return func(ctx context.Context, c *app.RequestContext) {
		global.Session = sessions.Default(c)
	}
}
