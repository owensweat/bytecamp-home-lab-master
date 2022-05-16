package main

import (
	"Go_test5/homework/model"
	"Go_test5/homework/query"
	"context"
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"math/rand"
	"time"
)


var db, _ = gorm.Open(mysql.Open("root:root@tcp(127.0.0.1:3306)/springsecurity?charset=utf8mb4&parseTime=True"))
var ctx context.Context

func main() {

	for i:=1; i <= 100 ;i++ {
		var maxNum = 10
		rand.Seed(time.Now().UnixNano())
		uuid := rand.Intn(maxNum)

		id := fmt.Sprintf("%v", uuid)

		u := query.Use(db).People

		count,err :=u.WithContext(ctx).Where(u.UUID.Eq(id)).First()

		if err!=nil{}

		if count==nil{
			u.WithContext(ctx).Create(&model.People{id, "owen", 0, 0})
		}else{
			u.WithContext(ctx).Where(u.UUID.Eq(id)).Update(u.Version,u.Version.Add(1))
		}

	}

	/*for i:=1; i <= 10 ;i++ {

		var maxNum = 10
		rand.Seed(time.Now().UnixNano())
		uuid := rand.Intn(maxNum)

		id := fmt.Sprintf("%v", uuid)

		fmt.Println(id)

		var user User

		user.UUID = id
		user.Version = 0

		//初始化
		db.Debug().Clauses(clause.OnConflict{
			Columns:[]clause.Column{{Name:"uuid"}},
			DoUpdates:clause.Assignments(map[string]interface{}{"version":gorm.Expr("version + ?",1)}),
		}).Create(&user)

	}*/





}

