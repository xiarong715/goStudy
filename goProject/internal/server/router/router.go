package router

import (
	"fmt"
	"goproject/internal/conf"
	"log"
)

func init() {
	conf.RunAfterConfInit(func() error {
		fmt.Println("init router")
		return nil
	})
}

func CreateRouter() {
	log.Println("router create router")
}
