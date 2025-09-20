package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type Context struct {
	Response http.ResponseWriter
	Request  *http.Request
}

func (c *Context) JSON(code int, obj interface{}) {
	c.Response.Header().Set("Content-Type", "application/json")
	c.Response.WriteHeader(code)
	json.NewEncoder(c.Response).Encode(obj)
}

type HandlerFunc func(*Context)

type Engine struct {
	routes map[string]map[string]HandlerFunc
}

func New() *Engine {
	return &Engine{
		routes: make(map[string]map[string]HandlerFunc),
	}
}

func (e *Engine) addRoute(method, path string, handler HandlerFunc) {
	if _, ok := e.routes[path]; !ok {
		e.routes[path] = make(map[string]HandlerFunc)
	}
	e.routes[path][method] = handler
}

func (e *Engine) GET(path string, handler HandlerFunc) {
	e.addRoute("GET", path, handler)
}

func (e *Engine) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	path := r.URL.Path
	if handlers, ok := e.routes[path]; ok {
		if handler, ok := handlers[r.Method]; ok {
			handler(&Context{Response: w, Request: r})
			return
		}
	}
	http.NotFound(w, r)
}

func (e *Engine) Run(addr string) error {
	fmt.Printf("Server starting on %s\n", addr)
	return http.ListenAndServe(addr, e)
}

type H map[string]interface{}

func main() {
	r := New()
	r.GET("/ping", func(c *Context) {
		c.JSON(200, H{
			"message": "pong",
		})
	})
	r.Run(":8080")
}