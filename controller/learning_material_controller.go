package controller

import "github.com/gofiber/fiber/v2"

type LearningMaterialController interface {
	FindAll(ctx *fiber.Ctx) error
	Create(ctx *fiber.Ctx) error
	FindByID(ctx *fiber.Ctx) error
	SaveProgress(ctx *fiber.Ctx) error
	Update(ctx *fiber.Ctx) error
	Delete(ctx *fiber.Ctx) error
	FindUserProgress(ctx *fiber.Ctx) error
	FindUserProgressById(ctx *fiber.Ctx) error
}
