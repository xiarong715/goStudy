package main

import (
	"log"

	"github.com/astaxie/beego/server/web"
)

func main() {
	log.Printf("app with beego")
	web.Run()
}
