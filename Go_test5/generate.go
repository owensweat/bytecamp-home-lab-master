package main

import (
	"Go_test5/homework/model"
	"gorm.io/driver/mysql"
	"gorm.io/gen"
	"gorm.io/gorm"
)

var db, _ = gorm.Open(mysql.Open("root:root@tcp(127.0.0.1:3306)/springsecurity?charset=utf8mb4&parseTime=True"))

// generate code
func main() {
	// 指定生成代码的具体(相对)目录，默认为：./query
	// 默认情况下需要使用WithContext之后才可以查询，但可以通过设置gen.WithoutContext避免这个操作
	g := gen.NewGenerator(gen.Config{
		// 最终package不能设置为model，在有数据库表同步的情况下会产生冲突，若一定要使用可以单独指定model package的新名字
		OutPath: "./homework/query",
		Mode:    gen.WithoutContext, // 默认情况下会跟随OutPath参数，在同目录下生成model目录
		/* Mode: gen.WithoutContext,*/
	})

	// 复用工程原本使用的SQL连接配置db(*gorm.DB)
	// 非必需，但如果需要复用连接时的gorm.Config或需要连接数据库同步表信息则必须设置

	db, _ := gorm.Open(mysql.Open("root:root@tcp(127.0.0.1:3306)/springsecurity?charset=utf8mb4&parseTime=True"))

	g.UseDB(db)
	// 所有需要实现查询方法的结构体
	//g.ApplyBasic(g.GenerateModelAs("users", "People"))


	g.ApplyInterface(func(method model.Method) {},g.GenerateModelAs("users","People"))

	//生成全部表
	//g.GenerateAllTable()

	// 执行并生成代码
	g.Execute()


}
