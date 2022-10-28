/*
 * @Author: changge <changge1519@gmail.com>
 * @Date: 2022-10-26 15:20:57
 * @LastEditTime: 2022-10-28 09:42:44
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

func main() {
	flag.Parse()

	// Config
	conf := configs.InitConfig(*configFile)

	// Run http server
	test.HttpServer(conf)
}
