package routes

import "net/http"

//RouteStruct RouteStruct
type RouteStruct struct {
	path string
	method func(http.ResponseWriter, *http.Request)
	httpMethod  string
	auth bool
}
