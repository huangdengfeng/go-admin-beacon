package dao

import (
	"context"
	"database/sql"
	log "github.com/sirupsen/logrus"
	"go-admin-beacon/internal/infrastructure/config"
	"go-admin-beacon/internal/infrastructure/constants"
	"go-admin-beacon/internal/infrastructure/errors"
	"golang.org/x/exp/slices"
	"gorm.io/gorm"
	"regexp"
	"strings"
)

// DB 句柄，如果直接使用全局变量，DB句柄初始化和go 全局变量初始化有顺序要求
type dao struct {
	db func() *gorm.DB
}

// 上下文DB句柄标识，可以用字符串
type dbKey struct {
}

var regex = regexp.MustCompile("([a-z0-9])([A-Z])")

// orderBy 例如 createTime asc,status desc
// 检查排序字段防止注入,转下划线
func checkAndConvertOrder(orderBy string, allowedFields []string) (string, error) {
	// 提取排序字段
	fields := strings.Split(orderBy, constants.Comma)
	for _, field := range fields {
		if !slices.Contains(allowedFields, strings.Split(field, " ")[0]) {
			log.Errorf("orderBy [%s] field [%s] not allowed", orderBy, field)
			return "", errors.OrderByNotAllowed
		}
	}
	// 驼峰转下划线
	snakeCase := regex.ReplaceAllString(orderBy, "${1}_${2}")
	result := strings.ToLower(snakeCase)

	return result, nil
}

func DoTransaction(ctx context.Context, fun func(ctx context.Context) error, opts ...*sql.TxOptions) error {
	return getDbFromContext(ctx).Transaction(func(tx *gorm.DB) error {
		err := fun(withTxDb(ctx, tx))
		return err
	}, opts...)
}

// 放入DB句柄，方便事务传播
func withTxDb(ctx context.Context, tx *gorm.DB) context.Context {
	return context.WithValue(ctx, dbKey{}, tx)
}

func getDbFromContext(ctx context.Context) *gorm.DB {
	value := ctx.Value(dbKey{})
	// 返回无事务的DB句柄
	if value == nil {
		if config.DebugEnable {
			return config.SqlClient.Debug().WithContext(ctx)
		}
		// WithContext 返回是每个会话，可以DB sql之前互不干扰
		return config.SqlClient.WithContext(ctx)
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
