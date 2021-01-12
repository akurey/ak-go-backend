package routes

import (
	"api-test/controllers"
	"api-test/middlewares"
	"github.com/gorilla/mux"
	"os"
)

func RouterConfig() *mux.Router {
	pathPrefix := os.Getenv("API_PATH_PREFIX")
	router := mux.NewRouter().StrictSlash(true).PathPrefix(pathPrefix).Subrouter()

	router.Use(middlewares.RequestLogger)

	// Index Request
	router.HandleFunc("/", controllers.IndexHandlerGET).Methods("GET")

	// Users Request
	router.HandleFunc("/users/register", controllers.UsersRegisterHandler).Methods("POST")
	router.HandleFunc("/users/login", controllers.UsersLogin).Methods("POST")

	mountModule(cardsRoutes, router)
	mountModule(productsRoutes, router)

	return router
}

func mountModule(module []RouteStruct, router *mux.Router) {
	for _, route := range module {
		path := "/api" + route.path
		if route.auth {
			router.HandleFunc(path, middlewares.JWTMiddleware(route.method)).Methods(route.httpMethod)
			return
		}
		router.HandleFunc(path, route.method).Methods(route.httpMethod)
	}
}


