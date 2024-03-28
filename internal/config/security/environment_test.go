package security

import "testing"

func TestEnvironment(t *testing.T) {
	t.Run("should return an environment struct", func(t *testing.T) {
		_, err := InitEnvironment(true)
		if err != nil {
			t.Errorf("error while parsing env variables")
		}
	})
	t.Run("should return an error if env variables are not set", func(t *testing.T) {
		_, err := InitEnvironment(false)
		if err != nil {
			t.Errorf("error while parsing env variables")
		}
	})
	t.Run("should return the auth URL", func(t *testing.T) {
		env, err := InitEnvironment(true)
		if err != nil {
			t.Errorf("error while parsing env variables")
		}
		if env.GetAuthURL() == "" {
			t.Errorf("error while getting auth URL")
		}
	})
}
