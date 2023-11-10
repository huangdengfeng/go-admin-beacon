// DB 句柄放入上下文
package dao

import (
	"go-admin-beacon/internal/infrastructure/config"
	"golang.org/x/net/context"
	"gorm.io/gorm"
)

const DbKey = "dbInstance"

// 放入DB句柄，方便事务传播
func withTxDb(ctx context.Context, tx *gorm.DB) context.Context {
	return context.WithValue(ctx, DbKey, tx)
}

func getDbFromContext(ctx context.Context) *gorm.DB {
	value := ctx.Value(DbKey)
	// 返回无事务的DB句柄
	if value == nil {
		return createDbWithContext(ctx)
	}
	return value.(*gorm.DB)
}

func getDb() *gorm.DB {
	if config.DebugEnable {
		return config.SqlClient.Debug().WithContext(context.Background())
	}
	// WithContext 返回是每个会话，可以DB sql之前互不干扰
	return config.SqlClient.WithContext(context.Background())
}

func createDbWithContext(ctx context.Context) *gorm.DB {
	if config.DebugEnable {
		return config.SqlClient.Debug().WithContext(context.Background())
	}
	// WithContext 返回是每个会话，可以DB sql之前互不干扰
	return config.SqlClient.WithContext(ctx)
}
