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
