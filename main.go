package main

import (
	"encoding/json"
	"keycloak-sample/external"
	"keycloak-sample/params"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/sign_up", handleSignUp)

	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}
}

// /sign_up
func handleSignUp(w http.ResponseWriter, r *http.Request) {
	// POST 以外は無視
	if r.Method != http.MethodPost {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	var form params.SignUp
	if err := json.NewDecoder(r.Body).Decode(&form); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// keycloakコンテナにユーザー登録
	if err := external.GetKeycloakClient().SignUp(r.Context(), form); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
