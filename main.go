package main

import (
	"encoding/json"
	"keycloak-sample/auth"
	"keycloak-sample/params"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", handleRoot)
	http.HandleFunc("/sign_up", handleSignUp)
	http.HandleFunc("/login", handleLogin)

	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}
}

// GET /
// NOTE: ログインしていない場合はリダイレクトするようにする
func handleRoot(w http.ResponseWriter, r *http.Request) {
	// GET 以外は無視
	if r.Method != http.MethodGet {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	// 認証情報の検証
	if err := auth.GetKeycloakClient().ValidateToken(r.Context()); err != nil {
		http.Error(w, err.Error(), http.StatusUnauthorized)
		return
	}

	if _, err := w.Write([]byte("Hello, World!")); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

// POST /sign_up
func handleSignUp(w http.ResponseWriter, r *http.Request) {
	// POST 以外は無視
	if r.Method != http.MethodPost {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	// KeycloakにPOSTするパラメータを作成
	var form params.SignUp
	if err := json.NewDecoder(r.Body).Decode(&form); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// keycloakコンテナにユーザー登録
	if err := auth.GetKeycloakClient().SignUp(r.Context(), form); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

// POST /login
func handleLogin(w http.ResponseWriter, r *http.Request) {
	// POST 以外は無視
	if r.Method != http.MethodPost {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	// KeycloakにPOSTするパラメータを作成
	var form params.Login
	if err := json.NewDecoder(r.Body).Decode(&form); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	client := auth.GetKeycloakClient()

	// keycloakコンテナにユーザー登録
	if err := client.Login(r.Context(), form); err != nil {
		http.Error(w, err.Error(), http.StatusUnauthorized)
		return
	}
}
