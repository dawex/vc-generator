package test_utils

import (
	"context"
	"time"

	"github.com/rs/zerolog/log"
	"github.com/testcontainers/testcontainers-go"
	"github.com/testcontainers/testcontainers-go/wait"
)

func SetupPgContainer() (string, string, func(), error) {
	ctx := context.Background()

	log.Info().Msg("Setup pg container")

	// Configure the PostgreSQL test container
	req := testcontainers.ContainerRequest{
		Image:        "postgres:17",
		ExposedPorts: []string{"5432/tcp"},
		Env: map[string]string{
			"POSTGRES_USER":     "test",
			"POSTGRES_PASSWORD": "test",
			"POSTGRES_DB":       "testdb",
		},
		Cmd:        []string{"postgres", "-c", "fsync=off"},
		WaitingFor: wait.ForLog("database system is ready to accept connections").WithOccurrence(2).WithStartupTimeout(20 * time.Second),
	}

	postgresContainer, err := testcontainers.GenericContainer(ctx, testcontainers.GenericContainerRequest{
		ContainerRequest: req,
		Started:          true,
	})
	if err != nil {
		return "", "", nil, err
	}

	// Get the mapped host and port
	mappedPort, err := postgresContainer.MappedPort(ctx, "5432")
	if err != nil {
		return "", "", nil, err
	}
	hostIP, err := postgresContainer.Host(ctx)
	if err != nil {
		return "", "", nil, err
	}

	cleanup := func() {
		if err := postgresContainer.Terminate(ctx); err != nil {
			log.Error().Msg("Error cleaning container")
		}
	}

	return hostIP, mappedPort.Port(), cleanup, nil
}
