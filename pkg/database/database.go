package database

import (
	"context"
	"database/sql"
	"fmt"
	"github.com/DATA-DOG/go-sqlmock"
	_ "github.com/go-sql-driver/mysql"
	"github.com/ngonghi/vian-backend/config"
	"github.com/uptrace/bun"
	"github.com/uptrace/bun/dialect/mysqldialect"
	"go.uber.org/zap"
)

// InitDatabase ... initializes the database connection.
func InitDatabase(configInstance *config.Config, logger *zap.Logger) (*bun.DB, error) {
	dbConfig := configInstance.Database

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?parseTime=true",
		dbConfig.Mysql.User, dbConfig.Mysql.Password, dbConfig.Mysql.Host, dbConfig.Mysql.Port, dbConfig.Mysql.Name)
	sqldb, err := sql.Open("mysql", dsn)
	if err != nil {
		return nil, err
	}

	db := bun.NewDB(sqldb, mysqldialect.New())
	db.AddQueryHook(NewLoggerQueryHook(logger))
	return db, nil
}

// PingDB ...
func PingDB(context context.Context, db *bun.DB, logger *zap.Logger) bool {
	var num int
	if err := db.QueryRowContext(context, "SELECT 1").Scan(&num); err != nil {
		logger.Error("Database error", zap.Error(err))
		return false
	}
	return true
}

// GetMockDB ... Get mock db connection for unit testing purpose
func GetMockDB() (*bun.DB, sqlmock.Sqlmock, error) {
	db, mock, err := sqlmock.New()
	if err != nil {
		return nil, nil, err
	}

	gdb := bun.NewDB(db, mysqldialect.New())
	if err != nil {
		return nil, nil, err
	}

	return gdb, mock, nil
}
