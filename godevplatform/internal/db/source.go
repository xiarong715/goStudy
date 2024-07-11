package db

type Source struct {
	ID         uint64 `json:"id"`
	SourceName string `json:"sourcename"`
}

func init() {
	addModel(&Source{})
}
