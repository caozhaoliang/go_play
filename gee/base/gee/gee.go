package gee

import (
	"net/http"
	"strings"
)

type HandlerFunc func(*Context)
type Engine struct {
	*RouterGroup
	groups []*RouterGroup
	router *router
}

func New() *Engine {
	e := &Engine{router: newRouter()}
	e.RouterGroup = &RouterGroup{engine: e}
	e.groups = []*RouterGroup{e.RouterGroup}
	return e
}
func (e *Engine) addRouter(method string, pattern string, handler HandlerFunc) {
	e.router.addRoute(method, pattern, handler)
}
func (e *Engine) Get(pattern string, handler HandlerFunc) {
	e.addRouter("GET", pattern, handler)
}
func (e *Engine) POST(pattern string, handler HandlerFunc) {
	e.addRouter("POST", pattern, handler)
}
func (e *Engine) Run(addr string) error {
	return http.ListenAndServe(addr, e)
}

// 引入 content
func (e *Engine) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	var middlewares []HandlerFunc
	for _, group := range e.groups {
		if strings.HasPrefix(req.URL.Path, group.prefix) {
			middlewares = append(middlewares, group.middlewares...)
		}
	}
	c := newContext(w, req)
	c.handlers = middlewares
	e.router.handle(c)
}
