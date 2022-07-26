package routes

import (
	"net/http"
	"ze-is-store/controller"
)

func GetRoutes() {
	http.HandleFunc("/", controller.Index)
	http.HandleFunc("/new", controller.New)

}
