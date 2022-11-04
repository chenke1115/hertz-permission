/*
 * @Author: changge <changge1519@gmail.com>
 * @Date: 2022-10-11 15:14:59
 * @LastEditTime: 2022-11-03 14:29:33
 * @Description: Do not edit
 */
package match

import (
	"regexp"
	"unicode"
)

const (
	/*
	 * include 0-9, a-z, A-Z, _, and Chinese characters
	 * '_' must appear in the middle and must not be repeated. For example, "__"
	 */
	nicknamePattern = `^[a-z0-9A-Z\p{Han}]+(_[a-z0-9A-Z\p{Han}]+)*$`

	/*
	 * Valid characters are 0-9, a-z, a-z, _
	 * The first letter cannot be _, 0-9
	 * The last letter cannot be _, and _ cannot be consecutive
	 */
	usernamePattern = `^[a-zA-Z][a-z0-9A-Z]*(_[a-z0-9A-Z]+)*$`

	/*
	 * Valid characters are 0-9, a-z, a-z, .
	 * The first letter cannot be ., 0-9
	 * The last letter cannot be ., and . cannot be consecutive
	 */
	keyPattern = `^[a-zA-Z][a-z0-9A-Z]*(.[a-z0-9A-Z]+)*$`
)

var (
	nicknameRegexp = regexp.MustCompile(nicknamePattern)
	usernameRegexp = regexp.MustCompile(usernamePattern)
	keyRegexp      = regexp.MustCompile(keyPattern)
)

/**
 * @description: Check whether it is a valid nickname.
 * @param {[]byte} b
 * @return {*}
 */
func IsNickname(b []byte) bool {
	if len(b) == 0 {
		return false
	}

	return nicknameRegexp.Match(b)
}

/**
 * @description: Same with func IsNickname(b []byte) bool
 * @param {string} str
 * @return {*}
 */
func IsNicknameString(str string) bool {
	if len(str) == 0 {
		return false
	}

	return nicknameRegexp.MatchString(str)
}

/**
 * @description: Check whether the user name is valid
 * @param {[]byte} b
 * @return {*}
 */
func IsUserName(b []byte) bool {
	if len(b) == 0 {
		return false
	}

	return usernameRegexp.Match(b)
}

/**
 * @description: Same with func IsUserName(b []byte) bool
 * @param {string} str
 * @return {*}
 */
func IsUserNameString(str string) bool {
	if len(str) == 0 {
		return false
	}

	return usernameRegexp.MatchString(str)
}

/**
 * @description: Check whether the character is Chinese
 * @param {string} str
 * @return {*}
 */
func IsChineseChar(str string) bool {
	for _, r := range str {
		if unicode.Is(unicode.Scripts["Han"], r) {
			return true
		}
	}

	return false
}

/**
 * @description: Check whether the key is valid
 * @param {string} str
 * @return {*}
 */
func IsKeyString(str string) bool {
	if len(str) == 0 {
		return false
	}

	return keyRegexp.MatchString(str)
}
