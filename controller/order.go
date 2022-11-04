package controller

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func OrderHandler() http.Handler {
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

		fmt.Println(string(b))
	})
}
