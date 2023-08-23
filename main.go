package main

import (
	"log"

	"github.com/ditrit/badaas-orm-example/standalone/models"
	"github.com/ditrit/badaas/orm"
	"github.com/ditrit/badaas/orm/logger"
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

	log.Println("You are ready to do queries with orm.NewQuery[models.MyModel]")
}

func NewDBConnection() (*gorm.DB, error) {
	return orm.Open(
		postgres.Open(orm.CreatePostgreSQLDSN("localhost", "root", "postgres", "disable", "badaas_db", 26257)),
		&gorm.Config{
			Logger: logger.Default.ToLogMode(logger.Info),
		},
	)
}
