package dto

import "github.com/google/uuid"

type ResponseGetJob struct {
	ID              uuid.UUID `json:"id"`
	Title           string    `json:"title"`
	Location        string    `json:"location"`
	WorkDescription string    `json:"work_description"`
	MinimumExp      string    `json:"minimum_exp"`
	Salary          string    `json:"salary"`
	WorkingDays     string    `json:"working_days"`
	WorkingHours    string    `json:"working_hours"`
	PhotoURL        string    `json:"photo_url"`
	ArticleURL      string    `json:"article_url"`
}

type RequestCreateJob struct {
	Title           string `json:"title" validate:"required,min=3"`
	Location        string `json:"location" validate:"required"`
	WorkDescription string `json:"work_description" validate:"required,min=10"`
	MinimumExp      string `json:"minimum_exp"`
	Salary          string `json:"salary" validate:"required"`
	WorkingDays     string `json:"working_days" validate:"required"`
	WorkingHours    string `json:"working_hours" validate:"required"`
	PhotoURL        string `json:"photo_url"`
	ArticleURL      string `json:"article_url" validate:"required"`
}

type ResponseCreateJob struct {
	Title           string `json:"title"`
	Location        string `json:"location"`
	WorkDescription string `json:"work_description"`
	MinimumExp      string `json:"minimum_exp"`
	Salary          string `json:"salary"`
	WorkingDays     string `json:"working_days"`
	WorkingHours    string `json:"working_hours"`
	PhotoURL        string `json:"photo_url"`
	ArticleURL      string `json:"article_url"`
}

type RequestUpdateJob struct {
	Title           string `json:"title" validate:"omitempty,min=3"`
	Location        string `json:"location"`
	WorkDescription string `json:"work_description" validate:"omitempty,min=10"`
	MinimumExp      string `json:"minimum_exp"`
	Salary          string `json:"salary"`
	WorkingDays     string `json:"working_days"`
	WorkingHours    string `json:"working_hours"`
	PhotoURL        string `json:"photo_url"`
	ArticleURL      string `json:"article_url" validate:"omitempty"`
}
