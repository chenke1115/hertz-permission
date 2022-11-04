/*
 * @Author: changge <changge1519@gmail.com>
 * @Date: 2022-11-03 15:47:04
 * @LastEditTime: 2022-11-03 16:07:51
 * @Description: Do not edit
 */
package date

import "time"

/**
 * @description: 时间戳转 Y-m-d H:i:s
 * @param {int} timestamp
 * @return {*}
 */
func DateFormat(timestamp int) string {
	tm := time.Unix(int64(timestamp), 0)
	return tm.Format("2006-01-02 15:04:05")
}

/**
 * @description: 时间戳转 Y-m-d
 * @param {int} timestamp
 * @return {*}
 */
func DateFormatYmd(timestamp int) string {
	tm := time.Unix(int64(timestamp), 0)
	return tm.Format("2006-01-02")
}

/**
 * @description: 当前时间 Y-m
 * @return {*}
 */
func DateYmFormat() string {
	tm := time.Now()
	return tm.Format("2006-01")
}

/**
 * @description: 当前时间 Y-m-d H:i:s
 * @return {*}
 */
func DateNowFormatStr() string {
	tm := time.Now()
	return tm.Format("2006-01-02 15:04:05")
}

/**
 * @description: 当前时间戳
 * @return {*}
 */
func DateUnix() int {
	t := time.Now().Local().Unix()
	return int(t)
}

/**
 * @description:当前年月日时分秒(time类型)
 * @return {*}
 */
func DateNowFormat() time.Time {
	tm := time.Now()
	return tm
}

/**
 * @description: 当前第几周
 * @return {*}
 */
func DateWeek() int {
	_, week := time.Now().ISOWeek()
	return week
}

/**
 * @description: 当前年、月、日
 * @return {*}
 */
func DateYMD() (int, int, int) {
	year, month, day := DateYmdInts()
	return year, month, day
}

/**
 * @description: 当前 Y-m-d
 * @return {*}
 */
func DateYmdFormat() string {
	tm := time.Now()
	return tm.Format("2006-01-02")
}

/**
 * @description: 当前年、月、日
 * @return {*}
 */
func DateYmdInts() (int, int, int) {
	timeNow := time.Now()
	year, month, day := timeNow.Date()
	return year, int(month), day
}
