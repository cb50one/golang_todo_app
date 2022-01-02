package main

import (
	"fmt"
	"golang_todo_app/app/controllers"
	"golang_todo_app/app/models"
)

func main() {
	fmt.Println(models.Db)
	controllers.StartMainServer()
}
