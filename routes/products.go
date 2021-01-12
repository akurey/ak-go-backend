package routes

import (
	"api-test/controllers"
	"net/http"
)

var productsRoutes = []RouteStruct {
	RouteStruct {
		path: "/products",
		httpMethod: http.MethodGet,
		method: controllers.IndexHandlerGET,
		auth: false,
	},
}
