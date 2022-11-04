package main

import (
	"net/http"

	"github.com/haagor/orderMP/controller"
)

func main() {
	http.Handle("/ticket", controller.TicketHandler())
	http.ListenAndServe(":8080", nil)
}
