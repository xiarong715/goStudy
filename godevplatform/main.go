package main

import (
	"godevplatform/internal/conf"
	"godevplatform/internal/db"
	"log"
	"os"
	"path"
)

func init() {
	log.Println("main init")
	dir, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}
	if err = os.Setenv("DATA_PATH", path.Join(dir, "build")); err != nil {
		log.Fatal(err)
	}

	if err := conf.InitConf(); err != nil {
		log.Fatal("initconf:", err)
	}
}

func main() {
	log.Println("devplatform")
	db.ItDB()
}
