package main

import (
	"encoding/json"
	"keycloak-sample/external"
	"keycloak-sample/params"
	"log"
	"net/http"

	"github.com/Nerzal/gocloak/v13"
)

var (
	token *gocloak.JWT
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
func handleRoot(w http.ResponseWriter, r *http.Request) {
	// GET 以外は無視
	if r.Method != http.MethodGet {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	// XXX: ログインしていない場合（やばいコード）
	// if token == nil {
	// 	// NOTE: redirectするようにしたい…！
	// 	http.Error(w, "Unauthorized", http.StatusUnauthorized)
	// 	return
	// }

	// 認証情報の検証
	if err := external.GetKeycloakClient().ValidateToken(r.Context(), token); err != nil {
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

// POST /login
func handleLogin(w http.ResponseWriter, r *http.Request) {
	// POST 以外は無視
	if r.Method != http.MethodPost {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	var form params.Login
	if err := json.NewDecoder(r.Body).Decode(&form); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// keycloakコンテナにユーザー登録
	got, err := external.GetKeycloakClient().Login(r.Context(), form)
	if err != nil {
		http.Error(w, err.Error(), http.StatusUnauthorized)
		return
	}

	// tokenを保持
	// NOTE: 本来ならcontextなどに保持すべきか
	token = got
}
