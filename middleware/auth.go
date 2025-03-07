package middleware

import "net/http"


func AuthMiddleware(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request){
		cookie, err := r.Cookie("session_token")
		if err != nil{
			http.Redirect(w, r, "/login", http.StatusSeeOther)
			return
		}
		username, err := GetSession(cookie.Value)
		if err != nil || username == "" {
			http.Redirect(w,r, "/login", http.StatusSeeOther)
			return
		}
		next(w,r)
	}
}