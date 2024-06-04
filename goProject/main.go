package main

import (
	"goproject/internal/conf"
	"goproject/internal/db"
	"goproject/internal/server/handler"
	"goproject/internal/server/handler/context"
	"goproject/internal/server/router"
	"log"
	"net/http"
)

func init() {
	if err := conf.InitConfig(); err != nil {
		log.Fatal(err)
	}
}

// 包中init的执行顺序是包引入的顺序；
// RunAfterConfInit 可控制加入的函数在配置初始化之后执行

// 引用其他包，又希望其他包的初始化与本包的初始化有先后顺序，就可用在某操作前或后执行一些函数；
func main() {
	db.CreateUser()
	handler.CreateHandler()
	router.CreateRouter()

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		ctx := &context.Context{W: w, R: r}
		ctx.ResponseJsonOK(map[string]interface{}{
			"token": "1234",
		})
	})
	http.HandleFunc("/msg", func(w http.ResponseWriter, r *http.Request) {
		ctx := &context.Context{W: w, R: r}
		ctx.ResponseString(200, "hello super man")
	})
	http.HandleFunc("/err", func(w http.ResponseWriter, r *http.Request) {
		ctx := &context.Context{W: w, R: r}
		ctx.ResonseJsonFailed(200, "something wrong")
	})
	http.ListenAndServe(":8081", nil)
}
