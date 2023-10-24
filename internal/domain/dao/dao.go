package dao

import (
	"go-admin-beacon/internal/infrastructure/config"
	"gorm.io/gorm"
)

// DB 句柄，如果直接使用全局变量，DB句柄初始化和go 全局变量初始化有顺序要求
type dao struct {
	db func() *gorm.DB
}

func getDb() *gorm.DB {
	if config.DebugEnable {
		return config.SqlClient.Debug()
	}
	return config.SqlClient
}
