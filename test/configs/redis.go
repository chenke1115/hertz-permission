/*
 * @Author: changge <changge1519@gmail.com>
 * @Date: 2022-10-12 17:54:30
 * @LastEditTime: 2022-10-13 09:54:06
 * @Description: Do not edit
 */
package configs

type Redis struct {
	Network  string `json:"network" yaml:"Network"`
	Addr     string `json:"addr" yaml:"Addr"`
	Password string `json:"password" yaml:"Password"`
	Size     int    `json:"size" yaml:"Size"`
}
