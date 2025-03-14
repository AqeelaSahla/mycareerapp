package entity

import (
	"mycareerapp/internal/domain/dto"
	"time"

	"github.com/google/uuid"
)

type Beasiswa struct {
	ID                  uuid.UUID `gorm:"type:char(36);primaryKey"`
	Title               string    `gorm:"type:varchar(100);not null"`
	EducationLevel      string    `gorm:"type:varchar(100);not null"`
	EducationLevel2     string    `gorm:"type:varchar(100)"`
	BeasiswaDescription string    `gorm:"type:text"`
	Organizer           string    `gorm:"type:text;not null"`
	Type                string    `gorm:"type:varchar(36);not null"`
	OpenReg             string    `gorm:"type:text;not null"`
	CloseReg            string    `gorm:"type:text;not null"`
	PhotoURL            string    `gorm:"type:text"`
	BeasiswaURL         string    `gorm:"type:text;not null"`
	CreatedAt           time.Time `gorm:"type:timestamp;autoCreateTime"`
}

func (p Beasiswa) ParseToDTOGet() dto.ResponseGetBeasiswa {
	return dto.ResponseGetBeasiswa{
		ID:                  p.ID,
		Title:               p.Title,
		EducationLevel:      p.EducationLevel,
		EducationLevel2:     p.EducationLevel2,
		BeasiswaDescription: p.BeasiswaDescription,
		Organizer:           p.Organizer,
		Type:                p.Type,
		OpenReg:             p.OpenReg,
		CloseReg:            p.CloseReg,
		PhotoURL:            p.PhotoURL,
		BeasiswaURL:         p.BeasiswaURL,
	}
}
