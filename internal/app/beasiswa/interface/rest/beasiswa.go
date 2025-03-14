package rest

import (
	"mycareerapp/internal/app/beasiswa/usecase"
	"mycareerapp/internal/domain/dto"
	"mycareerapp/internal/middleware"
	"net/http"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
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
	routerGroup.Post("/", middleware.Authentication, middleware.Authorization, BeasiswaHandler.GetAllBeasiswa)
	routerGroup.Get("/:id", middleware.Authentication, BeasiswaHandler.GetAllBeasiswa)
	routerGroup.Patch("/:id", middleware.Authentication, middleware.Authorization, BeasiswaHandler.GetAllBeasiswa)
	routerGroup.Delete("/:id", middleware.Authentication, middleware.Authorization, BeasiswaHandler.GetAllBeasiswa)
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

func (h BeasiswaHandler) GetSpecificBeasiswa(ctx *fiber.Ctx) error {

	beasiswaID, err := uuid.Parse(ctx.Params("id"))
	if err != nil {
		return fiber.NewError(http.StatusBadRequest, "BeasiswaID should an UUID")
	}

	res, err := h.BeasiswaUseCase.GetSpecificBeasiswa(beasiswaID)
	if err != nil {
		return err
	}

	return ctx.Status(http.StatusOK).JSON(res)
}

func (h BeasiswaHandler) CreateBeasiswa(ctx *fiber.Ctx) error {

	var request dto.RequestCreateBeasiswa

	err := ctx.BodyParser(&request)
	if err != nil {
		return err
	}

	err = h.Validator.Struct(request)
	if err != nil {
		return err
	}

	res, err := h.BeasiswaUseCase.CreateBeasiswa(request)
	if err != nil {
		return err
	}

	return ctx.Status(http.StatusOK).JSON(fiber.Map{
		"message": "Beasiswa created successfully",
		"payload": res,
	})
}

func (h BeasiswaHandler) UpdateBeasiswa(ctx *fiber.Ctx) error {

	var request dto.RequestUpdateBeasiswa

	err := ctx.BodyParser(&request)
	if err != nil {
		return err
	}

	BeasiswaID, err := uuid.Parse(ctx.Params("id"))
	if err != nil {
		return fiber.NewError(http.StatusBadRequest, "BeasiswaID should an UUID")
	}

	err = h.Validator.Struct(request)
	if err != nil {
		return err
	}

	err = h.BeasiswaUseCase.UpdateBeasiswa(BeasiswaID, request)
	if err != nil {
		return err
	}

	return ctx.SendStatus(http.StatusNoContent)
}

func (h BeasiswaHandler) DeleteBeasiswa(ctx *fiber.Ctx) error {

	BeasiswaID, err := uuid.Parse(ctx.Params("id"))
	if err != nil {
		return fiber.NewError(http.StatusBadRequest, "BeasiswaID should an UUID")
	}

	err = h.BeasiswaUseCase.DeleteBeasiswa(BeasiswaID)
	if err != nil {
		return err
	}

	return nil
}
