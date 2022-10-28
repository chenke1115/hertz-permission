/*
 * @Author: changge <changge1519@gmail.com>
 * @Date: 2022-08-26 10:27:26
 * @LastEditTime: 2022-08-29 14:56:35
 * @Description: Do not edit
 */
package configs

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strings"

	jsoniter "github.com/json-iterator/go"
	"gopkg.in/yaml.v3"
)

var (
	options      = defaultOptions()
	isInitConfig bool
)

/**
 * @description: Get config options
 * @return {*}
 */
func GetConf() *Options {
	if !isInitConfig {
		panic("please call InitConfig first")
	}
	return &options
}

/**
 * @description: Init config of iSmart
 * @param {string} fn
 * @return {*}
 */
func InitConfig(fn string) *Options {
	fd, err := os.Open(fn)
	if err != nil {
		log.Panicf(fmt.Sprintf("open conf file:%s error:%v", fn, err.Error()))
	}
	defer fd.Close()

	content, err := ioutil.ReadAll(fd)
	if err != nil {
		log.Panicf(fmt.Sprintf("read conf file:%s error:%v", fn, err.Error()))
	}

	if strings.HasSuffix(fn, ".json") {
		if err = jsoniter.Unmarshal(content, &options); err != nil {
			log.Panicf(fmt.Sprintf("unmarshal conf file:%s error:%v", fn, err.Error()))
		}
	} else if strings.HasSuffix(fn, ".yaml") {
		if err = yaml.Unmarshal(content, &options); err != nil {
			log.Panicf(fmt.Sprintf("unmarshal conf file:%s error:%v", fn, err.Error()))
		}
	}

	// turn true
	isInitConfig = true

	return &options
}
