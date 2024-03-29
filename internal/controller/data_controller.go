package controller

import (
	"fmt"
	"net/http"
	"rcp-api-data/internal/utils"
)

func DataController() http.HandlerFunc {
    return func(w http.ResponseWriter, r *http.Request) {
		
        switch r.Method {
		case "GET":
			token, err := utils.GetContextToken(r)
			if err != nil {
				http.Error(w, "Token introuvable", http.StatusUnauthorized)
				return
			}
			fmt.Println(token)
		}
    }
}

