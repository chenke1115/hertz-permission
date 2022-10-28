/*
 * @Author: changge <changge1519@gmail.com>
 * @Date: 2022-10-13 16:49:20
 * @LastEditTime: 2022-10-13 17:02:55
 * @Description: Do not edit
 */
package gorm

import (
	"github.com/go-sql-driver/mysql"
	"github.com/mattn/go-sqlite3"
)

const (
	ErrMySQLDupEntry            = 1062
	ErrMySQLDupEntryWithKeyName = 1586
)

/**
 * @description: Check if the error is unique constraint error
 * @param {error} err
 * @return {*}
 */
func IsUniqueConstraintError(err error) bool {
	// sqlite || mysql
	if sqliteErr, ok := err.(sqlite3.Error); ok {
		if sqliteErr.ExtendedCode == sqlite3.ErrConstraintUnique ||
			sqliteErr.ExtendedCode == sqlite3.ErrConstraintPrimaryKey {
			return true
		}
	} else if mysqlErr, ok := err.(*mysql.MySQLError); ok {
		if mysqlErr.Number == ErrMySQLDupEntry ||
			mysqlErr.Number == ErrMySQLDupEntryWithKeyName {
			return true
		}
	}
	return false
}
