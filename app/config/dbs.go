package config

import (
	"time"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

func Database() *gorm.DB {

	db, err := gorm.Open("sqlite3", "redventures")
	if err != nil {
		panic("failed to connect database")
	}
	db.DB().SetConnMaxLifetime(time.Minute*5);
	db.DB().SetMaxIdleConns(0);
	db.DB().SetMaxOpenConns(5);
	return db
}
