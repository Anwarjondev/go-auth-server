package handlers

import (
	"fmt"
	"net/http"

	"github.com/Anwarjondev/go-auth-server/middleware"
)

func DashboardHandler(w http.ResponseWriter, r *http.Request) {
	// Retrieve session token from cookie
	cookie, err := r.Cookie("session_token")
	if err != nil {
		http.Error(w, "Unauthorized access", http.StatusUnauthorized)
		return
	}

	// Retrieve username from the session table
	username, err := middleware.GetSession(cookie.Value)
	if err != nil {
		http.Error(w, "Session expired. Please log in again.", http.StatusUnauthorized)
		return
	}
	if username == "" {
		http.Error(w, "Invalid session. Please log in.", http.StatusUnauthorized)
		return
	}

	// Display correct username
	fmt.Fprintf(w, "Welcome to your dashboard, %s!", username)
}