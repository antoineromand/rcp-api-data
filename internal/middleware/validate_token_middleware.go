package middleware

import (
	"context"
	"fmt"
	"net/http"
	"rcp-api-data/internal/config/security"
	"strings"
)

type CtxTokenKey string

const TokenKey CtxTokenKey = "token"

func ValidateTokenMiddleware(next http.Handler, cfg *security.Environment) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Envoi de la requête GET à l'API d'authentification
		if cfg == nil {
			http.Error(w, "Erreur lors de la requête à l'API d'authentification 1", http.StatusInternalServerError)
			return
		}
		authHeader := r.Header.Get("Authorization")
		if !strings.HasPrefix(authHeader, "Bearer ") || len(strings.Split(authHeader, " ")) != 2 {
			http.Error(w, "Invalid 'Authorization' header", http.StatusUnauthorized)
		}
		// Envoi de la requête GET à l'API d'authentification avec un header Authorization
		req, err := http.NewRequest("GET", cfg.GetAuthURL()+cfg.RCP_AUTH_PREFIX+"/token-check", nil)
		if err != nil {
			http.Error(w, "Erreur lors de la requête à l'API d'authentification (step 2)", http.StatusInternalServerError)
			return
		}
		req.Header.Set("Authorization", authHeader)
		resp, err := http.DefaultClient.Do(req)
		if err != nil {
			fmt.Println(err)
			http.Error(w, "Erreur lors de la requête à l'API d'authentification (step 3)", http.StatusInternalServerError)
			return
		}
		defer resp.Body.Close()
		// Vérifiez si le statut de la réponse est 200
		if resp.StatusCode == http.StatusOK {

			// Extraire le bearer token du header authorize
			rawToken := strings.Split(authHeader, " ")[1]
			if rawToken == "" {
				http.Error(w, "Token non valide", http.StatusUnauthorized)
				return
			}
			// Décoder le token JWT
			token, err := security.DecodePayload(rawToken)
			if err != nil {
				http.Error(w, "Erreur lors du décodage du token JWT", http.StatusUnauthorized)
				return
			}
			// Ajouter les informations décodées à la requête, si nécessaire
			newRequest := r.WithContext(context.WithValue(r.Context(), TokenKey, token))
			// Passez à la prochaine manipulation dans la chaîne
			next.ServeHTTP(w, newRequest)
		} else {
			http.Error(w, "Accès non autorisé", http.StatusUnauthorized)
		}
	})
}
