package main

import (
	"godevplatform/internal/conf"
	"log"
)

func init() {
	if err := conf.InitConf(); err != nil {
		log.Fatal("initconf:", err)
	}
}

func main() {
	log.Println("devplatform")
}
