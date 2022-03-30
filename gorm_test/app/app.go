package app

import (
	"gorm.io/gorm"
)

type Product struct {
	gorm.Model
	Code  string
	Price uint
}

func Start(db *gorm.DB) {

	db.AutoMigrate(&Product{})

	db.Create(&Product{Code: "D42", Price: 100})

	var product Product
	db.First(&product, 1) //pk기준으로 product찾기
	db.First(&product, "code = ?", "D42")

	db.Model(&product).Update("Price", 200)
	db.Model(&product).Updates(Product{Price: 200, Code: "F42"})

	db.Delete(&product, 1)
}
