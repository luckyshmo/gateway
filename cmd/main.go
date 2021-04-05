package main

import (
	"log"

	"github.com/luckyshmo/gateway/config"
	"github.com/luckyshmo/gateway/pkg/repository"
	"github.com/luckyshmo/gateway/pkg/repository/pg"
	"github.com/luckyshmo/gateway/pkg/service"
	"github.com/sirupsen/logrus"
)

func main() {
	if err := run(); err != nil {
		log.Fatal(err)
	}
}

func run() error {
	// config
	cfg := config.Get()

	// kf := kafkaQueue.NewKafkaStore("", "")
	cfg1 := pg.Config{
		Host:     cfg.PgHOST,
		Port:     cfg.PgPORT,
		Username: cfg.PgUserName,
		DBName:   cfg.PgDBName,
		SSLMode:  cfg.PgSSLMode,
		Password: cfg.PgPAS,
	}
	p1, err := pg.NewPostgresDB(cfg1)
	if err != nil {
		logrus.Fatal("asdasd")
	}

	repo := repository.NewRepository(p1)
	services := service.NewService(repos)
	// handlers := handler.NewHandler(services)

	// confPG := models.Config{
	// 	Host:     "localhost",
	// 	Port:     "5432",
	// 	Username: "postgres",
	// 	Password: "example",
	// 	DBName:   "postgres",
	// 	SSLMode:  "disable",
	// }
	// var lol = d.DataBase{}
	// lol.Init(confPG)
	// err := f.ReadFile("/home/mihail/go/src/gateway/fileReader/test")
	// if err != nil {
	// 	return err
	// }
	// return nil
}
