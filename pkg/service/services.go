package service

import (
	"fmt"
	"time"

	"github.com/luckyshmo/gateway/models"
	"github.com/luckyshmo/gateway/pkg/repository"
	"github.com/luckyshmo/gateway/pkg/source"
)

type StorageService struct {
	repo repository.Storage
}

func (ss StorageService) WriteData(ch <-chan models.Data) error {
	for {
		select {
		case data := <-ch:
			ss.repo.WriteData(data)
		case <-time.After(500 * time.Millisecond):
			fmt.Println("BOOM!")
			return nil
		default:
			fmt.Println("    .")
			time.Sleep(50 * time.Millisecond)
		}
	}
}

type DataSourceService struct {
	source source.DataSource
}

func (dss *DataSourceService) ReadData(ch chan<- models.RawData) error {
	dss.source.ReadData(ch)
	return nil
}

func NewStorageService(repo repository.Storage) *StorageService {
	return &StorageService{repo: repo}
}

func NewSourceService(dataSource source.DataSource) *DataSourceService {
	return &DataSourceService{source: dataSource}
}

type ProcessService struct {
	kek string
}

func (ps *ProcessService) ProcessData(chRaw <-chan models.RawData, chData chan<- models.Data) error {

	for {
		select {
		case rawData := <-chRaw:
			chData <- models.Data{
				Id:           rawData.Id,
				TimeCreated:  rawData.Time,
				DataComposed: rawData.Data,
			}
		case <-time.After(500 * time.Millisecond):
			fmt.Println("BOOM!")
			return nil
		default:
			fmt.Println("    .")
			time.Sleep(50 * time.Millisecond)
		}
	}
}

func NewProcessService() *ProcessService {
	return &ProcessService{kek: ""}
}
