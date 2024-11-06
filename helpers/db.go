package helpers

import (
	"codeku.id/sametarget/model"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func ConnectDb() *gorm.DB {
	DB_USER := "postgres"
	DB_PASSWORD := "Welcome4$"
	DB_HOST := "localhost"
	DB_PORT := "5432"
	DB_NAME := "sametarget"

	dsn := "postgres://" + DB_USER + ":" + DB_PASSWORD + "@" + DB_HOST + ":" + DB_PORT + "/" + DB_NAME + "?sslmode=disable"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		panic("Failed to connect to database!")
	}

	return db

}

func MigrateDb(db *gorm.DB) {
	db.AutoMigrate(&model.User{})
	db.AutoMigrate(&model.Teams{})
	db.AutoMigrate(&model.UserTeam{})
	db.AutoMigrate(&model.Skill{})
	db.AutoMigrate(&model.UserSkill{})
}
