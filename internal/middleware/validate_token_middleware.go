package middleware

import (
	"context"
	"net/http"
	"rcp-api-data/internal/config"
)

type CtxTokenKey string

const TokenKey CtxTokenKey = "token"

func ValidateTokenMiddleware(next http.Handler) http.Handler {
    return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        // Envoi de la requête GET à l'API d'authentification
        cfg := config.GetConfig()
        if cfg == nil { 
            http.Error(w, "Erreur lors de la requête à l'API d'authentification 1", http.StatusInternalServerError)
            return 
        }
        resp, err := http.Get(cfg.GetAuthURL() + "/token-check")
        if err != nil {
            http.Error(w, "Erreur lors de la requête à l'API d'authentification 2", http.StatusInternalServerError)
            return
        }
        defer resp.Body.Close()
        // Vérifiez si le statut de la réponse est 200
        if resp.StatusCode == http.StatusOK {
            // Extraire le token JWT du cookie
            //cookie, err := r.Cookie("token")
            //fmt.Print(cookie.Value)
            cookie := "eyJ1dWlkIjoiYTU0MjUxZGMtODdhNi00YTkzLWFlYTYtNzU5YTIwYzBhZDZkIiwiZXhwIjoiMjAyNC0wMy0wMlQwOToxOToyNy4yNzVaIiwiaXNzdWVyIjoiYXV0aGVudGljYXRpb24tMS42LjEiLCJhdWRpZW5jZSI6ImFuZGVzaXRlIn0=.eyJ1dWlkIjoiOTA1ZTkwNWQtNDZkMy00NjE3LWFkNzctYzJiODNhNDMzODRhIiwidXNlcm5hbWUiOiJhbmRlc2l0ZSIsInJvbGVQZXJtaXNzaW9uIjp7ImFkbWluIjpbImFkbWluIl19fQ==.jk/hiegvenlvIFitfCyX+6C1TIFDOD0MczNq1vT+bneSHmdk6Yj33j3xeAC1a3LMO8z4E4mAp9E+m6mu9KzRDg=="
            if err != nil {
                http.Error(w, "Token JWT introuvable dans le cookie", http.StatusUnauthorized)
                return
            }

            // Décoder le token JWT
            token, err := config.DecodePayload(cookie)

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