package main

import (
	"fmt"

	"github.com/jinzhu/gorm"
)

var err error

func main() {

	config.DB, err = gorm.Open("mysql", "root:@tcp(127.0.0.1:3306)/testinger?charset=utf8&parseTime=True&loc=Local")

	if err != nil {
		fmt.Println("statuse: ", err)
	}
	defer config.DB.Close()
	config.DB.AutoMigrate(&models.Book{})

	r := routers.SetupRouter()
	// running
	r.Run()
}
