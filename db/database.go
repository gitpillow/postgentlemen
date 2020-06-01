package db

import (
	"github.com/gitpillow/postgentlemen/utils"
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
		utils.Log.Infof("create db: %v", DBName)
		file, err := os.Create(DBName)
		if err != nil {
			utils.Log.Fatalf("can not create sqlite db file: %v", DBName)
		}
		file.Close()
	}
}

func RemoveDB() {
	utils.Log.Infof("remove db: %v\n", DBName)
	err := os.RemoveAll(DBName)
	if err != nil {
		utils.Log.Fatalf("remove db error: %v\n", err)
	}
}

func GetDB() *gorm.DB {
	if DB != nil {
		return DB
	}

	db, err := gorm.Open("sqlite3", DBName)
	if err != nil {
		utils.Log.Fatalf("can not connect to sqlite db: %v\n", DBName)
	}

	DB = db
	utils.Log.Infof("get db connection: %v", DBName)
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
	utils.Log.Infof("start to migrate %v models", len(models))
	for _, model := range models {
		utils.Log.Infof("migrate model: %v", reflect.TypeOf(model).Elem().Name())
		GetDB().AutoMigrate(model)
	}
	utils.Log.Infof("migrate models done")
}
