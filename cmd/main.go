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
	_ = kafkaQueue.NewKafkaStore("", "")
	pgDB, err := pg.NewPostgresDB(cfg)
	if err != nil {
		return err
	}

	//Source init
	k, _ := filepath.Abs("../testData")
	fileDataSource, err := fileSource.NewFileSource(k)
	if err != nil {
		return err
	}

	repo := repository.NewRepository(pgDB)
	dataSource := source.NewDataSource(fileDataSource)
	services := service.NewService(repo, dataSource)

	chRaw := make(chan models.RawData)
	chData := make(chan models.Data)
	go services.Reader.ReadData(chRaw)
	go services.Process.ProcessData(chRaw, chData) //? no need to create extrenal and interface method //todo middleware
	go services.Writer.WriteData(chData)

	logrus.Print("App Started")

	quit := make(chan os.Signal, 1)
	//if app get SIGTERM it will exit
	signal.Notify(quit, syscall.SIGTERM, syscall.SIGINT)
	<-quit

	return nil
}
