package main

import (
	"fmt"
	"golang_todo_app/app/models"
)

func main() {
	// fmt.Println(config.Config.Port)
	// fmt.Println(config.Config.SQLDriver)
	// fmt.Println(config.Config.DbName)
	// fmt.Println(config.Config.LogFile)

	// log.Println("test")

	fmt.Println(models.Db)

	u := &models.User{
		Name:     "test",
		Email:    "test@example.com",
		PassWord: "testtest",
	}
	fmt.Println(u)

	u.CreateUser()
}
