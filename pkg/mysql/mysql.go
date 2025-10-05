package mysql

import (
	"log"

	"github.com/ovinc-cn/apicenter/v2/pkg/cfg"
	"gorm.io/gorm"
)

var db *gorm.DB

func init() {
	// open conn
	var err error
	db, err = openConn()
	if err != nil {
		log.Fatalf("[MySQL] connect to mysql failed; %s", err)
	}
	// debug
	if cfg.AppDebug() {
		db = db.Debug()
	}
	// try to connect
	sqlDB, err := db.DB()
	if err != nil || sqlDB.Ping() != nil {
		log.Fatalf("[MySQL] ping mysql failed; %s", err)
	}
	// set config
	sqlDB.SetMaxOpenConns(cfg.MySQLMaxOpenConns())
	sqlDB.SetMaxIdleConns(cfg.MySQLMaxIdleConns())
	sqlDB.SetConnMaxLifetime(cfg.MySQLConnMaxLifetime())
	// log
	log.Printf("[MySQL] connect to mysql success\n")
}

func DB() *gorm.DB {
	return db
}
