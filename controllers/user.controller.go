package controllers

import (
	"github.com/Reksaditya/project-management/models"
	"github.com/Reksaditya/project-management/services"
	"github.com/Reksaditya/project-management/utils"
	"github.com/gofiber/fiber/v2"
	"github.com/jinzhu/copier"
)

type UserController struct {
	service services.UserService
}

func NewUserController(s services.UserService) *UserController {
	return &UserController{service: s}
}

func (c *UserController) Register(ctx *fiber.Ctx) error {
	user := new(models.User)

	if err := ctx.BodyParser(user); err != nil {
		return utils.BadRequest(ctx, "Failed to parse data", err.Error())
	}

	if err := c.service.Register(user); err != nil {
		return utils.BadRequest(ctx, "Register failed", err.Error())
	}

	var userRes models.UserRes
	_ = copier.Copy(&userRes, &user)
	return utils.Success(ctx, "Register success", userRes)
}