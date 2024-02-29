package controller

import (
	"fmt"
	"net/http"
)

func DataController() http.HandlerFunc {
    return func(w http.ResponseWriter, r *http.Request) {

        switch r.Method {
		case "GET":
			fmt.Fprintf(w, "Hello, World!")
		case "POST":
			fmt.Fprintf(w, "Hello, World!")
		}
    }
}

