package database

import (
	"fmt"
	"my_spp/configs"
	"my_spp/models"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var (
	DB *gorm.DB
)

func Init() {
	InitDB()
	InitialMigration()
	// Seeders()
}

type DbSetup struct {
	DB_Username string
	DB_Password string
	DB_Port     string
	DB_Host     string
	DB_Name     string
}

func InitDB() {
	database := DbSetup{
		DB_Username: configs.DB_USERNAME,
		DB_Password: configs.DB_PASSWORD,
		DB_Port:     configs.DB_PORT,
		DB_Host:     configs.DB_HOST,
		DB_Name:     configs.DB_NAME,
	}

	connectionString := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=UTC",
		database.DB_Username,
		database.DB_Password,
		database.DB_Host,
		database.DB_Port,
		database.DB_Name,
	)

	var err error
	DB, err = gorm.Open(mysql.Open(connectionString), &gorm.Config{})
	if err != nil {
		panic(err)
	}
}

func InitialMigration() {
	DB.AutoMigrate(
		&models.User{},
		&models.Teacher{},
		&models.Student{},
		&models.Class{},
		&models.TeacherEducation{},
		&models.Spp{},
		&models.Transaction{},
		&models.ManualPayment{},
		&models.News{},
		&models.Comment{},
	)
}

// func Seeders() {
// 	// ADMIN SEEDERS
// 	adminPasswordHash, err := bcrypt.GenerateFromPassword([]byte("Admin@123"), bcrypt.DefaultCost)

// 	if err != nil {
// 		return
// 	}
// 	admin := []models.Admin{
// 		{
// 			ID:       uuid.New(),
// 			Username: "admin1",
// 			Password: string(adminPasswordHash),
// 			Name:     "Admin 1",
// 		},
// 		{
// 			ID:       uuid.New(),
// 			Username: "admin2",
// 			Password: string(adminPasswordHash),
// 			Name:     "Admin 2",
// 		},
// 		{
// 			ID:       uuid.New(),
// 			Username: "admin3",
// 			Password: string(adminPasswordHash),
// 			Name:     "Admin 3",
// 		},
// 	}

// 	for _, v := range admin {
// 		var exist models.Admin

// 		errCheck := DB.Where("username = ?", v.Username).First(&exist).Error

// 		if errCheck != nil {
// 			DB.Create(&v)
// 		}
// 	}
// }
