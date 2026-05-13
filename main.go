package main

import (
	"log"

	"github.com/Reksaditya/project-management/config"
	"github.com/Reksaditya/project-management/controllers"
	// "github.com/Reksaditya/project-management/database/seed"
	"github.com/Reksaditya/project-management/repositories"
	"github.com/Reksaditya/project-management/routes"
	"github.com/Reksaditya/project-management/services"
	"github.com/gofiber/fiber/v2"
)

func main() {
	config.LoadEnv()
	config.ConnectDB()

	// seed.AdminSeed()
	app := fiber.New()

	userRepo := repositories.NewUserRepository()
	userService := services.NewUserSevice(userRepo)
	userController := controllers.NewUserController(userService)

	routes.Setup(app, userController)

	port := config.AppConfig.AppPort
	log.Println("Server is running on port", port)
	log.Fatal(app.Listen(":" + port))
}