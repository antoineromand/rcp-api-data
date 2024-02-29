package middleware

import (
	"encoding/base64"
	"encoding/json"
	"net/http"
	"strings"
)

func ValidateTokenMiddleware(next http.Handler) http.Handler {
    return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        // Envoi de la requête GET à l'API d'authentification
        resp, err := http.Get("URL_DE_VOTRE_API/validateToken")
        if err != nil {
            http.Error(w, "Erreur lors de la requête à l'API d'authentification", http.StatusInternalServerError)
            return
        }
        defer resp.Body.Close()

        // Vérifiez si le statut de la réponse est 200
        if resp.StatusCode == http.StatusOK {
            // Extraire le token JWT du cookie
            cookie, err := r.Cookie("jwt")
            if err != nil {
                http.Error(w, "Token JWT introuvable dans le cookie", http.StatusUnauthorized)
                return
            }

            // Décoder le payload du token
            splitToken := strings.Split(cookie.Value, ".")
            if len(splitToken) != 3 {
                http.Error(w, "Token JWT invalide", http.StatusUnauthorized)
                return
            }

            payload, err := base64.RawURLEncoding.DecodeString(splitToken[1])
            if err != nil {
                http.Error(w, "Erreur lors du décodage du token JWT", http.StatusUnauthorized)
                return
            }

            var decodedPayload map[string]interface{}
            json.Unmarshal(payload, &decodedPayload)

            // Ajouter les informations décodées à la requête, si nécessaire
            // r = r.WithContext(context.WithValue(r.Context(), "decodedPayload", decodedPayload))

            // Passez à la prochaine manipulation dans la chaîne
            next.ServeHTTP(w, r)
        } else {
            http.Error(w, "Accès non autorisé", http.StatusUnauthorized)
        }
    })
}