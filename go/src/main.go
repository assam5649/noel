// デバッグ用コード
package main

import (
	"log"
	"src/model"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

var db *gorm.DB

func main() {
	db := model.Connect()
	db = db.Debug()
	router := gin.Default()
	userid := 1
	score, err := model.GetHighScore(db, userid)
	model.UpdateHighScore(db, score, userid)
	a, err := model.GetRanking(db)

	uid := 81
	uname := "assam"
	passwd := "tnv471"

	user := model.User{
		UserID:   uid,
		UserName: uname,
		Password: passwd,
	}

	model.CreateUser(db, &user)
	//model.Userの構造体をもつ変数のポインタを渡す

	var cities []string = []string{"TOKYO", "OSAKA", "OKINAWA"}
	c, e := model.GetScore(db, cities)
	log.Print(model.City[0].Score)

	for i := 0; i < len(cities); i++ {
		if e[i] != nil {
			panic(e[i].Error())
		}
	}

	log.Print(c)
	log.Print(score)

	if err != nil {
		panic(err.Error())
	}
	router.GET("/", func(ctx *gin.Context) {
		ctx.JSON(200, gin.H{
			"score":      score,
			"ranking":    a[1].Score,
			"city score": c,
		})
	})

	router.Run(":8080")
}
