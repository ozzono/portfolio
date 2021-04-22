package route

import (
	"encoding/json"
	"fmt"
	"net/http"
	"request/auth"
	"request/models"
	"strings"
)

const (
	magicSecretKey = "I'm really enjoying this challenge"
)

func OpenTheGates(token string) error {
	jwtWrapper := auth.JwtWrapper{
		SecretKey: magicSecretKey,
		Issuer:    "AuthService",
	}
	_, err := jwtWrapper.ValidateToken(token)
	if err != nil {
		return fmt.Errorf("jwtWrapper.ValidateToken err: %v", err)
	}
	return nil
}

func PayTheLoad(r *http.Request) (map[string]string, error) {
	payload := map[string]string{}
	err := json.NewDecoder(r.Body).Decode(&payload)
	if err != nil {
		return nil, fmt.Errorf("invalid json data")
	}

	_, ok := payload["token"]
	if !ok {
		payload["token"] = strings.TrimPrefix(r.Header.Get("authorization"), "Bearer ")
	}
	return payload, nil
}

func RefreshDB(w http.ResponseWriter, r *http.Request) {
	payload, err := PayTheLoad(r)
	if err != nil {
		http.Error(w, fmt.Sprintf("{\"msg\":\"%v\"}", err), http.StatusBadRequest)
		return
	}
	err = OpenTheGates(payload["token"])
	if err != nil {
		http.Error(w, fmt.Sprint("{\"msg\":\"unauthorized access\"}"), http.StatusBadRequest)
		return
	}
	_, err = models.DefaultDB()
	if err != nil {
		http.Error(w, fmt.Sprintf("{\"msg\":\"models.DefaultDB err - %v\"}", err), http.StatusBadRequest)
		return
	}
	http.Error(w, fmt.Sprint("{\"success\":\"successfully refreshed db\"}"), http.StatusOK)
}
