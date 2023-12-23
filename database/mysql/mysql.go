package mysql

import (
	"database/sql"
	"fmt"
	"os"
	"time"

	"github.com/armuh16/kbfinansia/config"
	"github.com/armuh16/kbfinansia/package/logger"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	gormLogger "gorm.io/gorm/logger"
)

type DB struct {
	Gorm *gorm.DB
	Sql  *sql.DB
}

func NewMysql(log *logger.LogRus) *DB {
	var err error
	dsn := fmt.Sprintf("%s:%s@tcp(%s:3306)/%s?charset=utf8mb4&parseTime=True&loc=Local", config.Get().DBUser, config.Get().DBPassword, config.Get().DBHost, config.Get().DBName)

	gormDB, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		Logger: gormLogger.Default.LogMode(gormLogger.Info),
		// Cache the prepared statement to improve performance
		PrepareStmt: true,
	})

	if err != nil {
		log.Fatal("🚀 Failed to connect to the Database! \n", err.Error())
		os.Exit(1)
	}

	sqldb, err := gormDB.DB()
	if err != nil {
		log.Fatalf("failed to get DB from gorm : %v", err.Error())
	}

	if err := sqldb.Ping(); err != nil {
		log.Fatalf("failed to ping database : %v", err.Error())
	}

	sqldb.SetMaxOpenConns(100)
	sqldb.SetMaxIdleConns(10)
	sqldb.SetConnMaxIdleTime(300 * time.Second)
	sqldb.SetConnMaxLifetime(time.Duration(300 * time.Second))

	fmt.Println("👍 Migration Successfully connected to the Database!")

	return &DB{
		Gorm: gormDB,
		Sql:  sqldb,
	}
}
