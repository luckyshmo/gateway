package repository

import "github.com/luckyshmo/gateway/models"

type Storage interface {
	WriteData(...models.RawData) error
}

type Repository struct {
	Storage
}

func NewRepository(rp Storage) *Repository {
	return &Repository{
		Storage: rp,
	}
}

// //TODO фнукции для БД общие
// //TODO а функции получение query и инициализации специфичны

// func (db *DataBase) Init(cfg models.Config) { //TODO вот это тсранно, получается для создания инстанса нужно импортировать зависимости???
// 	if cfg.DBName == "postgres" {
// 		var pgDb pg.PG
// 		pgDb.Init(cfg)
// 		db.dbI = &pgDb
// 	} else {
// 		var viDb vi.Vi
// 		viDb.Init(cfg)
// 		db.dbI = &viDb
// 	}
// }

// func (dbL *DataBase) InsertRawData(rawData []models.RawData) (int, error) {
// 	var id int
// 	var db = dbL.dbI //TODO пахнет

// 	query := db.GetInsertQuery(rawData)

// 	row := db.GetSqlDB().QueryRow(query) //TODO чета сложна выглядит

// 	if row.Err() != nil {
// 		log.Println(row.Err())
// 	}
// 	return id, nil
// }
