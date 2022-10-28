/*
 * @Author: changge <changge1519@gmail.com>
 * @Date: 2022-08-31 11:50:38
 * @LastEditTime: 2022-10-28 16:48:15
 * @Description: Do not edit
 */
package main

import (
	"flag"
	"path"

	"github.com/chenke1115/ismart-permission/pkg/model"
	"github.com/chenke1115/ismart-permission/test/configs"
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

	// Config of iSmart
	_ = configs.InitConfig(*configFile)

	if err := model.GetDB().Transaction(func(tx *gorm.DB) error {
		// TODO: logic of migrator

		return nil
	}); err != nil {
		hlog.Error(" err: ", err)
		hlog.Info("migrator fail....")
		return
	}
	hlog.Info("migrator success....")
}
