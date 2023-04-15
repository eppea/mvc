package config

import (
	"mvc/model"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func InitDB() {
	dsn := "root:EANHHUFWsX2ocayI4WXW@tcp(containers-us-west-143.railway.app:6475)/railway?charset=utf8mb4&parseTime=True&loc=Local"
	var err error
	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("connection DB")
	}
	migration()
}

func migration() {
	DB.AutoMigrate(&model.Transaction{})
}
