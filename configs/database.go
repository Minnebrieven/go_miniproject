package configs

import (
	"fmt"

	"swim-class/models"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func ConnectDB() (*gorm.DB, error) {
	// config := models.Config{
	// 	DB_Host:     os.Getenv("DB_HOST"),
	// 	DB_Port:     os.Getenv("DB_PORT"),
	// 	DB_Name:     os.Getenv("DB_NAME"),
	// 	DB_Username: os.Getenv("DB_USERNAME"),
	// 	DB_Password: os.Getenv("DB_PASSWORD"),
	// }

	connectionString := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local",
		AppConfig.DBUsername,
		AppConfig.DBPassword,
		AppConfig.DBHost,
		AppConfig.DBPort,
		AppConfig.DBName,
	)
	return gorm.Open(mysql.Open(connectionString), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})
}

func MigrateDB(db *gorm.DB) error {
	return db.AutoMigrate(
		models.User{},
		models.Instructor{},
		models.ClassCategory{},
		models.Class{},
		models.ClassParticipant{},
	)
}
