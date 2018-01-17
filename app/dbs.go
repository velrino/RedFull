package app

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	Model "./models"
)

func Database() *gorm.DB {
	//docker run --name mysqldocker -v /my/custom:/etc/mysql/conf.d -p 3369:3306 -e MYSQL_ROOT_PASSWORD=root -d mysql:5.7
	db, err := gorm.Open("mysql", "root:root@tcp(127.0.0.1:3369)/rv?charset=utf8&parseTime=True")
	if err != nil {
		panic("failed to connect database")
	}
	return db
}

func Migrations() {
	db := Database()

	db.DropTableIfExists(&Model.User{},&Model.Widget{})
	db.AutoMigrate(&Model.User{},&Model.Widget{})

	// Seeding tables:
	var users []Model.User = []Model.User{
		Model.User{ID: 1, Name: "Colin", Email: "colin@redventures.com", Password: "12345", Gravatar: "http://www.Gravatar.com/avatar/a51972ea936bc3b841350caef34ea47e?s=64&d=monsterID"},
		Model.User{ID: 2, Name: "Kyle", Email: "kyle@redventures.com", Password: "12345", Gravatar:"http://www.Gravatar.com/avatar/432f3e353c689fc37af86ae861d934f9?s=64&d=monsterID"},
		Model.User{ID: 3, Name: "Thomas", Email: "thomas@redventures.com", Password: "12345", Gravatar:"http://www.Gravatar.com/avatar/48009c6a27d25ef0ea03f985d1f186b0?s=64&d=monsterID"},
		Model.User{ID: 4, Name: "James", Email: "james@redventures.com", Password: "12345", Gravatar:"http://www.Gravatar.com/avatar/9372f138140c8578c82bbc77b2eca602?s=64&d=monsterID"},
	}
	
	for _, user := range users {
		db.Save(&user)
	}
	
	defer db.Close()
}

