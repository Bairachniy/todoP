package main

import (
	"fmt"
	"golang-book/todoP/services"
)

func main() {
	var toLoc services.Repo = new(services.ToLocal)
	var toBase services.Repo = new(services.ToBase)
	todoshka := services.NewTodos(toLoc)
	todoshka.Create("Hello")
	todoshka.Create("World")
	todoBase := services.NewTodos(toBase)
	todoBase.Create("World")
	fmt.Println(todoshka.GetAll(), todoBase.GetAll())
}
