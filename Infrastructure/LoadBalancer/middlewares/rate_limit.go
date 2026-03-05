package middlewares

import (
	"fmt"
	"net/http"

	"subscriptions/Infrastructure/LoadBalancer/libraries/parsers"

	"golang.org/x/time/rate"
)

func init() {
	Register("rate_limit", func(cfg map[string]any) (Middleware, error) {
		rps, _ := parsers.GetFloat(cfg, "rps", 20)
		burst, _ := parsers.GetInt(cfg, "burst", 50)

		if rps <= 0 || burst <= 0 {
			return nil, fmt.Errorf("rate_limit: invalid config (rps=%v, burst=%d), expected rps>0 and burst>0", rps, burst)
		}

		return RateLimit(rps, burst), nil
	})
}

func RateLimit(rps float64, burst int) Middleware {
	limiter := rate.NewLimiter(rate.Limit(rps), burst)

	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			if !limiter.Allow() {
				http.Error(w, "rate limit exceeded", http.StatusTooManyRequests)
				return
			}
			next.ServeHTTP(w, r)
		})
	}
}
