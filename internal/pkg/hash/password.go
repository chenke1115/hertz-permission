/*
 * @Author: changge <changge1519@gmail.com>
 * @Date: 2022-09-13 11:29:33
 * @LastEditTime: 2022-09-15 15:27:23
 * @Description: Do not edit
 */
package hash

import (
	"crypto/sha256"
	"encoding/hex"
)

/**
 * @description: Encrypt password by hash (Minimum memory consumption is sha256)
 * @param {string} password
 * @param {string} salt
 * @return {*}
 */
func GetHashedPassword(password string, salt string) string {
	tmp := sha256.Sum256([]byte(password + salt))
	return hex.EncodeToString(tmp[:])
}

/**
 * @description: check
 * @param {string} password
 * @param {string} salt
 * @param {string} hashedPassword
 * @return {*}
 */
func CheckHashedPassword(password string, salt string, hashedPassword string) bool {
	return GetHashedPassword(password, salt) == hashedPassword
}
