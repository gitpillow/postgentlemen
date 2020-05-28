package db

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/mattn/go-sqlite3"
	"os"
	"reflect"
)

var DBName = "pgmen.sqlite"
var DB *gorm.DB
var models []interface{}

func CreateDB() {
	if !CheckDBExists() {
		file, err := os.Create(DBName)
		if err != nil {
			fmt.Errorf("can not create sqlite db file: %v", DBName)
			os.Exit(1)
		}
		file.Close()
	}
}

func RemoveDB() {
	fmt.Printf("remove db: %v\n", DBName)
	err := os.RemoveAll(DBName)
	if err != nil {
		fmt.Errorf("remove db error: %v\n", err)
		os.Exit(1)
	}
}

func GetDB() *gorm.DB {
	if DB != nil {
		return DB
	}

	db, err := gorm.Open("sqlite3", DBName)
	if err != nil {
		fmt.Errorf("can not connect to sqlite db: %v\n", DBName)
		os.Exit(1)
	}

	DB = db
	fmt.Printf("get db connection: %v\n", DBName)
	return DB
}

func CheckDBExists() bool {
	_, err := os.Stat(DBName)
	if err == nil {
		return true
	}
	return false
}

func Register(model interface{}) {
	models = append(models, model)
}

func Migrate() {
	fmt.Printf("start to migrate %v models\n", len(models))
	for _, model := range models {
		fmt.Printf("migrate model: %v\n", reflect.TypeOf(model).Elem().Name())
		GetDB().AutoMigrate(model)
	}
	fmt.Printf("migrate models done\n")
}
