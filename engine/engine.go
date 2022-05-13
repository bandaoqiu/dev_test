package engine

import (
	"encoding/json"
	"net/http"
)

type StatusCode int

var (
	Success    StatusCode = 200
	NoValidate StatusCode = 400
	ServerErr  StatusCode = 500
)

type Context struct {
	Writer  http.ResponseWriter
	Request *http.Request
}
type H map[string]interface{}

func (ctx *Context) Query(param string) string {
	return ctx.Request.URL.Query().Get(param)
}
func (ctx *Context) PostForm(param string) string {
	return ctx.Request.PostFormValue(param)
}
func (ctx *Context) Write(code StatusCode, msg []byte) {
	ctx.Writer.WriteHeader(int(code))
	ctx.Writer.Write(msg)
}
func (ctx *Context) Json(data interface{}) {
	ctx.Writer.WriteHeader(int(Success))
	ctx.Writer.Header().Set("Content-Type", "application/json")
	bytes, err := json.Marshal(data)
	if err != nil {
		ctx.Writer.Write([]byte(err.Error()))
	} else {
		ctx.Writer.Write(bytes)
	}
}

type HandlerFunc func(ctx *Context)
type Engine struct {
	router map[string]HandlerFunc
}

func New() *Engine {
	return &Engine{
		router: make(map[string]HandlerFunc),
	}
}
func (t *Engine) addRouter(method, pattern string, hdl HandlerFunc) {
	key := method + "-" + pattern
	t.router[key] = hdl
}
func (t *Engine) Get(pattern string, hdl HandlerFunc) {
	t.addRouter("GET", pattern, hdl)
}
func (t *Engine) Post(pattern string, hdl HandlerFunc) {
	t.addRouter("POST", pattern, hdl)
}
func (t *Engine) Run(addr string) {
	http.ListenAndServe(addr, t)
}
func (t *Engine) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	key := req.Method + "-" + req.URL.Path
	if hdl, ok := t.router[key]; ok {
		hdl(&Context{
			Writer:  w,
			Request: req,
		})
	} else {
		w.WriteHeader(404)
		w.Write([]byte("router not found"))
	}
}
