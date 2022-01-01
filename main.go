package main

import (
	"fmt"
	"golang_todo_app/app/models"
)

func main() {
	fmt.Println(models.Db)

	// u := &models.User{
	// 	Name:     "test",
	// 	Email:    "test@example.com",
	// 	PassWord: "testtest",
	// }
	// fmt.Println(u)

	// u.CreateUser()

	// user, _ := models.GetUser(2)

	// fmt.Println(u)

	// u.Name = "test2"
	// u.Email = "test2@example.com"
	// u.UpdateUser()
	// u, _ = models.GetUser(1)
	// fmt.Println(u)

	// user.CreateTodo("First Todo")

	t, _ := models.GetTodo(1)
	fmt.Println(t)
}
