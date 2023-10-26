package config

import (
	"encoding/json"
	"flag"
	"fmt"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"gopkg.in/natefinch/lumberjack.v2"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
	"path/filepath"
	"strings"
	"sync"
	"time"
)

var once sync.Once

type Config struct {
	Server struct {
		Name string
		// eg :8080
		Listen          string
		ShutdownTimeout time.Duration
	}
	Log struct {
		Filename string
		// M
		MaxSize int
		// Day
		MaxAge          int
		Level           string
		ReportCaller    bool
		OutputToConsole bool
	}
	DB struct {
		Dsn          string
		User         string
		Password     string
		MaxIdleConns int
		// S
		ConnMaxLifetime time.Duration
		// S
		ConnMaxIdleTime time.Duration
		MaxOpenConns    int
	}
}

const (
	defaultConfigPath = "./conf"
	configName        = "application"
	configType        = "yaml"
)

var (
	DebugEnable      bool
	ServerConfigPath = defaultConfigPath
	Global           = new(Config)
	SqlClient        *gorm.DB
)

// Init 初始化业务全局配置
// config file dir default is ./conf and can be set by flag -configPath
func Init() {
	once.Do(func() {
		// 执行文件 -h 可以查看说明
		if ServerConfigPath == defaultConfigPath {
			flag.StringVar(&ServerConfigPath, "configPath", defaultConfigPath, "server config path")
			if !flag.Parsed() {
				flag.Parse()
			}
		}
		fullConfigPath, err := filepath.Abs(ServerConfigPath)
		if err != nil {
			log.Fatalf("find configPath abs err: %s", err)
		}
		log.Infof("read config file from %s", fullConfigPath)
		viper := viper.New()
		viper.AddConfigPath(fullConfigPath)
		viper.SetConfigName(configName)
		viper.SetConfigType(configType)
		// 环境变量不能有点，viper 对大小写的都能识别
		viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))
		viper.AutomaticEnv()
		if err = viper.ReadInConfig(); err != nil {
			log.Fatalf("read config file errors:%s", err)
		}
		if err = viper.Unmarshal(Global); err != nil {
			log.Fatalf("unmarshal config file errors:%s", err)
		}
		if marshal, err := json.Marshal(Global); err != nil {
			log.Fatalf("json unmarshal errors:%s", err)
		} else {
			log.Infof("config is:%s\n", marshal)
		}
		initLog()
		initDB()
	})
}

// 初始化log https://github.com/sirupsen/logrus
func initLog() {
	logConf := Global.Log
	if logConf.Level == "debug" {
		DebugEnable = true
	}
	level, err := log.ParseLevel(logConf.Level)
	if err != nil {
		panic(err)
	}
	log.SetLevel(level)

	logger := &lumberjack.Logger{
		Filename:   logConf.Filename,
		MaxSize:    logConf.MaxSize,
		MaxAge:     logConf.MaxAge,
		MaxBackups: 0,
		LocalTime:  true,
	}
	log.SetReportCaller(logConf.ReportCaller)
	log.SetFormatter(&log.TextFormatter{FullTimestamp: true, TimestampFormat: time.StampMilli})
	if !logConf.OutputToConsole {
		log.SetOutput(logger)
	}
}

// https://github.com/go-sql-driver/mysql
// https://gorm.io/zh_CN/docs/connecting_to_the_database.html
func initDB() {
	dbConf := Global.DB
	dsn := fmt.Sprintf(dbConf.Dsn, dbConf.User, dbConf.Password)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			// 驼峰表名，不加s
			SingularTable: true,
		},
	})
	if err != nil {
		log.Fatalf("connect db errors:%s", err)
	}
	// 添加连接池
	sqlDB, err := db.DB()
	if err != nil {
		log.Fatalf("connect db errors:%s", err)
	}
	sqlDB.SetMaxIdleConns(dbConf.MaxIdleConns)
	sqlDB.SetConnMaxLifetime(dbConf.ConnMaxLifetime * time.Second)
	sqlDB.SetConnMaxIdleTime(dbConf.ConnMaxIdleTime * time.Second)
	sqlDB.SetMaxOpenConns(dbConf.MaxOpenConns)
	SqlClient = db
}

func Shutdown() {
	if db, err := SqlClient.DB(); nil == err {
		db.Close()
		log.Info("closed db success")
	} else {
		log.Errorf("close db error:%s", err.Error())
	}
}
