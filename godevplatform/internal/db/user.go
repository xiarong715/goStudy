package db

import "log"

type User struct {
	ID       uint64 `json:"id"`
	Username string `json:"username"`
}

func init() {
	log.Println("db user init")
	addModel(&User{})
}
