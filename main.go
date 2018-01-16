package main

import (
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"net/http"
)

func Database() *gorm.DB {
	//DB Connection
	//docker run --name mysqldocker -v /my/custom:/etc/mysql/conf.d -p 3369:3306 -e MYSQL_ROOT_PASSWORD=root -d mysql:5.7
	db, err := gorm.Open("mysql", "root:root@tcp(127.0.0.1:3369)/redventures?charset=utf8&parseTime=True")
	if err != nil {
		panic("failed to connect database")
	}
	return db
}

func main() {

	//Migrate the schema
	db := Database()
	db.AutoMigrate(&UserMigration{}, &WidgetsMigration{})

	router := gin.Default()

	v1 := router.Group("/api")
	{
		v1.GET("/", FetchAll)
	}
	router.Run(":8979")

}

type LoginRequest struct {
	Email 		string		 `json:"email"`
	Password 	string		 `json:"password"`
}

type UserMigration struct  {
	gorm.Model
	id        uint   `json:"id"`
	email string `json:"email"`
	gravatar  string `json:"gravatar "`	
	name  string `json:"name "`
	password  string `json:"password"`
}

type WidgetsMigration struct  {
	gorm.Model
	id        uint   `json:"id"`
	color string `json:"color"`	
	inventory  int `json:"inventory"`
	name  string `json:"name "`
	price  string `json:"price"`
}

type Mock struct {
	Lorem 	string 	`json:"lorem"`
	Ipsum 		string 	`json:"ipsum"`
}

func FetchAll(c *gin.Context) {
	var _todos []Mock
	c.JSON(http.StatusOK, gin.H{"status": http.StatusOK, "data": _todos})
}