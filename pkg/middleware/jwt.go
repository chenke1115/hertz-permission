/*
 * @Author: changge <changge1519@gmail.com>
 * @Date: 2022-08-22 10:48:17
 * @LastEditTime: 2022-11-14 16:52:04
 * @Description: Do not edit
 */
package middleware

import (
	"context"
	"fmt"
	"time"

	"github.com/chenke1115/hertz-permission/internal/constant/consts"
	"github.com/chenke1115/hertz-permission/internal/constant/global"
	"github.com/chenke1115/hertz-permission/internal/constant/status"
	"github.com/chenke1115/hertz-permission/internal/pkg/array"
	"github.com/chenke1115/hertz-permission/internal/pkg/errors"
	"github.com/chenke1115/hertz-permission/internal/pkg/hash"
	"github.com/chenke1115/hertz-permission/pkg/model"

	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/common/hlog"
	"github.com/hertz-contrib/jwt"
)

var (
	identityKey    = consts.CurrentUserKey
	uidKey         = "uid"
	nameKey        = "name"
	accountKey     = "account"
	customerIDKey  = "customer_id"
	rolesKey       = "roles"
	permissionsKey = "permissions"

	JwtMiddlewareFunc = Jwt().MiddlewareFunc()
	JwtLoginHandler   = Jwt().LoginHandler
	JwtLogoutHandler  = Jwt().LogoutHandler
	JwtRefreshHandler = Jwt().RefreshHandler
)

type login struct {
	Username string `form:"username,required" json:"username,required"` //lint:ignore SA5008 ignoreCheck
	Password string `form:"password,required" json:"password,required"` //lint:ignore SA5008 ignoreCheck
}

/**
 * @description: the jwt middleware
 * @return *jwt.HertzJWTMiddleware
 */
func Jwt() *jwt.HertzJWTMiddleware {
	authMiddleware, err := jwt.New(&jwt.HertzJWTMiddleware{
		Realm:       consts.AppName,
		Key:         []byte("secret key"),
		Timeout:     time.Hour,
		MaxRefresh:  time.Hour,
		IdentityKey: identityKey,
		PayloadFunc: func(data interface{}) jwt.MapClaims {
			if v, ok := data.(*model.CurrentUser); ok {
				return jwt.MapClaims{
					uidKey:         v.ID,
					nameKey:        v.Name,
					accountKey:     v.Account,
					customerIDKey:  v.CustomerID,
					rolesKey:       v.Roles,
					permissionsKey: v.Permissions,
				}
			}
			return jwt.MapClaims{}
		},
		IdentityHandler: func(ctx context.Context, c *app.RequestContext) interface{} {
			claims := jwt.ExtractClaims(ctx, c)
			return &model.CurrentUser{
				ID:          int(claims[uidKey].(float64)),
				Name:        claims[nameKey].(string),
				Account:     claims[accountKey].(string),
				CustomerID:  int(claims[customerIDKey].(float64)),
				Roles:       array.ToArray(claims[rolesKey]),
				Permissions: []string{},
			}
		},
		Authenticator: func(ctx context.Context, c *app.RequestContext) (interface{}, error) {
			var loginVals login
			if err := c.BindAndValidate(&loginVals); err != nil {
				return "", jwt.ErrMissingLoginValues
			}

			username := loginVals.Username
			password := loginVals.Password

			if cuser, err := model.CheckUsernameAndPassword(username, hash.GetHashedPassword(password, consts.Salt)); err == nil {
				return &cuser, nil
			}

			return nil, jwt.ErrFailedAuthentication
		},
		Authorizator: func(data interface{}, ctx context.Context, c *app.RequestContext) bool {
			if _, ok := data.(*model.CurrentUser); ok {
				return true
			}
			return false
		},
		Unauthorized: func(ctx context.Context, c *app.RequestContext, code int, message string) {
			c.JSON(code, map[string]interface{}{
				"code":    code,
				"message": message,
			})
		},

		// TokenLookup is a string in the form of "<source>:<name>" that is used
		// to extract token from the request.
		// Optional. Default value "header:Authorization".
		// Possible values:
		// - "header:<name>"
		// - "query:<name>"
		// - "cookie:<name>"
		// - "param:<name>"
		TokenLookup: "header: Authorization, query: token, cookie: jwt",

		// TokenHeadName is a string in the header. Default value is "Bearer"
		TokenHeadName: "Bearer",

		// TimeFunc provides the current time. You can override it to use another time value. This is useful for testing or if your server uses a different time zone than your tokens.
		TimeFunc: time.Now,
	})

	if err != nil {
		hlog.Fatal("JWT Error:" + err.Error())
	}

	// When you use jwt.New(), the function is already automatically called for checking,
	// which means you don't need to call it again.
	errInit := authMiddleware.MiddlewareInit()

	if errInit != nil {
		hlog.Fatal("authMiddleware.MiddlewareInit() Error:" + errInit.Error())
	}

	return authMiddleware
}

/**
 * @description: NoRoute
 * @use h.NoRoute(jwt.MiddlewareFunc, jwt.NoRoute())
 * @return app.HandlerFunc
 */
func JwtNoRoute() app.HandlerFunc {
	return func(ctx context.Context, c *app.RequestContext) {
		claims := jwt.ExtractClaims(ctx, c)
		hlog.Warnf("NoRoute claims: %#v\n", claims)
		c.JSON(404, map[string]string{"code": "PAGE_NOT_FOUND", "message": "Page not found"})
	}
}

/**
 * @description: check is current user id
 * @param {int} id
 * @param {context.Context} ctx
 * @param {*app.RequestContext} c
 * @return {*}
 */
func CheckCurrentUser(id int, ctx context.Context, c *app.RequestContext) (err error) {
	// Get token data
	var user *model.CurrentUser
	if data, ok := c.Get(identityKey); ok {
		user = data.(*model.CurrentUser)
	}

	// Check
	if user.ID != id {
		err = errors.New(status.UserPermissionDeniedCode)
	}

	return
}

/**
 * @description: Get current user from token
 * @param {context.Context} ctx
 * @param {*app.RequestContext} c
 * @return {*}
 */
func GetCurrentUser(ctx context.Context, c *app.RequestContext) (user *model.CurrentUser, err error) {
	// Get token data
	if data, ok := c.Get(identityKey); ok {
		user = data.(*model.CurrentUser)
	}

	if user == nil {
		err = errors.New(errors.Forbidden)
		return
	}

	// Get cache
	cacheKey := fmt.Sprintf(consts.UserPermissionCacheKey, user.ID)
	cacheVal := global.Session.Get(cacheKey)
	if cacheVal != nil {
		user.Permissions = cacheVal.([]string)
		return
	}

	// Permission
	user.Permissions, _ = model.GetPermissionsByUID(user.ID)
	if user.Permissions == nil {
		user.Permissions = []string{}
	}

	// Set cache
	global.Session.Set(cacheKey, user.Permissions)
	global.Session.Save()

	return
}

/**
 * @description: Clean current user cache
 * @param {context.Context} ctx
 * @param {*app.RequestContext} c
 * @return {*}
 */
func CleanCurrentUserCache(ctx context.Context, c *app.RequestContext) (err error) {
	// Get token data
	var user *model.CurrentUser
	if data, ok := c.Get(identityKey); ok {
		user = data.(*model.CurrentUser)
	}

	if user == nil {
		err = errors.New(errors.Forbidden)
		return
	}

	cacheKey := fmt.Sprintf(consts.UserPermissionCacheKey, user.ID)
	cacheVal := global.Session.Get(cacheKey)
	if cacheVal != nil {
		global.Session.Clear()
	}

	return
}
