/*
 * @Author: changge <changge1519@gmail.com>
 * @Date: 2022-08-25 11:42:11
 * @LastEditTime: 2022-11-14 09:48:35
 * @Description: Do not edit
 */
package logs

import (
	"fmt"
	"os"
	"time"

	"github.com/chenke1115/hertz-permission/internal/configs"
	"github.com/chenke1115/hertz-permission/internal/pkg/file"

	"github.com/cloudwego/hertz/pkg/common/hlog"
)

/**
 * @description: write log to logfile
 * @return {*}
 */
func WriteLog(conf *configs.Options) *os.File {
	// set log level
	level := hlog.LevelInfo
	if conf.Debug {
		level = hlog.LevelDebug
	}
	hlog.SetLevel(level)

	path := conf.Log.Dir + time.Now().Format("200601") + "/"
	fileName := time.Now().Format("20060102") + ".log"
	_, err := os.Stat(path + fileName)
	switch {
	case os.IsNotExist(err):
		if err = file.MakeDir(path); err != nil {
			panic(fmt.Errorf("log file not exist:%v", err.Error()))
		}
	case os.IsPermission(err):
		panic(fmt.Errorf("log file permission:%v", err.Error()))
	}

	// output
	f, err := os.OpenFile(path+fileName, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		panic(fmt.Errorf("fail to open log file: %v", err.Error()))
	}
	hlog.SetOutput(f)

	// defer f.Close()

	return f
}
