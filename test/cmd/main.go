/*
 * @Author: changge <changge1519@gmail.com>
 * @Date: 2022-10-26 15:20:57
 * @LastEditTime: 2023-04-10 14:38:33
 * @Description: Do not edit
 */
package main

import (
	"flag"
	"path"

	"github.com/chenke1115/go-common/configs"
	"github.com/chenke1115/hertz-permission/test"
)

var configFile = flag.String("c", path.Base("app.yaml"), "config file")

// var configFile = flag.String("c", "/Users/changge/program/go/src/github.com/chenke1115/hertz-permission/app.yaml", "config file")

// @title                      hertz-permission
// @version                    1.0.0
// @BasePath                   /
// @host                       127.0.0.1:8080
// @schemes                    http
// @securityDefinitions.apikey authorization
// @name                       Authorization
// @in                         header
func main() {
	flag.Parse()

	// Config
	conf := configs.InitConfig(*configFile)

	// Run http server
	test.HttpServer(conf)
}
