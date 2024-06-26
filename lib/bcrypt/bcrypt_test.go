package bcrypt

import "testing"

func TestHashPassword(t *testing.T) {
	hash, err := HashPassword("password")
	if err != nil {
		t.Errorf("error hashing password %v", err)
	}

	if hash == "" {
		t.Error("expected hash not empty")
	}

	if hash == "password" {
		t.Error("expected hash to be different from password")
	}
}

func TestComparePassword(t *testing.T) {
	hash, err := HashPassword("password")
	if err != nil {
		t.Errorf("error hashing password %v", err)
	}

	if !ComparePassword(hash, []byte("password")) {
		t.Error("expected password to match hash")
	}

	if ComparePassword(hash, []byte("notpassword")) {
		t.Error("expected hash to not match hash")
	}
}
