package middleware

import (
	"net/http"
	"strings"
)

var whitelist = []string{
	"localhost:",
	"127.0.0.1:",
}

// Whitelist restricts access to the localRPC to local nodes
func Whitelist(w http.ResponseWriter, r *http.Request, next http.HandlerFunc) {

	youShallNotPass := true
	for _, host := range whitelist {
		if strings.Index(r.Host, host) != -1 {
			youShallNotPass = false
			break
		}
	}

	if youShallNotPass {
		return
	}

	next(w, r)
}
