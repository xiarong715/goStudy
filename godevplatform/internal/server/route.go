package server

import (
	"errors"
	"godevplatform/internal/conf"
	"godevplatform/internal/server/router/context"
	"log"
	"net/http"
	"regexp"
	"strings"
)

func init() {
	log.Println("route")
}

type Entrance struct{}

func NewEntrance() *Entrance {
	return &Entrance{}
}
func (e *Entrance) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	log.Println("raw path:", r.URL.Path)
	ctx := &context.Context{W: w, R: r}

	path := strings.TrimPrefix(r.URL.Path, conf.Conf.PathPrefix)
	log.Println("path:", path)

	// path
	pathRoute, err := getPathRoute(path)
	if err != nil {
		ctx.ResponseJsonFailed(1, err.Error())
		return
	}
	if pathRoute.method != r.Method && pathRoute.method != "*" {
		ctx.ResponseJsonFailed(1, "method not allowed")
		return
	}

	// action
	pathRoute.f(ctx)
}

type PathRoute struct {
	reg    *regexp.Regexp
	method string
	f      func(ctx *context.Context)
}

var PathRoutes []PathRoute = []PathRoute{
	{reg: regexp.MustCompile("^/api/v1/server$"), method: "*", f: ServerHandle},
}

func getPathRoute(path string) (PathRoute, error) {
	for _, pathroute := range PathRoutes {
		res := pathroute.reg.FindStringSubmatch(path)
		if len(res) != 0 {
			return pathroute, nil
		}
	}
	return PathRoute{}, errors.New("path route not found")
}

func ServerHandle(ctx *context.Context) {
	// action
}
