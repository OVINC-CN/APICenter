package mysql

import (
	"fmt"

	"github.com/ovinc-cn/apicenter/v2/pkg/cfg"
	"github.com/ovinc-cn/apicenter/v2/pkg/trace"
	"go.opentelemetry.io/otel/attribute"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/plugin/opentelemetry/tracing"
)

func openConn() (*gorm.DB, error) {
	// init dsn
	dsn := fmt.Sprintf(
		"%s:%s@tcp(%s)/%s?%s",
		cfg.MySQLUsername(),
		cfg.MySQLPassword(),
		cfg.MySQLAddr(),
		cfg.MySQLDatabase(),
		cfg.MySQLParams(),
	)
	// init db
	gormDB, err := gorm.Open(
		mysql.Open(dsn),
		&gorm.Config{DisableForeignKeyConstraintWhenMigrating: true},
	)
	if err != nil {
		return nil, err
	}
	// add trace
	if err := gormDB.Use(tracing.NewPlugin(
		tracing.WithoutMetrics(),
		tracing.WithAttributes(
			attribute.String(trace.AttributeDBSystem, "mysql"),
			attribute.String(trace.AttributeDBIP, cfg.MySQLAddr()),
			attribute.String(trace.AttributeDBInstance, cfg.MySQLDatabase()),
		),
	)); err != nil {
		return nil, err
	}
	return gormDB, nil
}
