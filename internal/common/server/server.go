package server

import (
	"context"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/dawex/vc-generator/internal/common/config"
	"github.com/rs/zerolog/log"
)

func Run(app_name string, app_config config.Config, handler http.Handler) error {
	// Handle SIGINT (CTRL+C) gracefully.
	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt)

	// Start HTTP server.
	log.Info().Msgf("Server is running on %s", app_config.Server.Addr)
	srv := &http.Server{
		Addr:         app_config.Server.Addr,
		Handler:      handler,
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
		IdleTimeout:  60 * time.Second,
	}
	srvErr := make(chan error, 1)
	go func() { srvErr <- srv.ListenAndServe() }()

	// Wait for interruption.
	select {
	case err := <-srvErr: // Error when starting HTTP server.
		return err
	case <-ctx.Done(): // Wait for first CTRL+C.
		stop() // Stop receiving signal notifications as soon as possible.
	}

	// When Shutdown is called, ListenAndServe immediately returns ErrServerClosed.
	return srv.Shutdown(context.Background())
}
