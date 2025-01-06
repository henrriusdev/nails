package store

import (
	"fmt"

	"github.com/lib/pq"

	"github.com/henrriusdev/nails/config"
)

func NewConnection(cfg config.EnvVar) (Queryable, error) {
	dsn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=%s",
		cfg.DBHost, cfg.DBPort, cfg.DBUser, cfg.DBPass, cfg.DBName, cfg.DBSSLMode)

	if err := connection.Ping(); err != nil {
		log.Error().Stack().Err(err).Msg("failed to ping db")
	}

	return connection, nil
}
