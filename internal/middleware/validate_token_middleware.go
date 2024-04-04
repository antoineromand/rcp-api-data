package middleware

import (
	"context"
	"net/http"
	"rcp-api-data/internal/config/security"
	"strings"

	"go.uber.org/zap"
)

type TokenMiddleware struct {
	cfg   *security.Environment
	sugar *zap.SugaredLogger
}

type CtxTokenKey string

const TokenKey CtxTokenKey = "token"

func NewTokenMiddleware(cfg *security.Environment, sugar *zap.SugaredLogger) *TokenMiddleware {
	return &TokenMiddleware{
		cfg:   cfg,
		sugar: sugar,
	}
}

func (t *TokenMiddleware) ValidateTokenMiddleware(next http.Handler) http.Handler {

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Envoi de la requête GET à l'API d'authentification
		if t.cfg == nil {
			t.sugar.Error("Erreur lors de la requête à l'API d'authentification 1, config nulle")
			http.Error(w, "Erreur lors de la requête à l'API d'authentification 1", http.StatusInternalServerError)
			return
		}
		authHeader := r.Header.Get("Authorization")
		if !strings.HasPrefix(authHeader, "Bearer ") || len(strings.Split(authHeader, " ")) != 2 {
			t.sugar.Error("Invalid 'Authorization' header, please check the format : 'Bearer <token>'")
			http.Error(w, "Invalid 'Authorization' header", http.StatusUnauthorized)
		}
		// Envoi de la requête GET à l'API d'authentification avec un header Authorization
		req, err := http.NewRequest("GET", t.cfg.GetAuthURL()+"/token-check", nil)
		if err != nil {
			t.sugar.Error("Erreur lors de la requête à l'API d'authentification (step 2)", err)
			http.Error(w, "Erreur lors de la requête à l'API d'authentification (step 2)", http.StatusInternalServerError)
			return
		}
		req.Header.Set("Authorization", authHeader)
		resp, err := http.DefaultClient.Do(req)
		if err != nil {
			t.sugar.Error("Erreur lors de la requête à l'API d'authentification (step 3)", err)
			http.Error(w, "Erreur lors de la requête à l'API d'authentification (step 3)", http.StatusInternalServerError)
			return
		}
		defer resp.Body.Close()
		// Vérifiez si le statut de la réponse est 200
		if resp.StatusCode == http.StatusOK {
			// Extraire le bearer token du header authorize
			rawToken := strings.Split(authHeader, " ")[1]
			if rawToken == "" {
				t.sugar.Error("Token non valide")
				http.Error(w, "Token non valide", http.StatusUnauthorized)
				return
			}
			// Décoder le token JWT
			token, err := security.DecodePayload(rawToken)
			if err != nil {
				t.sugar.Error("Erreur lors du décodage du token JWT", err)
				http.Error(w, "Erreur lors du décodage du token JWT", http.StatusUnauthorized)
				return
			}
			// Ajouter les informations décodées à la requête, si nécessaire
			newRequest := r.WithContext(context.WithValue(r.Context(), TokenKey, token))
			// Passez à la prochaine manipulation dans la chaîne
			next.ServeHTTP(w, newRequest)
		} else {
			t.sugar.Error("Accès non autorisé")
			http.Error(w, "Accès non autorisé", http.StatusUnauthorized)
		}
	})
}
