/*
 * @Author: changge <changge1519@gmail.com>
 * @Date: 2022-08-31 11:50:38
 * @LastEditTime: 2023-08-03 15:15:33
 * @Description: Do not edit
 */
package main

import (
	"errors"
	"flag"
	"path"

	"github.com/chenke1115/go-common/configs"
	"github.com/chenke1115/hertz-permission/pkg/model"
	"github.com/cloudwego/hertz/pkg/common/hlog"
	"gorm.io/gorm"
)

var configFile = flag.String("c", path.Base("app.yaml"), "config file")

/**
 * @description: main
 * @return {*}
 */
func main() {
	flag.Parse()

	// Config
	conf := configs.InitConfig(*configFile)

	if err := model.GetDB().Transaction(func(tx *gorm.DB) error {
		// Create super role
		if len(conf.App.User.Super) < 1 {
			return errors.New("please complete the config file")
		}
		role := model.Role{
			Name: "超级管理员",
			Key:  conf.App.User.Super[0],
		}
		err := role.Create(tx)
		if err != nil {
			return err
		}

		// Create super user
		userInfo := model.UserInfo{
			Name:    "Admin",
			Account: "admin",
		}
		err = userInfo.Create(tx)
		if err != nil {
			return err
		}

		// Assign role for user
		userRole := model.UserRole{
			UID:    1,
			RoleID: 1,
		}
		err = userRole.Create(tx)
		if err != nil {
			return err
		}

		return nil
	}); err != nil {
		hlog.Error(" err: ", err)
		hlog.Info("migrator fail....")
		return
	}
	hlog.Info("migrator success....")
}
