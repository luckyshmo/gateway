package config

import (
	"sync"

	"github.com/luckyshmo/gateway/tools"
	"github.com/sirupsen/logrus"

	"github.com/kelseyhightower/envconfig"
)

// Config. Should be filled from Env. Use launch.json(vscode) on local machine
type Config struct {
	PgHOST           string `envconfig:"PG_HOST"`
	PgPORT           string `envconfig:"PG_PORT"`
	PgPAS            string `envconfig:"PG_PAS"`
	PgSSLMode        string `envconfig:"PG_SSLMODE"`
	PgMigrationsPath string `envconfig:"PG_MIGRATIONS_PATH"`
	PgUserName       string `envconfig:"PG_USERNAME"`
	PgDBName         string `envconfig:"PG_DBNAME"`

	InfluxUrl    string `envconfig:"INFLUX_URL"`
	InfluxToken  string `envconfig:"INFLUX_TOKEN"`
	InfluxOrg    string `envconfig:"INFLUX_ORG"`
	InfluxBucket string `envconfig:"INFLUX_BUCKET"`

	SocketHost string `envconfig:"SOCKET_HOST"`

	Environment string `envconfig:"ENV"`

	AppPort  string `envconfig:"APP_PORT"`
	LogLevel string `envconfig:"LOG_LEVEL"`
}

var (
	config Config
	once   sync.Once
)

// Get reads config from environment. Once.
func Get() *Config {
	once.Do(func() {
		err := envconfig.Process("", &config)
		if err != nil {
			logrus.Fatal(err)
		}
		tools.PrintEmptyStructFields(config)
	})
	return &config
}
