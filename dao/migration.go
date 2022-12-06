package dao

import (
	"fmt"
	"os"
)

func Migration() {
	//自动迁移模式
	err := _db.Set("gorm:table_options", "charset=utf8mb4").AutoMigrate()
	if err != nil {
		fmt.Println("register table fail")
		os.Exit(0)
	}
	fmt.Println("register table success")
}
