package config

import (
	"fmt"
	"os"

	_ "github.com/joho/godotenv"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

// func LoadEnv() {
// 	err := godotenv.Load()
// 	if err != nil {
// 		panic("failed to load file")
// 	}
// }

type DBConfig struct {
	Username string
	Password string
	Host     string
	Port     string
	Name     string
}

func ConnectToDB() *gorm.DB {
	var dbConfig DBConfig = DBConfig{
		Username: os.Getenv("DB_USER"),
		Password: os.Getenv("DB_PASSWORD"),
		Host:     os.Getenv("DB_HOST"),
		Port:     os.Getenv("DB_PORT"),
		Name:     os.Getenv("DB_NAME"),
	}
	fmt.Println("dbconfig : ", dbConfig)
	fmt.Println("test")

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		dbConfig.Username,
		dbConfig.Password,
		dbConfig.Host,
		dbConfig.Port,
		dbConfig.Name)

	var err error
	fmt.Println("dsn : ", dsn)
	DB, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	// DB, err = gorm.Open("mysql", dsn)
	if err != nil {
		fmt.Println("failed")
		panic("Database Connection Error")
	}
	fmt.Println("Success")
	return DB
}
