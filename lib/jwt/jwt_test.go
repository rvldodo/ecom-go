package jwt

import (
	"testing"
)

func TestCreateJWT(t *testing.T) {
	token, err := CreateToken(1)
	if err != nil {
		t.Errorf("error create token %v", err)
	}

	if token == "" {
		t.Error("expected token not to be empty")
	}
}
