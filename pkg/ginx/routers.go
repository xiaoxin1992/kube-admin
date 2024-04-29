package ginx

import "github.com/gin-gonic/gin"

var route []router

type router struct {
	Method   string
	URI      string
	Handlers gin.HandlerFunc
	Auth     bool
}

func Register(method string, uri string, auth bool, handlers gin.HandlerFunc) {
	route = append(route, router{
		Method:   method,
		URI:      uri,
		Handlers: handlers,
		Auth:     auth,
	})
}

func init() {
	route = make([]router, 0)
}
