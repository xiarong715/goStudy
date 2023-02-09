package singleton

import "sync"

type singleton struct{}

var instance *singleton
var once sync.Once

// sync.Once 实现单实例
func GetInstance() *singleton {
	once.Do(func() {
		instance = &singleton{}
	})
	return instance
}
