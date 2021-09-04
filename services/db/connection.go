package db

import (
	"github.com/Pauloo27/archvium/logger"
	"github.com/Pauloo27/archvium/model"
	"github.com/Pauloo27/archvium/utils"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var Connection *gorm.DB

// TODO: change port to uint? it would require parse from the env...
// just to be stringified later...
func Connect(host, username, password, dbname, port string) error {
	dsn := utils.Fmt(
		"host=%s user=%s password=%s dbname=%s port=%s",
		host, username, password, dbname, port,
	)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		CreateBatchSize: 1000,
	})

	Connection = db

	return err
}

func Setup() {
	Connection.AutoMigrate(&model.User{})

	err := Connection.Create(&model.User{
		Username: "admin",
		Password: utils.EnvString("AUTH_ADMIN_PASSWORD"),
		Email:    "admin@localhost",
	}).Error

	if err != nil && !utils.IsNotUnique(err) {
		logger.HandleFatal(err, "Cannot create admin user")
	}
}
