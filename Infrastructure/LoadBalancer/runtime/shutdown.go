package runtime

import (
	"context"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"subscriptions/Infrastructure/LoadBalancer/libraries"
)

func GracefulShutdown(s *http.Server) {
	stop := make(chan os.Signal, 1)
	signal.Notify(stop, syscall.SIGINT, syscall.SIGTERM)

	sig := <-stop
	libraries.Infof("shutdown signal received: %s", sig.String())

	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	if err := s.Shutdown(ctx); err != nil {
		libraries.Errorf("server shutdown error: %v", err)
	}
}
