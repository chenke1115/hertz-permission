/*
 * @Author: changge <changge1519@gmail.com>
 * @Date: 2022-08-26 11:53:03
 * @LastEditTime: 2022-10-25 14:40:41
 * @Description: Do not edit
 */
package configs

type Options struct {
	Debug    bool     `json:"debug" yaml:"Debug"`
	Server   Sever    `json:"server" yaml:"Server"`
	Log      Log      `json:"log" yaml:"Log"`
	Database Database `json:"database" yaml:"Database"`
	Redis    Redis    `json:"redis" yaml:"Redis"`
	Swagger  Swagger  `json:"swagger" yaml:"Swagger"`
}

/**
 * @description: default
 * @return {*}
 */
func defaultOptions() (option Options) {
	return
}
