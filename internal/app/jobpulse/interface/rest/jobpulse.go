package rest

import (
	"mycareerapp/internal/app/jobpulse/usecase"
	"mycareerapp/internal/domain/dto"
	"mycareerapp/internal/middleware"
	"net/http"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

type JobpulseHandler struct {
	Validator       *validator.Validate
	JobpulseUseCase usecase.JobpulseUsecaseItf
	Middleware      middleware.MiddleWareI
}

func NewJobpulseHandler(routerGroup fiber.Router, validator *validator.Validate, jobpulseUseCase usecase.JobpulseUsecaseItf, middleware middleware.MiddleWareI) {
	JobpulseHandler := JobpulseHandler{
		Validator:       validator,
		JobpulseUseCase: jobpulseUseCase,
		Middleware:      middleware,
	}

	routerGroup = routerGroup.Group("/jobs")

	routerGroup.Get("/", middleware.Authentication, JobpulseHandler.GetAllJobs)
	routerGroup.Post("/", middleware.Authentication, middleware.Authorization, JobpulseHandler.CreateJob)
	routerGroup.Get("/:id", middleware.Authentication, JobpulseHandler.GetSpecificJob)
	routerGroup.Patch("/:id", middleware.Authentication, middleware.Authorization, JobpulseHandler.UpdateJob)
	routerGroup.Delete("/:id", middleware.Authentication, middleware.Authorization, JobpulseHandler.DeleteJob)
}

func (h JobpulseHandler) GetAllJobs(ctx *fiber.Ctx) error {

	res, err := h.JobpulseUseCase.GetAllJobs()
	if err != nil {
		return err
	}

	return ctx.JSON(fiber.Map{
		"message": res,
	})
}

func (h JobpulseHandler) GetSpecificJob(ctx *fiber.Ctx) error {

	jobID, err := uuid.Parse(ctx.Params("id"))
	if err != nil {
		return fiber.NewError(http.StatusBadRequest, "jobID should an UUID")
	}

	res, err := h.JobpulseUseCase.GetSpecificJob(jobID)
	if err != nil {
		return err
	}

	return ctx.Status(http.StatusOK).JSON(res)
}

func (h JobpulseHandler) CreateJob(ctx *fiber.Ctx) error {

	var request dto.RequestCreateJob

	err := ctx.BodyParser(&request)
	if err != nil {
		return err
	}

	//validasi
	err = h.Validator.Struct(request)
	if err != nil {
		return err
	}

	res, err := h.JobpulseUseCase.CreateJob(request)
	if err != nil {
		return err
	}

	return ctx.Status(http.StatusOK).JSON(fiber.Map{
		"message": "Job created successfully",
		"payload": res,
	})
}

func (h JobpulseHandler) UpdateJob(ctx *fiber.Ctx) error {

	var request dto.RequestUpdateJob

	err := ctx.BodyParser(&request)
	if err != nil {
		return err
	}

	jobID, err := uuid.Parse(ctx.Params("id"))
	if err != nil {
		return fiber.NewError(http.StatusBadRequest, "jobID should an UUID")
	}

	err = h.Validator.Struct(request)
	if err != nil {
		return err
	}

	err = h.JobpulseUseCase.UpdateJob(jobID, request)
	if err != nil {
		return err
	}

	return ctx.SendStatus(http.StatusNoContent)
}

func (h JobpulseHandler) DeleteJob(ctx *fiber.Ctx) error {

	jobID, err := uuid.Parse(ctx.Params("id"))
	if err != nil {
		return fiber.NewError(http.StatusBadRequest, "jobID should an UUID")
	}

	err = h.JobpulseUseCase.DeleteJob(jobID)
	if err != nil {
		return err
	}

	return nil
}
