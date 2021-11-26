package db

import (
	"eicesoft/web-demo/config"
	"fmt"
	"time"

	"github.com/pkg/errors"
	"go.uber.org/zap"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
)

var _ Repo = (*DbRepo)(nil)

type Repo interface {
	p()
	GetDbR() *gorm.DB
	GetDbW() *gorm.DB
	DbRClose() error
	DbWClose() error
	Shutdown(logger *zap.Logger)
}

type DbRepo struct {
	DbR *gorm.DB
	DbW *gorm.DB
}

func New() (*DbRepo, error) {
	cfg := config.Get().MySQL
	dbr, err := dbConnect(cfg.Read.User, cfg.Read.Pass, cfg.Read.Addr, cfg.Read.Name)
	if err != nil {
		return nil, err
	}

	dbw, err := dbConnect(cfg.Write.User, cfg.Write.Pass, cfg.Write.Addr, cfg.Write.Name)
	if err != nil {
		return nil, err
	}

	return &DbRepo{
		DbR: dbr,
		DbW: dbw,
	}, nil
}

func (d *DbRepo) p() {}

func (d *DbRepo) GetDbR() *gorm.DB {
	return d.DbR
}

func (d *DbRepo) GetDbW() *gorm.DB {
	return d.DbW
}

func (d *DbRepo) DbRClose() error {
	sqlDB, err := d.DbR.DB()
	if err != nil {
		return err
	}
	return sqlDB.Close()
}

func (d *DbRepo) DbWClose() error {
	sqlDB, err := d.DbW.DB()
	if err != nil {
		return err
	}
	return sqlDB.Close()
}

func (d *DbRepo) Shutdown(logger *zap.Logger) {
	if err := d.DbWClose(); err != nil {
		logger.Error("dbw close err", zap.Error(err))
	} else {
		logger.Info("dbw close success")
	}

	if err := d.DbRClose(); err != nil {
		logger.Error("dbr close err", zap.Error(err))
	} else {
		logger.Info("dbr close success")
	}
}

func dbConnect(user, pass, addr, dbName string) (*gorm.DB, error) {
	dsn := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8mb4&parseTime=%t&loc=%s",
		user,
		pass,
		addr,
		dbName,
		true,
		"Local")

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			SingularTable: true,
		},
		//Logger: logger.Default.LogMode(logger.Info), // 日志配置
	})

	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("[db connection failed] Database name: %s", dbName))
	}

	db.Set("gorm:table_options", "CHARSET=utf8mb4")

	cfg := config.Get().MySQL.Base

	sqlDB, err := db.DB()
	if err != nil {
		return nil, err
	}

	// 设置连接池 用于设置最大打开的连接数，默认值为0表示不限制.设置最大的连接数，可以避免并发太高导致连接mysql出现too many connections的错误。
	sqlDB.SetMaxOpenConns(cfg.MaxOpenConn)

	// 设置最大连接数 用于设置闲置的连接数.设置闲置的连接数则当开启的一个连接使用完成后可以放在池里等候下一次使用。
	sqlDB.SetMaxIdleConns(cfg.MaxIdleConn)

	// 设置最大连接超时
	sqlDB.SetConnMaxLifetime(time.Minute * cfg.ConnMaxLifeTime)

	// 使用插件
	db.Use(&TracePlugin{})

	return db, nil
}
