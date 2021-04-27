package pg

import (
	"github.com/luckyshmo/gateway/config"
	"github.com/rotisserie/eris"
	"github.com/sirupsen/logrus"

	"github.com/golang-migrate/migrate/v4"

	_ "github.com/golang-migrate/migrate/v4/database/mysql"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
)

// runPgMigrations runs Postgres migrations
func RunPgMigrations(cfg *config.Config) error { //? can be run from Makefile
	if cfg.PgMigrationsPath == "" {
		logrus.Warn("No migration path provided")
		return nil
	}
	if cfg.PgHOST == "" || cfg.PgPORT == "" {
		return eris.New("no cfg.PgURL provided")
	}

	connectionString := "postgres://" + cfg.PgUserName + ":" + cfg.PgPAS + "@" + cfg.PgHOST + "/" + cfg.PgDBName + "?sslmode=" + cfg.PgSSLMode
	logrus.Info(connectionString)

	m, err := migrate.New(
		cfg.PgMigrationsPath,
		connectionString,
	)
	if err != nil {
		return err
	}
	if err := m.Up(); err != nil && err != migrate.ErrNoChange {
		return err
	}
	return nil
}
