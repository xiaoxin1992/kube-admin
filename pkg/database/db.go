package database

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/xiaoxin1992/kube-admin/pkg/config"
	"github.com/xiaoxin1992/kube-admin/pkg/logger"
	"go.uber.org/zap"
	"time"
)

var p *pool

func GetPool() *pool {

	if p == nil {
		p = &pool{
			logger: logger.GetLogger().S("MySQL"),
		}
	}
	return p
}

type pool struct {
	db     *sql.DB
	logger *zap.SugaredLogger
}

func (p *pool) Create() error {
	host := config.GetConfig().MySQL.Host
	port := config.GetConfig().MySQL.Port
	username := config.GetConfig().MySQL.User
	password := config.GetConfig().MySQL.Password
	dbName := config.GetConfig().MySQL.Database

	dataSourceName := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8&loc=Local&parseTime=true", username, password, host, port, dbName)
	db, err := sql.Open("mysql", dataSourceName)
	if err != nil {
		p.logger.Errorf("open mysql error %v", err)
		return err
	}
	if err = db.Ping(); err != nil {
		p.logger.Errorf("mysql ping error %v", err)
		return err
	}
	maxLifeTime := time.Duration(config.GetConfig().MySQL.MaxLifetime) * time.Second
	maxOpen := config.GetConfig().MySQL.MaxOpen
	maxIdle := config.GetConfig().MySQL.MaxIdle
	db.SetConnMaxLifetime(maxLifeTime)
	db.SetMaxOpenConns(maxOpen)
	db.SetMaxIdleConns(maxIdle)

	p.db = db
	p.logger.Infof("Create MySQL Connection Pool Successfully.")
	return nil
}

func (p *pool) Destroy() {
	if p.db != nil {
		p.logger.Infof("Destroy MySQL Connection Pool Successfully.")
		err := p.db.Close()
		if err != nil {
			p.logger.Errorf("Destroy MySQL Connection Pool Error: %v", err)
		}
	}
}

func (p *pool) GetSqlDB() *sql.DB {
	return p.db
}
