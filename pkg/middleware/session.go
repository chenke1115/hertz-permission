/*
 * @Author: changge <changge1519@gmail.com>
 * @Date: 2022-09-19 17:44:44
 * @LastEditTime: 2022-11-18 16:24:48
 * @Description: Do not edit
 */
package middleware

import (
	"context"
	"fmt"

	"github.com/chenke1115/go-common/configs"
	"github.com/chenke1115/hertz-common/global"
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
	conf := configs.GetConf()
	redisConf := conf.Redis
	// New redis store
	store, err := redis.NewStore(
		redisConf.Size,
		redisConf.Network,
		redisConf.Addr,
		redisConf.Password,
		[]byte(conf.Session.Secret),
	)
	if err != nil {
		panic(fmt.Errorf("缓存初始化失败: %v", err.Error()))
	}

	return sessions.Sessions(conf.Session.Name, store)
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
