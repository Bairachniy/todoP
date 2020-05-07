package main

import (
	"fmt"
	todoServ "golang-book/todoP/services"
)

func main(){
	todoshka:=new(todoServ.Todo)
	todoshka.Create("First todo")
	todoshka.Create("Second todo")
	//fmt.Println(todoshka.GetAll())
	//todoshka.ClearAllDb()
	fmt.Println(todoshka.GetAll())
}
