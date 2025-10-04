package mysql

import (
	"database/sql"
	"log"

	"github.com/ovinc-cn/apicenter/v2/pkg/cfg"
	"gorm.io/gorm"
)

var db *sql.DB

func init() {
	gormDB, err := openConn()
	if err != nil {
		log.Fatalf("[MySQL] connect to mysql failed; %s", err)
	}
	// try to connect
	db, err = gormDB.DB()
	if err != nil || db.Ping() != nil {
		log.Fatalf("[MySQL] ping mysql failed; %s", err)
	}
	// set config
	db.SetMaxOpenConns(cfg.MySQLMaxOpenConns())
	db.SetMaxIdleConns(cfg.MySQLMaxIdleConns())
	db.SetConnMaxLifetime(cfg.MySQLConnMaxLifetime())
	// log
	log.Printf("[MySQL] connect to mysql success\n")
}

func DB() *sql.DB {
	return db
}

func GormDB() (*gorm.DB, error) {
	return openConn()
}
