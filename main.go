package main

import (
	"log"

	"github.com/FrancoLiberali/cql"
	"github.com/FrancoLiberali/cql-quickstart/models"
	"github.com/FrancoLiberali/cql/logger"
	"gorm.io/driver/sqlite"
)

func main() {
	db, err := NewDBConnection()
	if err != nil {
		panic(err)
	}

	err = db.GormDB.AutoMigrate(
		models.MyModel{},
	)
	if err != nil {
		panic(err)
	}

	log.Println("You are ready to do queries with cql.Query[models.MyModel]")
}

func NewDBConnection() (*cql.DB, error) {
	return cql.Open(
		sqlite.Open("sqlite:db"),
		&cql.Config{
			Logger: logger.Default.ToLogMode(logger.Info),
		},
	)
}
