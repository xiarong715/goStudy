package db

import (
	"fmt"
	"goproject/internal/conf"
	"log"
)

func init() {
	conf.RunAfterConfInit(func() error {
		fmt.Println("init db")
		return nil
	})
}

func CreateUser() {
	log.Println("db create user")
}
