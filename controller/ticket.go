package controller

import (
	"fmt"
	"net/http"
)

func TicketHandler() http.Handler {
	return http.HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {
		if request.Method != "POST" {
			http.Error(writer, "Invalid request method.", 405)
			return
		}

		fmt.Println("Hello world")
	})
}
