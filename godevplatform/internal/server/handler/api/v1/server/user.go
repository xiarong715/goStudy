package server

import (
	"godevplatform/internal/server/router/context"
)

func CreateUser(ctx *context.Context) {
	SetSession(ctx)
	param := struct {
		Name string `json:"name"`
	}{}
	ctx.Param.ToVal(&param)
	ctx.ResponseJsonOK(map[string]interface{}{
		"name": param.Name,
	})
}
