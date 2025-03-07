package middleware

import (
	"log"
	"net/http"
	"os"

	"github.com/gorilla/csrf"
	"github.com/joho/godotenv"
)


func CSRFMiddleware(next http.Handler) http.Handler {

	err := godotenv.Load()
	if err != nil {
		log.Println("Warning: No .env file found")
	}
	csrfkey := os.Getenv("CSRF_SECRET") 
	if csrfkey == "" {
		log.Fatal("CSRF is not set in .env file")
	}
	csrfProtection := csrf.Protect([]byte("V+rS+M/9NaAVE6y0DbOi9j6Ej+fbYWOco7cQVLD/S40="),
		csrf.Secure(false),
	)
	return csrfProtection(next)
}