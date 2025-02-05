package config

import (
    "gorm.io/driver/mysql"
    "gorm.io/gorm"
)


// gorm.Open(): Ye function gorm ORM (Object Relatinal Mapping) ko database se connect karne ke liye use hota hai.
// DSN (Data Source Name) ek string hota hai jo ek database ke saath connection establish karne ke liye zaruri information ko specify karta hai
// mysql.Open(dsn): Ye gorm ko MySQL driver ke saath connect karta hai
// &gorm.Config{}: Ye optional configuration object hai jo gorm ke default behavior ko customize karne ke liye use kiya jata hai
// MySQL driver ek library hai jo Go programming language ko MySQL database ke saath interact karne ki capability provide karta hai

var (
	db *gorm.DB    //
)

func ConnectORMToDataBase(){
	
	// Yeh connection string database ko identify karne ka kaam karti hai 
	// aur uske saath connection establish karne ke liye zaroori settings ko define karti hai.
	dsn := "akhil:Axlesharma@12@/simplerest?charset=utf8&parseTime=True&loc=Local"
 // Update your connection string as needed
	d, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err!=nil{
		panic(err)
	}
	db = d;
}

func GetDB() *gorm.DB {
	return db
}