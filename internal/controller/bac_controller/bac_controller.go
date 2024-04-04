package baccontroller

import (
	"net/http"
	"rcp-api-data/internal/utils"
)

func BacController() http.HandlerFunc {

	return func(w http.ResponseWriter, r *http.Request) {
		sugar := utils.GetLogger()

		switch r.Method {
		case "POST":
			token, err := utils.GetContextToken(r)
			if err != nil {
				http.Error(w, "Token introuvable", http.StatusUnauthorized)
				return
			}
			sugar.Info("Token: ", token)
		}
	}
}
