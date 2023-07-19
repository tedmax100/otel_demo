package database

import (
	"fmt"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	"github.com/uptrace/opentelemetry-go-extra/otelsql"
	"github.com/uptrace/opentelemetry-go-extra/otelsqlx"
	semconv "go.opentelemetry.io/otel/semconv/v1.10.0"
	"go.uber.org/fx"
)

var Module = fx.Options(
	fx.Provide(ProvideDB),
)

func ProvideDB(dsn string) (*sqlx.DB, error) {
	db, err := otelsqlx.Open("mysql", dsn,
		otelsql.WithAttributes(semconv.DBSystemMSSQL))
	if err != nil {
		fmt.Println("connection to mysql failed:", err)
		return nil, err
	}
	err = db.Ping()
	if err != nil {
		fmt.Println("connection to mysql failed:", err)
		return nil, err
	}
	return db, nil
}
