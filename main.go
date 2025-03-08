package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/Anwarjondev/go-auth-server/database"
	"github.com/Anwarjondev/go-auth-server/handlers"
	"github.com/Anwarjondev/go-auth-server/middleware"
	"github.com/joho/godotenv"
)

func main() {

	godotenv.Load()

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	mux := http.NewServeMux()
	database.InitializeDB()
	defer database.DB.Close()
	fs := http.FileServer(http.Dir("static"))
	mux.Handle("/static/", http.StripPrefix("/static/", fs))

	mux.HandleFunc("/", handlers.HomeHandler)
	mux.HandleFunc("/register", handlers.RegisterHandler)
	mux.HandleFunc("/login", handlers.LoginHandler)
	mux.HandleFunc("/dashboard", middleware.AuthMiddleware(handlers.DashboardHandler))
	mux.HandleFunc("/logout", handlers.Logout)

	secureMux := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Security-Policy", "default-src 'self'")
		w.Header().Set("X-Frame-Options", "DENY")
		w.Header().Set("X-XSS-Protection", "1: mode=block")
		w.Header().Set("X-Content-Type-Options", "nosniff")
		mux.ServeHTTP(w, r)
	})

	fmt.Println("server is running on 8080 port")
	http.ListenAndServe(":"+port, secureMux)
}
