package middleware

import (
	"context"
	"net/http"
	"rcp-api-data/internal/config/security"
	"strings"

	"go.uber.org/zap"
)

type CtxTokenKey string

const TokenKey CtxTokenKey = "token"

func ValidateTokenMiddleware(next http.Handler, cfg *security.Environment, sugar *zap.SugaredLogger) http.Handler {

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Envoi de la requête GET à l'API d'authentification
		if cfg == nil {
			sugar.Error("Erreur lors de la requête à l'API d'authentification 1, config nulle")
			http.Error(w, "Erreur lors de la requête à l'API d'authentification 1", http.StatusInternalServerError)
			return
		}
		authHeader := r.Header.Get("Authorization")
		if !strings.HasPrefix(authHeader, "Bearer ") || len(strings.Split(authHeader, " ")) != 2 {
			sugar.Error("Invalid 'Authorization' header, please check the format : 'Bearer <token>'")
			http.Error(w, "Invalid 'Authorization' header", http.StatusUnauthorized)
		}
		// Envoi de la requête GET à l'API d'authentification avec un header Authorization
		req, err := http.NewRequest("GET", cfg.GetAuthURL()+"/token-check", nil)
		if err != nil {
			sugar.Error("Erreur lors de la requête à l'API d'authentification (step 2)", err)
			http.Error(w, "Erreur lors de la requête à l'API d'authentification (step 2)", http.StatusInternalServerError)
			return
		}
		req.Header.Set("Authorization", authHeader)
		resp, err := http.DefaultClient.Do(req)
		if err != nil {
			sugar.Error("Erreur lors de la requête à l'API d'authentification (step 3)", err)
			http.Error(w, "Erreur lors de la requête à l'API d'authentification (step 3)", http.StatusInternalServerError)
			return
		}
		defer resp.Body.Close()
		// Vérifiez si le statut de la réponse est 200
		if resp.StatusCode == http.StatusOK {
			// Extraire le bearer token du header authorize
			rawToken := strings.Split(authHeader, " ")[1]
			if rawToken == "" {
				sugar.Error("Token non valide")
				http.Error(w, "Token non valide", http.StatusUnauthorized)
				return
			}
			// Décoder le token JWT
			token, err := security.DecodePayload(rawToken)
			if err != nil {
				sugar.Error("Erreur lors du décodage du token JWT", err)
				http.Error(w, "Erreur lors du décodage du token JWT", http.StatusUnauthorized)
				return
			}
			// Ajouter les informations décodées à la requête, si nécessaire
			newRequest := r.WithContext(context.WithValue(r.Context(), TokenKey, token))
			// Passez à la prochaine manipulation dans la chaîne
			next.ServeHTTP(w, newRequest)
		} else {
			sugar.Error("Accès non autorisé")
			http.Error(w, "Accès non autorisé", http.StatusUnauthorized)
		}
	})
}
