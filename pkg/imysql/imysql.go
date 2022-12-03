package imysql

import (
	"fmt"

	"github.com/coldwind/artist/pkg/ilog"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	"go.uber.org/zap"
)

type Service struct {
	handler  *gorm.DB
	host     string
	port     int
	username string
	password string
	db       string
	prefix   string
	debug    bool
	maxIdle  int
	maxOpen  int
}

type Option func(*Service)

func New(opts ...Option) *Service {
	s := &Service{}
	for _, f := range opts {
		f(s)
	}

	return s
}

// MysqlInit 初始化
func (s *Service) Run() error {
	connArgs := fmt.Sprintf("%s:%s@(%s:%d)/%s?charset=utf8&parseTime=True&loc=Local", s.username, s.password, s.host, s.port, s.db)
	handle, err := gorm.Open("mysql", connArgs)
	if err != nil {
		return err
	}

	gorm.DefaultTableNameHandler = func(db *gorm.DB, defaultTableName string) string {
		return s.prefix + defaultTableName
	}

	if s.debug {
		handle.LogMode(true)
	}
	s.handler = handle

	// 最大空闲连接数
	s.handler.DB().SetMaxIdleConns(s.maxIdle)
	// 最大连接数
	s.handler.DB().SetMaxOpenConns(s.maxOpen)

	return nil
}

func (m *Service) Handle() *gorm.DB {
	return m.handler
}

func (m *Service) Close() {
	err := m.handler.Close()
	ilog.Info("gorm closed", zap.Error(err))
}
