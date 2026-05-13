package seed

import (
	"log"

	"github.com/Reksaditya/project-management/config"
	"github.com/Reksaditya/project-management/models"
	"github.com/Reksaditya/project-management/utils"
	"github.com/google/uuid"
)

func AdminSeed() {
	password, _ := utils.HashPassword("admin123")
	
	admin := models.User{
		Name: "Admin",
		Email: "admin@example.com",
		Password: password,
		Role: "admin",
		PublicID: uuid.New(),
	}

	if err := config.DB.FirstOrCreate(&admin, models.User{Email: admin.Email}).Error; err != nil {
		log.Println("Failed to Seed Admin", err)
	} else {
		log.Println("Admin Seeded Successfully")
	}
}