package domain

import "github.com/jinzhu/gorm"

var DB *gorm.DB

func InitDb() {
	var err error
	DB, err = gorm.Open("sqlite3", "test.db")
	if err != nil {
		panic("failed to connect database")
	}
}

