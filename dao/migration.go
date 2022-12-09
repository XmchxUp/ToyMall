package dao

import (
	"fmt"
	"os"
	"xm-mall/model"
)

func Migration() {
	//自动迁移模式
	err := _db.Set("gorm:table_options", "charset=utf8mb4").
		AutoMigrate(&model.Address{},
			&model.Admin{},
			&model.Carousel{},
			&model.Cart{},
			&model.Category{},
			&model.Favorite{},
			&model.Notice{},
			&model.Order{},
			&model.ProductImg{},
			&model.Product{},
			&model.User{})
	if err != nil {
		fmt.Println("register table fail")
		os.Exit(0)
	}
	fmt.Println("register table success")
}
