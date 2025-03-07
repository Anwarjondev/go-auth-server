package handlers

import (
	"database/sql"
	"html/template"
	"net/http"

	"github.com/Anwarjondev/go-auth-server/database"
	"github.com/Anwarjondev/go-auth-server/middleware"
	"github.com/gorilla/csrf"
	"golang.org/x/crypto/bcrypt"
)

func RegisterHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		http.ServeFile(w, r, "templates/register.html")
		return
	}

	username := r.FormValue("username")
	password := r.FormValue("password")

	hashedpassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		http.Error(w, "Error processing password", http.StatusInternalServerError)
		return
	}
	_, err = database.DB.Exec("Insert into users (username, password) Values(?, ?)", username, hashedpassword)
	if err != nil {
		http.Error(w, "User already exists", http.StatusBadRequest)
		return
	}
	http.Redirect(w, r, "/login", http.StatusSeeOther)
}
func LoginHandler(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.ParseFiles("templates/login.html"))
	if r.Method != http.MethodPost {
		tmpl.Execute(w, map[string]interface{}{
			"CSRFToken": csrf.Token(r),
		})
		return
	}
	username := r.FormValue("username")
	password := r.FormValue("password")

	var storePassword string
	err := database.DB.QueryRow("Select password from users where username = ?", username).Scan(&storePassword)
	if err == sql.ErrNoRows {
		http.Error(w, "Invalid username or password", http.StatusUnauthorized)
		return
	} else if err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
	err = bcrypt.CompareHashAndPassword([]byte(storePassword), []byte(password))
	if err != nil {
		http.Error(w, "Invalid password", http.StatusUnauthorized)
		return
	}
	sessionToken, err := middleware.CreateSession(username)
	if err != nil {
		http.Error(w, "Could not create session", http.StatusInternalServerError)
		return
	}
	http.SetCookie(w, &http.Cookie{
		Name:     "session_token",
		Value:    sessionToken,
		Path:     "/",
		HttpOnly: true,
	})
	http.Redirect(w, r, "/dashboard", http.StatusSeeOther)
}

func Logout(w http.ResponseWriter, r *http.Request) {
	cookie, err := r.Cookie("session_token")
	if err != nil {
		_ = middleware.DestroySession(cookie.Value)
	}
	http.SetCookie(w, &http.Cookie{
		Name:   "session_token",
		Value:  "",
		Path:   "/",
		MaxAge: -1,
	})
	http.Redirect(w, r, "/login", http.StatusSeeOther)
}
