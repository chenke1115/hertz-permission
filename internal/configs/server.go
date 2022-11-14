/*
 * @Author: changge <changge1519@gmail.com>
 * @Date: 2022-08-26 14:52:29
 * @LastEditTime: 2022-08-26 14:57:22
 * @Description: Do not edit
 */
package configs

type Http struct {
	Addr string `json:"addr" yaml:"Addr"`
}

type Sever struct {
	Http Http `json:"http" yaml:"Http"`
}
