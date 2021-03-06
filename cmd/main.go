package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/luckyshmo/gateway/config"
	process "github.com/luckyshmo/gateway/pkg/processors"
	"github.com/luckyshmo/gateway/pkg/repository/influx"
	"github.com/luckyshmo/gateway/pkg/repository/kafkaQueue"
	"github.com/luckyshmo/gateway/pkg/repository/pg"
	"github.com/luckyshmo/gateway/pkg/service"
	"github.com/luckyshmo/gateway/pkg/source"
	"github.com/luckyshmo/gateway/pkg/source/fileSource"
	"github.com/luckyshmo/gateway/pkg/source/socket"
	"github.com/rotisserie/eris"
	"github.com/sirupsen/logrus"
)

//TODO proper close socket and over connection while graceful shutdown
func main() {
	if err := run(); err != nil {
		// format error for Prod
		formattedJSON := eris.ToJSON(err, true)
		logrus.Error(formattedJSON)

		// format error for Debug
		formattedStr := eris.ToString(err, true)
		fmt.Println(formattedStr)
	}
}

func run() error {
	// Config
	cfg := config.Get()
	fileSource.ReadFile("../catalina")
	// logger configuration
	lvl, err := logrus.ParseLevel(cfg.LogLevel)
	if err != nil {
		logrus.SetLevel(logrus.DebugLevel) //using debug lvl if we can't parse
		logrus.Warn("Using debug level logger")
	} else {
		logrus.SetLevel(lvl)
	}
	if cfg.Environment == "production" {
		var JSONF = new(logrus.JSONFormatter)
		JSONF.TimestampFormat = time.RFC3339
		logrus.SetFormatter(JSONF)
	}

	//Storage init
	_, err = kafkaQueue.NewKafkaStore("", "") //example write to Kafka
	if err != nil {
		return eris.Wrap(err, "Error Init kafka")
	}
	pgDB, err := pg.NewPostgresDB(cfg)
	if err != nil {
		return eris.Wrap(err, "Error Init PG")
	}
	defer pgDB.SqlDB.Close()

	// pgDB1, err := pg.NewPostgresDB(cfg)
	// if err != nil {
	// 	return err
	// }
	// defer pgDB1.SqlDB.Close()

	inf, err := influx.NewInfluxWriter(cfg)
	if err != nil {
		return eris.Wrap(err, "Error Init influx")
	}
	//Source init
	// path, err := filepath.Abs("../testData")
	// if err != nil {
	// 	return err
	// }
	// _, err = fileSource.NewFileSource(path) //example. Read from file
	// if err != nil {
	// 	return err
	// }

	//Process Init
	jsonProcessService := process.NewJsonProcessService()

	sock := socket.NewSocketSource(cfg)
	dataSource := source.NewDataSource(sock)
	ps := process.NewProcessService(jsonProcessService)
	services := service.NewService(inf, pgDB, dataSource, ps)

	//Run program
	services.Init()

	logrus.Print("App Started")

	quit := make(chan os.Signal, 1)
	//if app get SIGTERM it will exit
	signal.Notify(quit, syscall.SIGTERM, syscall.SIGINT)
	<-quit

	return nil
}
