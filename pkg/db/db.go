package db

import (
	"github.com/sirius2001/loon/pkg"
	"github.com/sirius2001/loon/pkg/log"
	"strings"
	"gorm.io/driver/mysql"
	"gorm.io/driver/postgres"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var db *gorm.DB

func dial(dsn string) error {
	log.Info("app database dial", "dsn", dsn)
	switch {
	case strings.HasPrefix(dsn, "user:"):
		// MySQL
		return dailMysql(dsn)
	case strings.HasPrefix(dsn, "host="):
		// PostgreSQL
		return dailPostgres(dsn)
	case strings.HasPrefix(dsn, "file:"):
		// SQLite
		return dailSQLite(dsn)
	default:
		return pkg.ErrDbNotSupport
	}
}

func dailSQLite(dsn string) error {
	var err error
	db, err = gorm.Open(sqlite.Open(dsn), &gorm.Config{})
	if err != nil {
		return pkg.ErrDialDB
	}
	return nil
}

func dailMysql(dsn string) error {
	var err error
	db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		return pkg.ErrDialDB
	}
	return nil
}

func dailPostgres(dsn string) error {
	var err error
	db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return pkg.ErrDialDB
	}
	return nil
}

func NewDB(merge bool, dsn string) error {
	if err := dial(dsn); err != nil {
		return err
	}

	if merge {
		if err := AutoMerge(); err != nil {
			return err
		}
	}

	return nil
}

func GetDB() *gorm.DB {
	return db
}
