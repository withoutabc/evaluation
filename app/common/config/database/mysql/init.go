package mysql

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"time"
)

func InitDB(database string) *gorm.DB {
	// 数据库连接配置
	dsn := database
	// 创建连接池
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
	// 设置数据库连接池参数
	sqlDB, _ := db.DB()
	// 设置数据库连接池最大连接数
	sqlDB.SetMaxOpenConns(100)
	// 连接池最大允许的空闲连接数，如果没有sql任务需要执行的连接数大于20，超过的连接会被连接池关闭
	sqlDB.SetMaxIdleConns(20)
	// 设置连接的最大复用时间
	sqlDB.SetConnMaxLifetime(time.Hour)
	return db
}
