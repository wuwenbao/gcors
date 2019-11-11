package gocors

import (
	"net/http"
)

//cors 跨域
type cors struct {
	nextHandler  http.Handler
	allowOrigin  string
	allowMethods string
	allowHeaders string
}

func New(handler http.Handler) *cors {
	return &cors{
		nextHandler:  handler,
		allowOrigin:  "*",
		allowMethods: "*",
		allowHeaders: "",
	}
}

func (c *cors) Allow(origin, methods, headers string) {
	c.allowOrigin = origin
	c.allowMethods = methods
	c.allowHeaders = headers
}

func (c *cors) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", c.allowOrigin)
	w.Header().Set("Access-Control-Allow-Methods", c.allowMethods)
	w.Header().Set("Access-Control-Allow-Headers", c.allowHeaders)
	if r.Method == http.MethodOptions && r.Header.Get("Access-Control-Request-Method") != "" {
		w.WriteHeader(http.StatusOK)
		return
	}
	c.nextHandler.ServeHTTP(w, r)
}
