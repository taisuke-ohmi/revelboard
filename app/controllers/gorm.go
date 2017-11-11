package controllers

import (
	"log"
	"myapp/app/models"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	"github.com/revel/revel"
)

var DB *gorm.DB

func InitDB() {
	dbinfo, ok := revel.Config.String("db")
	if !ok {
		log.Fatalln("db.infoは設定されていません")
	}

	db, err := gorm.Open("mysql", dbinfo)
	if err != nil {
		log.Panicf("Failed to connect to database: %v\n", err)
	}

	db.AutoMigrate(&models.Comment{}) // ここで table の作成を行う
	DB = db
}
