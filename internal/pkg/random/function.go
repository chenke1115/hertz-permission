/*
 * @Author: changge <changge1519@gmail.com>
 * @Date: 2022-10-10 15:05:50
 * @LastEditTime: 2022-10-11 10:24:37
 * @Description: Do not edit
 */
package random

import (
	"math/rand"
	"time"
)

/**
 * @description: get random string with length
 * @param {int} l
 * @return {*}
 */
func GetRandomString(l int) string {
	str := "0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	bytes := []byte(str)
	result := []byte{}
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	for i := 0; i < l; i++ {
		result = append(result, bytes[r.Intn(len(bytes))])
	}
	return string(result)
}
