package auth

import (
	"os"
	"testing"
)

func init() {
	os.Setenv("MONGOHOSTNAME", "localhost")
}

var (
	generatedToken   string
	OMGThatSecretKEY = "I'm really enjoying this challenge"
	testEmail        = "nice@challen.ge"
)

func TestNewToken(t *testing.T) {
	jwtWrapper := JwtWrapper{
		SecretKey:       OMGThatSecretKEY,
		Issuer:          "AuthService",
		ExpirationHours: 1,
	}
	token, err := jwtWrapper.NewToken(testEmail)
	if err != nil {
		t.Fatalf("jwtWrapper.NewToken err: %v", err)
	}
	generatedToken = token
}

func TestValidateToken(t *testing.T) {
	jwtWrapper := JwtWrapper{
		SecretKey: OMGThatSecretKEY,
		Issuer:    "AuthService",
	}
	claims, err := jwtWrapper.ValidateToken(generatedToken)
	if err != nil {
		t.Fatalf("jwtWrapper.ValidateToken err: %v", err)
	}

	if claims.Email != testEmail {
		t.Fatalf("email didn't match; expecting %s - found %s", testEmail, claims.Email)
	}
	if claims.Issuer != "AuthService" {
		t.Fatalf("issuer didn't match; expecting 'AuthService' - found %s", claims.Issuer)
	}
}
