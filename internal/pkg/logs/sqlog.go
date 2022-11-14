/*
 * @Author: changge <changge1519@gmail.com>
 * @Date: 2022-10-21 10:07:47
 * @LastEditTime: 2022-11-14 09:48:13
 * @Description: Do not edit
 */
package logs

import (
	"os"
	"time"

	"github.com/chenke1115/hertz-permission/internal/configs"
	"github.com/chenke1115/hertz-permission/internal/pkg/file"

	"github.com/cloudwego/hertz/pkg/common/hlog"
)

/**
 * @description: log of sql
 * @param {*configs.Options} conf
 * @return {*}
 */
func SqlLog(conf *configs.Options) *os.File {
	path := conf.Log.Dir + time.Now().Format("200601") + "/"
	fileName := time.Now().Format("20060102") + "_sql.log"
	_, err := os.Stat(path + fileName)
	switch {
	case os.IsNotExist(err):
		if err = file.MakeDir(path); err != nil {
			hlog.Fatalf("Log File Not Exist:%v", err)
		}
	case os.IsPermission(err):
		hlog.Fatalf("Log File Permission:%v", err)
	}

	f, err := os.OpenFile(path+fileName, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		hlog.Fatalf("Fail to Open Log File:%v", err)
	}

	// defer func() {
	// 	err := f.Close()
	// 	if err != nil {
	// 		hlog.Warnf("Sql Log File Close Fail: %v", err.Error())
	// 	} else {
	// 		hlog.Infof("Sql Log File Is Close")
	// 	}
	// }()

	return f
}
