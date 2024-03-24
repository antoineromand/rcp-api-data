package security

import "testing"
func TestJwt(t *testing.T) {
	t.Run("should decode the payload", func(t *testing.T) {
		token := "eyJ1dWlkIjoiZDY0ZmZmYTgtZGNlYi00ZDY0LThlYzAtN2Y4MmMwYjFkZTQwIiwiZXhwIjoiMjAyNC0wMy0yNVQyMTo1Njo0OC4xMjJaIiwiaXNzdWVyIjoiYXV0aGVudGljYXRpb24tMS44LjEiLCJhdWRpZW5jZSI6ImFuZGVzaXRlIn0=.eyJ1dWlkIjoiOTA1ZTkwNWQtNDZkMy00NjE3LWFkNzctYzJiODNhNDMzODRhIiwidXNlcm5hbWUiOiJhbmRlc2l0ZSIsInJvbGVQZXJtaXNzaW9uIjp7ImFkbWluIjpbImFkbWluIl19fQ==.MVqhYVIzs2A3XZTwAM387h0TYYCmA+CxK9nYNy9ul4Mr4FgYAl9Y7pWOOXx7MqyPqk9uEJXYndrDCQJfm0U6Ag=="
		_, err := DecodePayload(token)
		if err != nil {
			t.Errorf("Error while decoding payload: %v", err)
		}
	})
	t.Run("should return an error if the token is invalid", func(t *testing.T) {
		token := "ererz.erzer"
		_, err := DecodePayload(token)
		if err == nil {
			t.Errorf("invalid token")
		}
	})
	t.Run("should return an error if the token is empty", func(t *testing.T) {
		token := ""
		_, err := DecodePayload(token)
		if err == nil {
			t.Errorf("token is empty")
		}
	})
}
