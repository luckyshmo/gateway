package main

import (
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/luckyshmo/gateway/config"
	"github.com/luckyshmo/gateway/pkg/repository"
	"github.com/luckyshmo/gateway/pkg/repository/pg"
	"github.com/luckyshmo/gateway/pkg/service"
	"github.com/luckyshmo/gateway/pkg/source"
	"github.com/luckyshmo/gateway/pkg/source/socket"
	"github.com/sirupsen/logrus"
)

func main() {
	if err := run(); err != nil {
		logrus.Fatal(err)
	}
}

func run() error {
	// config
	cfg := config.Get()

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
	// _ = kafkaQueue.NewKafkaStore("", "") //example write to Kafka
	pgDB, err := pg.NewPostgresDB(cfg)
	if err != nil {
		return err
	}
	defer pgDB.SqlDB.Close()

	pgDB1, err := pg.NewPostgresDB(cfg)
	if err != nil {
		return err
	}
	defer pgDB.SqlDB.Close()

	//Source init
	// path, err := filepath.Abs("../testData")
	// if err != nil {
	// 	return err
	// }
	// _, err = fileSource.NewFileSource(path) //example. Read from file
	// if err != nil {
	// 	return err
	// }

	sock := socket.NewSocketSource(cfg)

	//Init interfaces
	validRepo := repository.NewRepository(pgDB1)
	invalidRepo := repository.NewRepository(pgDB)
	dataSource := source.NewDataSource(sock)
	services := service.NewService(validRepo, invalidRepo, dataSource)

	//Run program
	services.Init()

	logrus.Print("App Started")

	quit := make(chan os.Signal, 1)
	//if app get SIGTERM it will exit
	signal.Notify(quit, syscall.SIGTERM, syscall.SIGINT)
	<-quit

	return nil
}
