package main

import (
	"fmt"
	"log"

	"github.com/FrancoLiberali/cql"
	"github.com/FrancoLiberali/cql-quickstart/models"
	"github.com/FrancoLiberali/cql/logger"
	"gorm.io/driver/postgres"
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
		postgres.Open(
			fmt.Sprintf(
				"user=%s password=%s host=%s port=%d sslmode=%s dbname=%s",
				"root", "postgres", "localhost", 26257, "disable", "cql_db",
			),
		),
		&gorm.Config{
			Logger: logger.Default.ToLogMode(logger.Info),
		},
	)
}
