package db

import (
	"github.com/gitpillow/postgentlemen/utils"
	"github.com/jinzhu/gorm"
	_ "github.com/mattn/go-sqlite3"
	"os"
	"reflect"
)

// sqlite file name for rest resource restore
var DBName = "pgmen.sqlite"

// gorm db access
var DB *gorm.DB

// gorm entity models
var models []interface{}

// CreateDB create sqlite file if not exist
func CreateDB() {
	if !CheckDBExists() {
		utils.Log.Infof("create db: %v", DBName)
		f, err := os.Create(DBName)
		if err != nil {
			utils.Log.Fatalf("can not create sqlite db file: %v", DBName)
		}
		f.Close()
	}
}

// RemoveDB delete sqlite file
func RemoveDB() {
	utils.Log.Infof("remove db: %v\n", DBName)
	err := os.RemoveAll(DBName)
	if err != nil {
		utils.Log.Fatalf("remove db error: %v\n", err)
	}
}

// CheckDBExists return if the sqlite db file exists
func CheckDBExists() bool {
	_, err := os.Stat(DBName)
	if err == nil {
		return true
	}
	return false
}

// GetDB return a gorm db access
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


// Register receive entity model from other package init() func
func Register(model interface{}) {
	models = append(models, model)
}

// Migrate create database structure by registered entity models
func Migrate() {
	utils.Log.Infof("start to migrate %v models", len(models))
	for _, model := range models {
		utils.Log.Infof("migrate model: %v", reflect.TypeOf(model).Elem().Name())
		GetDB().AutoMigrate(model)
	}
	utils.Log.Infof("migrate models done")
}