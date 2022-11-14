/*
 * @Author: changge <changge1519@gmail.com>
 * @Date: 2022-08-19 17:09:52
 * @LastEditTime: 2022-11-14 14:15:19
 * @Description: Do not edit
 */
package middleware

import (
	"time"

	"github.com/cloudwego/hertz/pkg/app"
	"github.com/hertz-contrib/cors"
)

// CORS
func Cors() app.HandlerFunc {
	return cors.New(cors.Config{
		AllowOrigins: []string{"*"}, // Allowed domains, need to bring schema
		AllowMethods: []string{
			"GET",
			"POST",
			"PUT",
			"PATCH",
			"DELETE",
			"HEAD",
		}, // Allowed request methods
		AllowHeaders: []string{
			"Origin",
			"Content-Length",
			"Content-Type",
			"Cookie",
			"Set-Cookie",
			"Accept",
			"Token",
			"Authorization",
			"X-CSRF-TOKEN",
			"X-XSRF-TOKEN",
			"X-Access-Token",
			"X-Requested-With",
		}, // Allowed request headers
		ExposeHeaders: []string{
			"Access-Control-Allow-Origin",
			"Access-Control-Allow-Headers",
			"Content-Length",
			"Content-Type",
			"Cache-Control",
		}, // Request headers allowed in the upload_file
		AllowCredentials: true, // Whether cookies are attached
		AllowOriginFunc: func(origin string) bool { // Custom domain detection with lower priority than AllowOrigins
			return origin == "*"
		},
		MaxAge: 12 * time.Hour, // Maximum length of upload_file-side cache preflash requests (seconds)
	})
}
