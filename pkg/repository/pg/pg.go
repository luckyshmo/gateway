package pg

import (
	"fmt"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq" //driver
	"github.com/luckyshmo/gateway/config"
	"github.com/luckyshmo/gateway/models"
	"github.com/sirupsen/logrus"
)

type PG struct {
	SqlDB *sqlx.DB
}

func (vi *PG) WriteData(...models.ValidPackage) error {
	logrus.Info("write to pg")
	return nil
}

func (vi *PG) WriteRawData(...models.RawData) error {
	logrus.Info("write to pg RAW")
	return nil
}

// type ConfigPG struct { //TODO remove?
// 	Host     string
// 	Port     string
// 	Username string
// 	Password string
// 	DBName   string
// 	SSLMode  string
// }

func NewPostgresDB(cfg *config.Config) (*PG, error) {

	db, err := sqlx.Open("postgres", fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s sslmode=%s",
		cfg.PgHOST, cfg.PgPORT, cfg.PgUserName, cfg.PgDBName, cfg.PgPAS, cfg.PgSSLMode))
	if err != nil {
		return nil, err
	}

	err = db.Ping()
	if err != nil {
		return nil, err
	}

	// err = migrations.RunPgMigrations()
	// if err != nil {
	// 	return nil, err
	// }

	return &PG{
		SqlDB: db,
	}, nil
}

// func (pg *PG) GetSqlDB() *sql.DB {
// 	return pg.SqlDB
// }

// //Init is custom init PG func
// func (pg *PG) Init(conf models.Config) error {
// 	connStr := fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s sslmode=%s",
// 		conf.Host, conf.Port, conf.Username, conf.DBName, conf.Password, conf.SSLMode)

// 	db, err := sql.Open("postgres", connStr)
// 	if err != nil {
// 		return err
// 	}
// 	pg.SqlDB = db

// 	const query = `
// 		CREATE TABLE IF NOT EXISTS rawdata (
// 		  id SERIAL PRIMARY KEY,
// 		  uuid TEXT,
// 		  time TEXT,
// 		  data TEXT
// 	)`

// 	_, err = pg.SqlDB.Exec(query)
// 	if err != nil {
// 		return err
// 	}
// 	// pg.SqlDB.Close()
// 	return nil
// }

// func (pg PG) GetInsertQuery(rawData []models.RawData) string {
// 	query := fmt.Sprintf("INSERT INTO " + "rawdata" + " (uuid, time, data) values \n")

// 	// INSERT INTO products (product_no, name, price) VALUES
// 	// (1, 'Cheese', 9.99),
// 	// (2, 'Bread', 1.99),
// 	// (3, 'Milk', 2.99);

// 	for _, raw := range rawData {
// 		if raw.Data != "" {
// 			query += fmt.Sprintf("('%s', '%s', '%s'), \n", raw.Id, raw.Time, raw.Data)
// 		}
// 	}
// 	query = strings.TrimSpace(query)
// 	query = query[:len(query)-1] + ";"
// 	return query
// }
