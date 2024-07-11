package db

import (
	"godevplatform/internal/conf"
	"log"
	"path"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var model []interface{}
var runnerBeforeDBInit []func() error
var db *gorm.DB

func init() {
	log.Println("db gorm")
	conf.RunAfterConfInit(initDB)
}

func initDB() error {
	var err error
	log.Println("before initDB")
	for _, runner := range runnerBeforeDBInit {
		if err := runner(); err != nil {
			return err
		}
	}

	log.Println("initDB")
	dbpath := path.Join(conf.Conf.DataPath, "devplatform.db")
	db, err = gorm.Open(sqlite.Open(dbpath), &gorm.Config{})
	if err != nil {
		return err
	}
	if err = db.AutoMigrate(model...); err != nil {
		return err
	}

	return nil
}

func addModel(m interface{}) {
	model = append(model, m)
}

func RunBeforeDBInit(f func() error) {
	runnerBeforeDBInit = append(runnerBeforeDBInit, f)
}

func ItDB() {
	log.Println("ItDB")
}
