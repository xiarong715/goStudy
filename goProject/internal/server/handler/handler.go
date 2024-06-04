package handler

import (
	"fmt"
	"goproject/internal/conf"
	"log"
)

func init() {
	conf.RunAfterConfInit(func() error {
		fmt.Println("init handler")
		return nil
	})
}

func CreateHandler() {
	log.Println("handler create handler")
}
