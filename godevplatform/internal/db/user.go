package db

type User struct {
	ID       uint64 `json:"id"`
	Username string `json:"username"`
}

func init() {
	addModel(&User{})
}
