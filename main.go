package main

import (
	"net/http"

	"github.com/haagor/orderMP/controller"
)

func main() {
	http.Handle("/ticket", controller.OrderHandler())
	http.ListenAndServe(":8080", nil)
}
