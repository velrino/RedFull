package app

import (
	_ "github.com/jinzhu/gorm/dialects/mysql"
	Model "github.com/velrino/RedFull/app/models"
	Config "github.com/velrino/RedFull/app/config"
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

	var widgets []Model.Widget = []Model.Widget{
		Model.Widget{ID: 1,Name: "Losenoid",Color: "blue",Price: "9.99",Inventory: 51,Melts: true},
		Model.Widget{ID: 2,Name: "Rowlow",Color: "red",Price: "4.00",Inventory: 7,Melts: true},
		Model.Widget{ID: 3,Name: "Printure",Color: "green",Price: "5.55",Inventory: 18,Melts: false},
		Model.Widget{ID: 4,Name: "Claster",Color: "off-white",Price: "12.56",Inventory: 9,Melts: true},
		Model.Widget{ID: 5,Name: "Pepelexa",Color: "purple",Price: "0.99",Inventory: 0,Melts: false},
	}

	for _, user := range users {
		db.Save(&user)
	}
	
	for _, widget := range widgets {
		db.Save(&widget)
	}

	defer db.Close()
}

