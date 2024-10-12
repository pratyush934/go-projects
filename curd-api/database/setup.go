package database

import (
	"fmt"
	"github.com/pratyush934/go-projects/curd-api/models"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDataBase() {
	url := "root:Pratyush@123@/gotesting?charset=utf8&parseTime=True&loc=Local"
	database, err := gorm.Open(mysql.Open(url), &gorm.Config{})

	if err != nil {
		fmt.Println(err)
		panic(`Failed to connect with database Please look at the file database/setup.go`)
	} else {
		fmt.Println("Yes We have done this and successfully setup the gotesting database")
	}

	if err := database.AutoMigrate(&models.Post{}); err != nil {
		fmt.Println(err)
		panic("Failed to AutoMigrate please look at the issue")
	}

	DB = database

}
