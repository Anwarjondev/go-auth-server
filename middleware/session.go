package middleware

import (
	"database/sql"
	"math/rand"
	"time"

	"github.com/Anwarjondev/go-auth-server/database"
)

var letters = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789")

func generateSessionToken() string {
	rand.Seed(time.Now().UnixNano())
	token := make([]rune, 32)
	for i := range token {
		token[i] = letters[rand.Intn(len(letters))]
	}
	return string(token)
}

func CreateSession(username string) (string, error) {
	token := generateSessionToken()
	_, err := database.DB.Exec("Insert into sessions (username, session_token) values(?,?)", username, token)
	if err != nil {
		return "", err
	}
	return token, nil
}

func GetSession(token string) (string, error) {
	var username string
	err := database.DB.QueryRow("select username from sessions where session_token = ?", token).Scan(&username)
	if err == sql.ErrNoRows {
		return "", nil
	}
	return username, err
}

func DestroySession(token string) error {
	_, err := database.DB.Exec("delete from sessions where session_token = ?", token)
	return err
}
