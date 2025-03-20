package app

import (
	"time"

	"github.com/dawex/vc-generator/internal/common/config"
	"github.com/dawex/vc-generator/internal/common/db"
	"github.com/dawex/vc-generator/internal/common/logger"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/render"
	"gorm.io/gorm"

	event_handler "github.com/dawex/vc-generator/internal/services/event/adapters/handler"
	event_repository "github.com/dawex/vc-generator/internal/services/event/adapters/repository"
	event_core "github.com/dawex/vc-generator/internal/services/event/core"
	event_ports "github.com/dawex/vc-generator/internal/services/event/ports"

	vc_handler "github.com/dawex/vc-generator/internal/services/verifiable-credential/adapters/handler"
	vc_repository "github.com/dawex/vc-generator/internal/services/verifiable-credential/adapters/repository"
	vc_core "github.com/dawex/vc-generator/internal/services/verifiable-credential/core"
	vc_ports "github.com/dawex/vc-generator/internal/services/verifiable-credential/ports"
)

func NewApp(app_name string, app_config config.Config) (*gorm.DB, *chi.Mux) {
	// Init logger
	logger.NewZerolog(app_config)

	// Init db
	gormDb := db.NewPostgres(app_config)

	// Init router
	router := chi.NewRouter()
	pathPrefix := "/v1"

	// Register Middlewares
	router.Use(
		logger.Logger, // Use custom logger

		// Set a timeout value on the request context (ctx), that will signal
		// through ctx.Done() that the request has timed out and further
		// processing should be stopped.
		middleware.Timeout(10*time.Second),

		// middlewares
		middleware.Recoverer,
		middleware.Heartbeat(pathPrefix+"/healthcheck"),

		// response type is forced to JSON
		render.SetContentType(render.ContentTypeJSON),
	)

	// Register Services
	event_ports.HandlerWithOptions(event_handler.NewHandler(event_core.New(event_repository.New(gormDb))), event_ports.ChiServerOptions{BaseURL: pathPrefix, BaseRouter: router})
	vc_ports.HandlerWithOptions(vc_handler.NewHandler(vc_core.New(app_config, vc_repository.New(gormDb), event_repository.New(gormDb))), vc_ports.ChiServerOptions{BaseURL: pathPrefix, BaseRouter: router})

	return gormDb, router
}
