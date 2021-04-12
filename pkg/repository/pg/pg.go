package pg

import (
	"fmt"

	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq" //driver
	"github.com/luckyshmo/gateway/config"
	"github.com/luckyshmo/gateway/models"
	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"
)

type PG struct {
	SqlDB *sqlx.DB
}

var (
	rawTable   = "raw_data"
	validTable = "valid_data"
)

func (pg *PG) WriteData(vp ...models.ValidPackage) error {
	logrus.Info(fmt.Sprintf("Writing %d valid packages to pg", len(vp)))
	err := pg.SqlDB.Ping()
	if err != nil {
		return errors.Wrap(err, "no DB connections")
	}
	for i, v := range vp {
		var id uuid.UUID
		query := fmt.Sprintf("INSERT INTO %s (id, dev_eui, time_cr, time_p, data_f, raw_data) values ($1, $2, $3, $4, $5, $6) RETURNING id", validTable)

		row := pg.SqlDB.QueryRow(query, v.Id, v.DevEui, v.TimeCreated, v.TimePackage, v.Data, v.RawData)
		if err := row.Scan(&id); err != nil {
			logrus.Warn("No resp from DB")
		}

		row.Err()
		if err := row.Err(); err != nil {
			return errors.Wrap(err, fmt.Sprintf("error executing query number %d", i))
		}
	}

	return nil
}

func (pg *PG) WriteRawData(rd ...models.RawData) error {
	err := pg.SqlDB.Ping()
	if err != nil {
		return errors.Wrap(err, "no DB connections")
	}
	logrus.Info(fmt.Sprintf("Writing %d invalid packages to pg", len(rd)))
	for i, r := range rd {
		var id uuid.UUID
		query := fmt.Sprintf("INSERT INTO %s (id, time_cr, data_r) values ($1, $2, $3) RETURNING id", rawTable)

		row := pg.SqlDB.QueryRow(query, r.Id, r.Time, r.Data)
		if err := row.Scan(&id); err != nil {
			logrus.Warn("No resp from DB")
		}
		row.Err()
		if err := row.Err(); err != nil {
			return errors.Wrap(err, fmt.Sprintf("error executing query number %d", i))
		}
	}

	return nil
}

func NewPostgresDB(cfg *config.Config) (*PG, error) {

	db, err := sqlx.Open("postgres", fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s sslmode=%s",
		cfg.PgHOST, cfg.PgPORT, cfg.PgUserName, cfg.PgDBName, cfg.PgPAS, cfg.PgSSLMode))
	if err != nil {
		return nil, errors.Wrap(err, "err open SQLx")
	}

	err = db.Ping()
	if err != nil {
		return nil, errors.Wrap(err, "no DB connections")
	}

	if err = RunPgMigrations(cfg); err != nil {
		return nil, errors.Wrap(err, "migrations failed")
	}

	return &PG{
		SqlDB: db,
	}, nil
}
