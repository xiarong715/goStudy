package server

import (
	"godevplatform/internal/conf"
	"godevplatform/internal/server/router/context"
	"net/http"
	"time"
)

func SetSession(ctx *context.Context) {
	http.SetCookie(ctx.W, &http.Cookie{
		Name:     conf.Conf.Cookie.CookieName,
		Value:    conf.Conf.Cookie.CookieValue,
		Path:     conf.Conf.Cookie.CookiePath,
		Domain:   conf.Conf.Domain,
		Expires:  time.Now().Add(time.Minute * 10),
		Secure:   true,
		HttpOnly: true,
	})
}

func ClearSession(ctx *context.Context) {
	http.SetCookie(ctx.W, &http.Cookie{
		Name:     conf.Conf.Cookie.CookieName,
		Path:     conf.Conf.Cookie.CookiePath,
		Domain:   conf.Conf.Domain,
		MaxAge:   -1,
		Secure:   true,
		HttpOnly: true,
	})
}
