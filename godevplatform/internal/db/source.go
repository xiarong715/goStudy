package db

import "log"

type Source struct {
	ID         uint64 `json:"id"`
	SourceName string `json:"sourcename"`
}

func init() {
	log.Println("db source init")
	addModel(&Source{})
}
