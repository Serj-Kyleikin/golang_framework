package middlewares

import (
	"net/http"

	"subscriptions/Infrastructure/LoadBalancer/libraries"
)

func init() {
	Register("logging", func(cfg map[string]any) (Middleware, error) {
		return Logging, nil
	})
}

func Logging(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		libraries.Infof("%s %s %s", r.RemoteAddr, r.Method, r.URL.Path)
		next.ServeHTTP(w, r)
	})
}
