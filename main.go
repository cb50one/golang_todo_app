package main

import (
	"fmt"
	"golang_todo_app/app/models"
)

func main() {
	fmt.Println(models.Db)

	// u := &models.User{
	// 	Name:     "test2",
	// 	Email:    "test2@example.com",
	// 	PassWord: "testtest",
	// }
	// u.CreateUser()

	// user, _ := models.GetUser(2)

	// u.Name = "test2"
	// u.Email = "test2@example.com"
	// u.UpdateUser()
	// u, _ = models.GetUser(1)
	// fmt.Println(u)

	// user2.CreateTodo("Third Todo")

	// todos, _ := user.GetTodosByUser()
	// for _, v := range todos {
	// 	fmt.Println(v)
	// }
	todo, _ := models.GetTodo(1)
	todo.Content = "Updated First Content"
	todo.UpdateTodo()
	// fmt.Println(t)

}
