package main

import (
	"os"
	"os/signal"
	"path/filepath"
	"syscall"

	"github.com/luckyshmo/gateway/config"
	"github.com/luckyshmo/gateway/models"
	"github.com/luckyshmo/gateway/pkg/repository"
	"github.com/luckyshmo/gateway/pkg/repository/kafkaQueue"
	"github.com/luckyshmo/gateway/pkg/repository/pg"
	"github.com/luckyshmo/gateway/pkg/service"
	"github.com/luckyshmo/gateway/pkg/source"
	"github.com/luckyshmo/gateway/pkg/source/fileSource"
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

	//Storage init
	kf := kafkaQueue.NewKafkaStore("", "") //example write to Kafka
	pgDB, err := pg.NewPostgresDB(cfg)
	if err != nil {
		return err
	}

	//Source init
	path, err := filepath.Abs("../testData")
	if err != nil {
		return err
	}
	_, err = fileSource.NewFileSource(path) //example. Read from file
	if err != nil {
		return err
	}

	sock := socket.NewSocketSource()

	//Init interfaces
	validRepo := repository.NewRepository(pgDB)
	dataSource := source.NewDataSource(sock)
	invalidRepo := repository.NewRepository(kf)
	services := service.NewService(validRepo, invalidRepo, dataSource)

	//Run program
	chRaw := make(chan models.RawData)
	chValid := make(chan models.Data)
	chInvalid := make(chan models.RawData)
	go services.Reader.ReadData(chRaw)
	go services.Process.SortData(chRaw, chValid, chInvalid) //? no need to create extrenal and interface method //todo middleware
	go services.Writer.WriteData(chValid)
	go services.Writer.WriteRawData(chInvalid)

	// go services.Writer.WriteData(chInvalid) //TODO

	logrus.Print("App Started")

	quit := make(chan os.Signal, 1)
	//if app get SIGTERM it will exit
	signal.Notify(quit, syscall.SIGTERM, syscall.SIGINT)
	<-quit

	return nil
}
