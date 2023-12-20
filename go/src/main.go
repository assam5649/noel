package main

import (
	"log"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type menus struct {
	ID   uint
	NAME string
}

func main() {
	var users menus
	/*USER := "root"
	  PASS := "pass"
	  PROTOCOL := "tcp(hackdenoel-mysql:3306)"
	  DBNAME := "noel-db"
	  CONNECT := USER + ":" + PASS + "@" + PROTOCOL + "/" + DBNAME + "?parseTime=true"
	  db, err := gorm.Open(mysql.Open(CONNECT), &gorm.Config{})*/
	user := "root"
	pass := "pass"
	progocol := "tcp(hackdenoel-mysql:3306)"
	dbname := "noel-db"
	connect := user + ":" + pass + "@" + progocol + "/" + dbname
	db, err := gorm.Open(mysql.Open(connect), &gorm.Config{})
	/*var count int

	count = 5
	if err != nil {
		//panic(err.Error())
		for count > 1 {
			time.Sleep(time.Second * 2)
			count--
			log.Println("retry...")
			gorm.Open(mysql.Open(connect), &gorm.Config{})
		}
	}*/

	router := gin.Default()

	result := db.Find(&users)

	if result.Error != nil {
		log.Println("Error fetching data:", result.Error)
	}

	log.Printf("ID: %d, Name: %s", users.ID, users.NAME)

	router.GET("/", func(ctx *gin.Context) {
		ctx.JSON(200, gin.H{
			"ID":   users.ID,
			"NAME": users.NAME,
		})
	})
	if err != nil {
		log.Fatal("Server Run Failed.: ", err)
	}
	router.Run(":8080")
}

/*
func DBMigrate(db *gorm.DB) *gorm.DB {
    db.AutoMigrate(&menus{})
    return db
}*/
//DBMigrate(db)
