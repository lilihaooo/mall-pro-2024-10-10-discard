package test

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
	"testing"
)

func TestGorm(t *testing.T) {
	dsn := "root:a111111@tcp(47.109.31.106:3306)/redu?charset=utf8&parseTime=True&loc=UTC"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("failed to connect database")
	}
	fmt.Println(db)
	//var categoryIDs []uint
	//var categoryName []string
	//err := db.Model(product.Category{}).Where("level = 3").Pluck("id", &categoryIDs).Error

}
