package middleware

import (
	"log"
	"net/http"
	"time"
)

// Logger logs info about every request to the console
func Logger(w http.ResponseWriter, r *http.Request, next http.HandlerFunc) {
	log.Printf(
		"%s\t%s\t%s",
		r.Method,
		r.RequestURI,
		time.Now(),
	)

	next(w, r)
}
