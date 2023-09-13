package main

import (
	"fmt"

	"github.com/ekonuma/webtodo/db"
	"github.com/ekonuma/webtodo/model"
)

func main(){
	dbConn := db.NewDB()
	defer fmt.Println("Successfully Migrated")
	defer db.CloseDB(dbConn)
	dbConn.AutoMigrate(&model.User{}, &model.Task{})
}