package gee

import "net/http"

// H interface 类似java 中的Object
type H map[string]interface{}

type Context struct {
	Writer    http.ResponseWriter
	Req       *http.Request
	Path      string
	Method    string
	StateCode int
}

func newContext(w http.ResponseWriter, r *http.Request) *Context {
	return &Context{
		Writer: w,
		Req:    r,
		Path:   r.URL.Path,
		Method: r.Method,
	}
}

// PostForm Context的结构体方法
func (c *Context) PostForm(key string) string {
	return c.Req.URL.Query().Get(key)
}
