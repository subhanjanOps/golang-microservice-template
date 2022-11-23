package gorm

import (
	customError "app-server-gateway-service/pkg/custom_error"
	dbLogger "app-server-gateway-service/pkg/loggers/db_logger"
	"database/sql"
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"time"
)

type IDbConfig interface {
	GetDbConnection() *gorm.DB
}
type DbConfig struct {
	port            string
	host            string
	username        string
	password        string
	dbName          string
	db              *gorm.DB
	connectionRetry time.Duration
}

func NewDbConfig(port, host, dbName, username, password string) *DbConfig {
	dbConfigStr := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?parseTime=true&loc=Local", username, password, host, port, dbName)
	db, err := gorm.Open(
		mysql.New(
			mysql.Config{
				DSN: dbConfigStr,
			},
		),
		&gorm.Config{
			Logger: dbLogger.NewDbLogger(),
		},
	)
	customError.CheckError(err)
	conn, err := db.DB()
	defer func(conn *sql.DB) {
		err := conn.Close()
		customError.CheckError(err)
	}(conn)

	return &DbConfig{
		port:            port,
		host:            host,
		username:        username,
		password:        password,
		dbName:          dbName,
		db:              db,
		connectionRetry: 0,
	}
}
