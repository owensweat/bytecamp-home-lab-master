package main
/**
gen
 */
import (
	"Go_test5/homework/query"
	"context"
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)
var db, _ = gorm.Open(mysql.Open("root:root@tcp(127.0.0.1:3306)/springsecurity?charset=utf8mb4&parseTime=True"))
var ctx context.Context


func main() {

	// 测试
	count, _ := query.Use(db).People.GetMaxVersionCount()

	fmt.Println(count)







}
