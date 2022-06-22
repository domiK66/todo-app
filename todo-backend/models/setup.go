package models

import (
	gorm "github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

var DB *gorm.DB

func ConnectDatabase() {
	database, err := gorm.Open("sqlite3", "test.db")
  
	if err != nil {
	  panic("Failed to connect to database!")
	}
  
	database.AutoMigrate(&Task{})
	database.AutoMigrate(&Priority{})
	database.AutoMigrate(&Todolist{})
	DB = database
}




//https://sourceforge.net/projects/mingw-w64/
//https://code.visualstudio.com/docs/cpp/config-mingw