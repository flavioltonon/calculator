package router

import (
	"net/http"

	"calculator/application/http/controller"

	"github.com/gorilla/mux"
)

type Router struct {
	core *mux.Router
}

func New(controller *controller.Controller) *Router {
	router := mux.NewRouter()
	router.HandleFunc("/add", controller.Add).Methods(http.MethodPost)
	router.HandleFunc("/subtract", controller.Subtract).Methods(http.MethodPost)
	router.HandleFunc("/multiply", controller.Multiply).Methods(http.MethodPost)
	router.HandleFunc("/divide", controller.Divide).Methods(http.MethodPost)
	return &Router{core: router}
}

func (r *Router) AddMiddleware(middleware Middleware) {
	r.core.Use(mux.MiddlewareFunc(middleware))
}

func (r *Router) ServeHTTP(writer http.ResponseWriter, request *http.Request) {
	r.core.ServeHTTP(writer, request)
}
