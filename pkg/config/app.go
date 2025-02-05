package config

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var (
	db *gorm.DB // This will hold the database connection object
)

func ConnectORMToDataBase() {
	// Corrected DSN connection string
	dsn := "akhil:Axlesharma@12@tcp(127.0.0.1:3306)/simplerest?charset=utf8&parseTime=True&loc=Local"
	// Open the connection with GORM and MySQL driver
	d, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err) // If connection fails, panic and exit
	}
	db = d // Assign the database connection to the global variable
}

func GetDB() *gorm.DB {
	return db // Return the database connection to use elsewhere
}
