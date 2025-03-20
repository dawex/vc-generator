package integration_test

import (
	"fmt"
	"testing"

	"github.com/dawex/vc-generator/internal/app"
	"github.com/dawex/vc-generator/internal/common/config"
	test_utils "github.com/dawex/vc-generator/test/integration/test-utils"
	"github.com/go-chi/chi/v5"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	"gorm.io/gorm"
)

const app_name = "vc-generator_test"

func TestHandler(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Control Suite")
}

var (
	db         *gorm.DB
	router     *chi.Mux
	dbHost     string
	dbPort     string
	pgShutdown func()
	app_config config.Config
)

var _ = BeforeSuite(func() {
	var err error
	dbHost, dbPort, pgShutdown, err = test_utils.SetupPgContainer()
	Expect(err).ToNot(HaveOccurred())

	// Initialize handler to test
	app_config = config.Config{
		Server: config.Server{
			Env: "DEV",
		},
		Db: config.Db{
			Addr: fmt.Sprintf("host=%s user=test password=test dbname=testdb port=%s sslmode=disable TimeZone=UTC", dbHost, dbPort),
		},
		Logs: config.Logs{
			Level: "DEBUG",
		},
		Security: config.Security{
			Seed: "sxdw993fhn3stuihb9maiaw762a69uh7",
		},
		Issuer: config.Issuer{
			ID:   "urn:uuid:f4d9f1ad-9965-4948-8243-38e4b6e704c6",
			Name: "TEST ISSUER",
		},
	}
	db, router = app.NewApp(app_name, app_config)
})

var _ = AfterSuite(func() {
	// Clean up the test environment
	pgShutdown()
})
