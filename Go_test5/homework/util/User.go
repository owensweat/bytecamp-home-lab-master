package main

import (
	"fmt"
)

type User struct {
	UUID string
	Name string
	Age int
	Version int
}


func main() {
	//创建表
	db.AutoMigrate(&User{})
	result :=db.Migrator().CreateTable(&User{})

	fmt.Println(result)
}
