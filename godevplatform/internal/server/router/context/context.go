package context

import (
	"net/http"

	json "github.com/json-iterator/go"
)

type Context struct {
	W     http.ResponseWriter
	R     *http.Request
	Param json.Any
}

type resp struct {
	ErrCode int    `json:"err_code"`
	Reply   string `json:"reply"`
	Result  any    `json:"result"`
}

func (c *Context) ResponseJson(statuscode int, jsonValue any) error {
	json, err := json.Marshal(jsonValue)
	if err != nil {
		return err
	}
	c.W.WriteHeader(statuscode)
	c.W.Header().Set("Content-Type", "application/json")
	_, err = c.W.Write(json)
	if err != nil {
		return err
	}
	return nil
}

func (c *Context) ResponseJsonOK(data any) error {
	return c.ResponseJson(http.StatusOK, resp{
		ErrCode: 0,
		Reply:   "ok",
		Result:  data,
	})
}

func (c *Context) ResponseJsonFailed(errcode int, data any) error {
	return c.ResponseJson(http.StatusOK, resp{
		ErrCode: errcode,
		Reply:   "failed",
		Result:  data,
	})
}
