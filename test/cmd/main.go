/*
 * @Author: changge <changge1519@gmail.com>
 * @Date: 2022-10-26 15:20:57
 * @LastEditTime: 2022-10-31 13:47:48
 * @Description: Do not edit
 */
package main

import (
	"flag"
	"path"

	"github.com/chenke1115/ismart-permission/test"
	"github.com/chenke1115/ismart-permission/test/configs"
)

var configFile = flag.String("c", path.Base("app.yaml"), "config file")

// @title    iSmart-Permission
// @version  1.0.0
// @BasePath /
// @host     127.0.0.1:8080
// @schemes  http
func main() {
	flag.Parse()

	// Config
	conf := configs.InitConfig(*configFile)

	// Run http server
	test.HttpServer(conf)
}
