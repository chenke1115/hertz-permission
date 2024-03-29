/*
 * @Author: changge <changge1519@gmail.com>
 * @Date: 2022-08-29 14:36:34
 * @LastEditTime: 2023-03-17 13:58:27
 * @Description: Do not edit
 */
package model

import (
	"database/sql/driver"
	"encoding/json"
	"fmt"
	"log"
	"reflect"
	"strings"
	"sync"
	"time"

	"github.com/chenke1115/go-common/configs"
	myLog "github.com/chenke1115/hertz-common/pkg/logs/hlog"
	"github.com/chenke1115/hertz-common/pkg/query"
	"gorm.io/driver/mysql"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

type StringArr []string
type Json map[string]interface{}
type Time time.Time

var db *gorm.DB
var once sync.Once

// Tables use for migration
var Tables []interface{} = []interface{}{
	Permission{},
	Role{}, RolePermission{},
	UserRole{}, UserInfo{}, User{},
}

type DateModel struct {
	CreatedAt *Time `json:"created_at" gorm:"type:datetime"`
	UpdatedAt *Time `json:"updated_at" gorm:"type:datetime"`
}

/**
 * @description: get
 * @return {*}
 */
func GetDB() *gorm.DB {
	once.Do(func() {
		loadDB()
	})
	return db
}

/**
 * @description: load
 * @return {*}
 */
func loadDB() {
	var dialect gorm.Dialector
	database := configs.GetConf().Database
	driver := database.Driver
	debug := configs.GetConf().Debug

	// dsn
	var dsn string
	switch driver {
	case "mysql":
		format := "%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local"
		dsn = fmt.Sprintf(format, database.Username, database.Password, database.Host,
			database.Port, database.Dbname)
		dialect = mysql.Open(dsn)
	case "postgres", "postgresql":
		format := "host=%s port=%d user=%s password=%s dbname=%s sslmode=disable TimeZone=%s"
		dsn = fmt.Sprintf(format, database.Host, database.Port, database.Username,
			database.Password, database.Dbname, "Asia/Shanghai")
		dialect = postgres.Open(dsn)
	default:
		panic(fmt.Errorf("invalid dialector %v", driver))
	}

	// connet
	sqldb, err := gorm.Open(dialect, &gorm.Config{
		// 初始化时禁用外键约束
		DisableForeignKeyConstraintWhenMigrating: true,
	})
	if err != nil {
		panic(fmt.Errorf("数据库连接失败 %v", err.Error()))
	}

	// db log level
	logMode := logger.Warn
	if debug {
		logMode = logger.Info
	}

	// set db session
	db = sqldb.Session(&gorm.Session{
		Logger: logger.New(log.New(myLog.SqlLog(configs.GetConf()), "", log.LstdFlags), logger.Config{
			SlowThreshold:             200 * time.Millisecond,
			LogLevel:                  logMode,
			IgnoreRecordNotFoundError: true,
			Colorful:                  false,
		}),
	})

	if debug {
		// Automatic migration and Up to date
		if err = db.AutoMigrate(Tables...); err != nil {
			log.Panicf("migrate err:%s", err.Error())
		}

		if err = migratorDrop(db); err != nil {
			log.Panicf("migrate err:%s", err.Error())
		}
	}
}

/**
 * @description: drop
 * @param {*gorm.DB} db
 * @return {*}
 */
func migratorDrop(db *gorm.DB) error {
	err := db.Transaction(func(tx *gorm.DB) (err error) {
		return err
	})

	return err
}

/**
 * @description: Convert to String before storing to database
 * @param {interface{}} value
 * @return {*}
 */
func (data *StringArr) Scan(val interface{}) (err error) {
	if val == nil {
		return nil
	}

	if payload, ok := val.([]byte); ok {
		var value []string
		err = json.Unmarshal(payload, &value)
		if err == nil {
			*data = value
		}
	}

	return
}

/**
 * @description: Data is converted to JSON before being read
 * @return {*}
 */
func (data *StringArr) Value() (driver.Value, error) {
	if data == nil {
		return nil, nil
	}

	return json.Marshal([]string(*data))
}

/**
 * @description: StringArr get
 * @return {*}
 */
func (data *StringArr) Get() []string {
	if data == nil {
		return nil
	}

	return []string(*data)
}

/**
 * @description: Convert to String before storing to database
 * @param {interface{}} value
 * @return {*}
 */
func (data *Json) Scan(value interface{}) (err error) {
	if value == nil {
		return nil
	}

	switch val := value.(type) {
	case []byte:
		err = json.Unmarshal(val, &data)
	case string:
		err = json.Unmarshal([]byte(val), &data)
	default:
		err = fmt.Errorf("val type is valid, is %+v", val)
	}

	return
}

/**
 * @description: Data is converted to JSON before being read
 * @return {*}
 */
func (data *Json) Value() (driver.Value, error) {
	vi := reflect.ValueOf(data)
	if vi.IsZero() {
		return nil, nil
	}

	return json.Marshal(Json(*data))
}

/**
 * @description: Format time string
 * @param {[]byte} data
 * @return {*}
 */
func (t *Time) UnmarshalJSON(data []byte) error {
	if string(data) == "null" {
		return nil
	}

	var err error
	str := string(data)
	timeStr := strings.Trim(str, "\"")
	t1, err := time.Parse("2006-01-02 15:04:05", timeStr)
	*t = Time(t1)
	return err
}

/**
 * @description: MarshalJSON on format Time field with %Y-%m-%d %H:%M:%S
 * @return {*}
 */
func (t Time) MarshalJSON() ([]byte, error) {
	formatted := fmt.Sprintf("\"%v\"", time.Time(t).Format("2006-01-02 15:04:05"))
	return []byte(formatted), nil
}

/**
 * @description: Scan valueof time.Time
 * @param {interface{}} value
 * @return {*}
 */
func (t *Time) Scan(v interface{}) error {
	if value, ok := v.(time.Time); ok {
		*t = Time(value)
		return nil
	}
	return fmt.Errorf("can not convert %v to timestamp", v)
}

/**
 * @description: Value insert timestamp into mysql need this function.
 * @return {*}
 */
func (t Time) Value() (driver.Value, error) {
	// Type:Time to time.Time
	var zeroTime time.Time
	tlt := time.Time(t)

	//判断给定时间是否和默认零时间的时间戳相同
	if tlt.UnixNano() == zeroTime.UnixNano() {
		return nil, nil
	}
	return tlt, nil
}

/**
 * @description: Crud all record with query conditions
 * @param {*query.PaginationQuery} p
 * @param {*gorm.DB} tx
 * @param {interface{}} list
 * @return {*}
 */
func crudAll(p *query.PaginationQuery, tx *gorm.DB, list interface{}) (total int64, err error) {
	// Default param
	if p.Offset < 1 {
		p.Offset = 1
	}

	if p.Limit < 1 {
		p.Limit = 10
	}

	// Count
	if err = tx.Count(&total).Error; err != nil {
		return 0, err
	}

	// Data
	if err = tx.Limit(int(p.Limit)).Offset(int(p.Limit * (p.Offset - 1))).Find(list).Error; err != nil {
		return 0, err
	}

	return
}
