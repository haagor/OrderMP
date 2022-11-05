package controller

import (
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/haagor/orderMP/adapter"
)

func OrderHandler(ordersChan chan string, pa adapter.PostgresAdapter) http.Handler {
	return http.HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {
		if request.Method != "POST" {
			http.Error(writer, "Invalid request method.", 405)
			return
		}

		b, err := ioutil.ReadAll(request.Body)
		if err != nil {
			http.Error(writer, "Bad request Data.", 400)
			return
		}

		select {
		case ordersChan <- string(b):
			fmt.Println(string(b) + "\n")
			writer.WriteHeader(http.StatusOK)
			return
		default:
			http.Error(writer, "Internal Server Error.", 500)
			return
		}
	})
}
