package routes

import (
	"api-test/controllers"
	"net/http"
)

var cardsRoutes = []RouteStruct {
	RouteStruct {
		path: "/cards",
		httpMethod: http.MethodGet,
		method: controllers.IndexHandlerGET,
		auth: false,
	},
}
