package main

import (
	"net/http"
	"os"
	"time"

	serverhttp "subscriptions/Infrastructure/LoadBalancer/http"
	"subscriptions/Infrastructure/LoadBalancer/libraries"
	"subscriptions/Infrastructure/LoadBalancer/libraries/config"
	"subscriptions/Infrastructure/LoadBalancer/libraries/parsers"
	"subscriptions/Infrastructure/LoadBalancer/middlewares"
	"subscriptions/Infrastructure/LoadBalancer/runtime"
)

func main() {

	mux := http.NewServeMux()

	fs := http.FileServer(http.Dir("./public"))
	mux.Handle("/public/", http.StripPrefix("/public/", fs))

	backendsEnv := os.Getenv("API_BACKENDS")
	backends := parsers.SplitCSV(backendsEnv)

	mux.Handle("/api/", serverhttp.LoadBalancer(backends))

	handler := middlewares.Apply(mux)

	port := config.Env("GATEWAY_INT_PORT", "8443")
	cert := config.Env("CERT_PATH", "/app/certs/cert.pem")
	key := config.Env("KEY_PATH", "/app/certs/key.pem")

	server := &http.Server{
		Addr:         ":" + port,
		Handler:      handler,
		ReadTimeout:  15 * time.Second,
		WriteTimeout: 30 * time.Second,
		IdleTimeout:  120 * time.Second,
	}

	go func() {
		libraries.Infof("HTTPS gateway started on :%s, cert=%s, key=%s", port, cert, key)

		if err := server.ListenAndServeTLS(cert, key); err != nil && err != http.ErrServerClosed {
			libraries.Errorf("https server failed: %v", err)
			os.Exit(1)
		}
	}()

	runtime.GracefulShutdown(server)
}
