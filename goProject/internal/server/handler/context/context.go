package context

import (
	"encoding/json"
	"net/http"
)

type Context struct {
	W http.ResponseWriter
	R *http.Request
}

func (c *Context) ResonseJson(statuscode int, data interface{}) error {
	res, err := json.Marshal(data)
	if err != nil {
		return err
	}
	c.W.WriteHeader(statuscode)
	c.W.Header().Set("Content-Type", "application/json")
	_, err = c.W.Write(res)
	return err
}

func (c *Context) ResponseJsonOK(data interface{}) error {
	return c.ResonseJson(http.StatusOK, map[string]interface{}{
		"err_code": 0,
		"reply":    "ok",
		"result":   data,
	})
}

func (c *Context) ResonseJsonFailed(errcode uint8, data interface{}) error {
	return c.ResonseJson(http.StatusOK, map[string]interface{}{
		"err_code": errcode,
		"reply":    "failed",
		"result":   data,
	})
}

func (c *Context) ResponseString(statuscode int, data string) error {
	c.W.WriteHeader(statuscode)
	_, err := c.W.Write([]byte(data))
	return err
}

// 重复出现2遍，尝试提取出函数
