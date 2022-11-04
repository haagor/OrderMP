package controller

import (
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/haagor/orderMP/adapter"
	"github.com/haagor/orderMP/model"
)

func OrderHandler(pa adapter.PostgresAdapter) http.Handler {
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

		order, err := model.StringToOrder(string(b))
		if err != nil {
			http.Error(writer, "Bad request Data.", 400)
			return
		}

		err = pa.AddOrderWithProduct(order)
		if err != nil {
			http.Error(writer, "Internal Server Error.", 500)
			return
		}
		fmt.Println(order)
	})
}
