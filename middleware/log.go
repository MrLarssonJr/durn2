package middleware

import (
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

var Log = struct {
	// A middleware that logs what remote accessed what path
	// using the standard logger.
	Access mux.MiddlewareFunc
} {
	func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			log.Printf("%s accessed by %s", r.URL.Path, r.RemoteAddr)
			next.ServeHTTP(w, r)
		})
	},
}