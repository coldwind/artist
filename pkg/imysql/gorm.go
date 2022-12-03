package imysql

import (
	"fmt"
	"wnwdkj_ws/boot"
	"wnwdkj_ws/conf"
	"wnwdkj_ws/pkg/logger"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	"go.uber.org/zap"
)

type MysqlClass struct {
	handler *gorm.DB
}

var instance *MysqlClass = nil

func New() *MysqlClass {
	if instance == nil {
		instance = &MysqlClass{}
	}

	return instance
}

// MysqlInit 初始化
func (m *MysqlClass) Run(args *boot.BootArgs) {
	mysqlConf := conf.New().GetMysqlConf()
	connArgs := fmt.Sprintf("%s:%s@(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local", mysqlConf.Username, mysqlConf.Password, mysqlConf.Host, mysqlConf.Port, mysqlConf.DB)
	handle, err := gorm.Open("mysql", connArgs)
	if err != nil {
		return
	}

	gorm.DefaultTableNameHandler = func(db *gorm.DB, defaultTableName string) string {
		return mysqlConf.Prefix + defaultTableName
	}

	if mysqlConf.Debug {
		handle.LogMode(true)
	}
	m.handler = handle

	// 最大空闲连接数
	m.handler.DB().SetMaxIdleConns(mysqlConf.MaxIdle)
	// 最大连接数
	m.handler.DB().SetMaxOpenConns(mysqlConf.MaxOpen)
}

func (m *MysqlClass) Handle() *gorm.DB {
	return instance.handler
}

func (m *MysqlClass) Close() {
	err := instance.handler.Close()
	logger.Info("gorm closed", zap.Error(err))
}
