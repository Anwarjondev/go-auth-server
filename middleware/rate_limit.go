package middleware

import (
	"net/http"
	"sync"
	"time"
)

var clients =  make(map[string]time.Time)
var lock = sync.Mutex{}

func RateLimitMiddleware(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		ip := r.RemoteAddr

		lock.Lock()
		lastAttempt, exists := clients[ip]
		lock.Unlock()

		if exists && time.Since(lastAttempt) < 5 * time.Second {
			http.Error(w, "Two many requests, please try again later", http.StatusTooManyRequests)
			return
		}
		lock.Lock()
		clients[ip] = time.Now()
		lock.Unlock()

		next(w,r)
	}
}