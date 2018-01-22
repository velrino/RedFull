package config

import (
	"time"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

func Database() *gorm.DB {
	//docker run --name mysqldocker -v /my/custom:/etc/mysql/conf.d -p 3369:3306 -e MYSQL_ROOT_PASSWORD=root -d mysql:5.7
	db, err := gorm.Open("sqlite3", "test.db")
	if err != nil {
		panic("failed to connect database")
	}
	db.DB().SetConnMaxLifetime(time.Minute*5);
	db.DB().SetMaxIdleConns(0);
	db.DB().SetMaxOpenConns(5);
	return db
}
