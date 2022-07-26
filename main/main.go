package main

import (
	"net/http"
	"ze-is-store/routes"
)

func main() {
	routes.GetRoutes()
	http.ListenAndServe(":8000", nil)
}
