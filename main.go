package main

import (
	"log"

	"github.com/FrancoLiberali/cql"
	"github.com/FrancoLiberali/cql-quickstart/models"
	"github.com/FrancoLiberali/cql/logger"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func main() {
	gormDB, err := NewDBConnection()
	if err != nil {
		panic(err)
	}

	err = gormDB.AutoMigrate(
		models.MyModel{},
	)
	if err != nil {
		panic(err)
	}

	log.Println("You are ready to do queries with cql.Query[models.MyModel]")
}

func NewDBConnection() (*gorm.DB, error) {
	return cql.Open(
		sqlite.Open("sqlite:db"),
		&gorm.Config{
			Logger: logger.Default.ToLogMode(logger.Info),
		},
	)
}
