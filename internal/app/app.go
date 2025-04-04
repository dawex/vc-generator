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

	negotiationcontracts_handler "github.com/dawex/vc-generator/internal/services/negotiation-contracts/adapters/handler"
	negotiationcontracts_repository "github.com/dawex/vc-generator/internal/services/negotiation-contracts/adapters/repository"
	negotiationcontracts_core "github.com/dawex/vc-generator/internal/services/negotiation-contracts/core"
	negotiationcontracts_ports "github.com/dawex/vc-generator/internal/services/negotiation-contracts/ports"

	compliancelogs_handler "github.com/dawex/vc-generator/internal/services/compliance-logs/adapters/handler"
	compliancelogs_repository "github.com/dawex/vc-generator/internal/services/compliance-logs/adapters/repository"
	compliancelogs_core "github.com/dawex/vc-generator/internal/services/compliance-logs/core"
	compliancelogs_ports "github.com/dawex/vc-generator/internal/services/compliance-logs/ports"

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
	negotiationcontracts_ports.HandlerWithOptions(negotiationcontracts_handler.NewHandler(negotiationcontracts_core.New(negotiationcontracts_repository.New(gormDb))), negotiationcontracts_ports.ChiServerOptions{BaseURL: pathPrefix, BaseRouter: router})
	compliancelogs_ports.HandlerWithOptions(compliancelogs_handler.NewHandler(compliancelogs_core.New(compliancelogs_repository.New(gormDb))), compliancelogs_ports.ChiServerOptions{BaseURL: pathPrefix, BaseRouter: router})
	vc_ports.HandlerWithOptions(vc_handler.NewHandler(vc_core.New(app_config, vc_repository.New(gormDb), compliancelogs_repository.New(gormDb), negotiationcontracts_repository.New(gormDb))), vc_ports.ChiServerOptions{BaseURL: pathPrefix, BaseRouter: router})

	return gormDb, router
}
