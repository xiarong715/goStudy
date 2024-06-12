package main

import (
	"log"
	"strings"
)

var path = "/proxy/api/v1/backend"

func main() {
	paths := strings.Split(path, "/proxy")
	log.Println(paths, len(paths))
}
