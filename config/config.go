package config

import (
	"fmt"
	"log"
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

	sqlDB, err := DB.DB()
	if err != nil {
		log.Fatalf("Error getting SQL DB instance from Gorm: %v", err)
	}

	sqlDB.SetMaxOpenConns(100)  // Maximum number of open connections to the database
	sqlDB.SetMaxIdleConns(10)   // Maximum number of connections in the idle connection pool
	sqlDB.SetConnMaxLifetime(0) // Maximum amount of time a connection may be reused

	return DB
}

// CloseDB closes the database connection.
func CloseDB() {
	sqlDB, err := DB.DB()
	if err != nil {
		log.Printf("Error getting SQL DB instance from Gorm for closing: %v", err)
		return
	}
	err = sqlDB.Close()
	if err != nil {
		log.Printf("Error closing the database connection: %v", err)
	}
}
