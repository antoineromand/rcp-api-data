package controller

import (
	"fmt"
	"net/http"
	"rcp-api-data/internal/middleware"
)

func DataController() http.HandlerFunc {
    return func(w http.ResponseWriter, r *http.Request) {

        switch r.Method {
		case "GET":
			fmt.Print(r.Context().Value(middleware.TokenKey))
		  	fmt.Fprintf(w, "Hello, World!")
		}
    }
}

