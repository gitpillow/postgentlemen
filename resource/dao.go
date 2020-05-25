package resource

import (
	"fmt"
	"github.com/jinzhu/gorm"
	"os"
)

const DBName = "pgmen.db"

var DB *gorm.DB

func CreateDB() {
	if !CheckDBExists() {
		file, err := os.Create(DBName)
		if err != nil {
			fmt.Errorf("can not create sqlite db file: %v", DBName)
		}
		file.Close()
	}
}

func GetDB() *gorm.DB {
	if DB != nil {
		return DB
	}

	db, err := gorm.Open("sqlite3", DBName)
	if err != nil {
		fmt.Errorf("can not connect to sqlite db: %v", DBName)
	}
	DB = db
	return DB
}

func CheckDBExists() bool {
	_, err := os.Stat(DBName)
	if err == nil {
		return true
	}
	return false
}
