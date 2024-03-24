package security

import "testing"

func TestEnvironment(t *testing.T) {
	t.Run("should return an environment struct", func(t *testing.T) {
		env := InitEnvironment(true)
		if env == nil {
			t.Errorf("error while parsing env variables")
		}
	})
	t.Run("should return an error if env variables are not set", func(t *testing.T) {
		env := InitEnvironment(false)
		if env == nil {
			t.Errorf("error while parsing env variables")
		}
	})
	t.Run("should return the auth URL", func(t *testing.T) {
		env := InitEnvironment(true)
		if env.GetAuthURL() == "" {
			t.Errorf("error while getting auth URL")
		}
	})
}