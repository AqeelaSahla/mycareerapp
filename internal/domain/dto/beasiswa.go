package dto

import "github.com/google/uuid"

type ResponseGetBeasiswa struct {
	ID                  uuid.UUID `json:"id"`
	Title               string    `json:"title"`
	EducationLevel      string    `json:"educational_level"`
	EducationLevel2     string    `json:"educational_level2"`
	BeasiswaDescription string    `json:"beasiswa_description"`
	Organizer           string    `json:"organizer"`
	Type                string    `json:"type"`
	OpenReg             string    `json:"open_reg"`
	CloseReg            string    `json:"close_reg"`
	PhotoURL            string    `json:"photo_url"`
	BeasiswaURL         string    `json:"beasiswa_url"`
}

type RequestCreateBeasiswa struct {
	Title               string `json:"title" validate:"required,min=3"`
	EducationLevel      string `json:"educational_level" validate:"required"`
	EducationLevel2     string `json:"educational_level2"`
	BeasiswaDescription string `json:"beasiswa_description" validate:"required,min=10"`
	Organizer           string `json:"organizer" validate:"required"`
	Type                string `json:"type" validate:"required"`
	OpenReg             string `json:"open_reg" validate:"required"`
	CloseReg            string `json:"close_reg" validate:"required"`
	PhotoURL            string `json:"photo_url"`
	BeasiswaURL         string `json:"beasiswa_url" validate:"required"`
}

type ResponseCreateBeasiswa struct {
	Title               string `json:"title"`
	EducationLevel      string `json:"educational_level"`
	EducationLevel2     string `json:"educational_level2"`
	BeasiswaDescription string `json:"beasiswa_description"`
	Organizer           string `json:"organizer"`
	Type                string `json:"type"`
	OpenReg             string `json:"open_reg"`
	CloseReg            string `json:"close_reg"`
	PhotoURL            string `json:"photo_url"`
	BeasiswaURL         string `json:"beasiswa_url"`
}

type RequestUpdateBeasiswa struct {
	Title               string `json:"title" validate:"omitempty,min=3"`
	EducationLevel      string `json:"educational_level" validate:"omitempty"`
	EducationLevel2     string `json:"educational_level2"`
	BeasiswaDescription string `json:"beasiswa_description" validate:"omitempty,min=10"`
	Organizer           string `json:"organizer" validate:"omitempty"`
	Type                string `json:"type" validate:"omitempty"`
	OpenReg             string `json:"open_reg" validate:"omitempty"`
	CloseReg            string `json:"close_reg" validate:"omitempty"`
	PhotoURL            string `json:"photo_url"`
	BeasiswaURL         string `json:"beasiswa_url" validate:"omitempty"`
}
