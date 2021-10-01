package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

type musics struct {
	Music_ID     int    `json:music_ID`
	Music_URL    string `json:music_URL`
	Music_artist string `json:music_artist`
	Music_name   string `json:music_name`
} //データをjsonで受け取る為の構造体

func gormConnect() (database *gorm.DB) {
	DBMS := "mysql"
	USER := "root"
	PASS := "root"
	PROTOCOL := "tcp(localhost:3306)"
	DBNAME := "otokake"
	CONNECT := USER + ":" + PASS + "@" + PROTOCOL + "/" + DBNAME
	db, err := gorm.Open(DBMS, CONNECT)

	if err != nil {
		panic(err.Error())
	}
	fmt.Println("db connected: ", &db)
	return db
}

func main() {
	db := gormConnect()
	defer db.Close()

	db.LogMode(true) //gormのログモード
	r := gin.Default()

	r.GET("/musics/:id", func(c *gin.Context) {
		return_music := []musics{}
		id := c.Param("id")
		db.Limit(5).Where("music_ID >= ?", id).Find(&return_music)
		c.JSON(http.StatusOK, return_music)
		//idに入力された値から5件取得する。本番環境では10件とかに変える?
	})

	r.Run(":8080")
}
