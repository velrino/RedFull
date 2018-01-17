package app

import (
	_ "github.com/jinzhu/gorm/dialects/mysql"
	Model "./models"
	Config "./config"
)

func DatabaseInit() {
	Migrations()
}

func Migrations() {

	db := Config.Database()

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

