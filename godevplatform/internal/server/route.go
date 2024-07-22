package server

import (
	"errors"
	"godevplatform/internal/conf"
	"godevplatform/internal/server/handler/api/v1/server"
	"godevplatform/internal/server/router/context"
	"io"
	"log"
	"net/http"
	"regexp"
	"strings"

	json "github.com/json-iterator/go"
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

var pathRoutes []PathRoute = []PathRoute{
	{reg: regexp.MustCompile("^/api/v1/server$"), method: "*", f: ServerHandle},
}

func getPathRoute(path string) (PathRoute, error) {
	for _, pathroute := range pathRoutes {
		res := pathroute.reg.FindStringSubmatch(path)
		if len(res) != 0 {
			return pathroute, nil
		}
	}
	return PathRoute{}, errors.New("path route not found")
}

type ActionRoute struct {
	action string
	mehod  string
	f      func(ctx *context.Context)
}

var actionRoutes = map[string]ActionRoute{
	"create_user":   {"create_user", "post", server.CreateUser},
	"clear_session": {"clear_session", "post", server.ClearSession},
}

type RequestParam struct {
	Action string   `json:"action"`
	Param  json.Any `json:"param"`
}

func ServerHandle(ctx *context.Context) {
	// action
	body, err := io.ReadAll(ctx.R.Body)
	if err != nil {
		ctx.ResponseJsonFailed(1, "read body:"+err.Error())
		return
	}
	param := RequestParam{}
	if err = json.Unmarshal(body, &param); err != nil {
		ctx.ResponseJsonFailed(1, "unmarshal body:"+err.Error())
		return
	}
	ctx.Param = param.Param
	actionroute, ok := actionRoutes[param.Action]
	if !ok {
		ctx.ResponseJsonFailed(1, "action not found")
		return
	}
	actionroute.f(ctx)
}
