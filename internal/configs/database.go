/*
 * @Author: changge <changge1519@gmail.com>
 * @Date: 2022-08-26 16:00:00
 * @LastEditTime: 2022-08-29 16:33:53
 * @Description: Do not edit
 */
package configs

type Database struct {
	Driver   string `json:"driver" yaml:"Driver"`
	Host     string `json:"host" yaml:"Host"`
	Port     int    `json:"port" yaml:"Port"`
	Username string `json:"username" yaml:"Username"`
	Password string `json:"password" yaml:"Password"`
	Dbname   string `json:"dbname" yaml:"Dbname"`
}
