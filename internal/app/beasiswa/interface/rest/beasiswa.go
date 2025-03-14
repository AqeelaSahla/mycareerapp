package rest

import (
	"mycareerapp/internal/app/beasiswa/usecase"
	"mycareerapp/internal/middleware"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
)

type BeasiswaHandler struct {
	Validator       *validator.Validate
	BeasiswaUseCase usecase.BeasiswaUsecaseItf
	Middleware      middleware.MiddleWareI
}

func NewBeasiswaHandler(routerGroup fiber.Router, validator *validator.Validate, beasiswaUseCase usecase.BeasiswaUsecaseItf, middleware middleware.MiddleWareI) {
	BeasiswaHandler := BeasiswaHandler{
		Validator:       validator,
		BeasiswaUseCase: beasiswaUseCase,
		Middleware:      middleware,
	}

	routerGroup = routerGroup.Group("/beasiswa")

	routerGroup.Get("/", middleware.Authentication, BeasiswaHandler.GetAllBeasiswa)
}

func (h *BeasiswaHandler) GetAllBeasiswa(ctx *fiber.Ctx) error {

	res, err := h.BeasiswaUseCase.GetAllBeasiswa()
	if err != nil {
		return err
	}

	return ctx.JSON(fiber.Map{
		"message": res,
	})
}
