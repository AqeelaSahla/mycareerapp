package entity

import (
	"mycareerapp/internal/domain/dto"
	"time"

	"github.com/google/uuid"
)

type Jobpulse struct {
	ID              uuid.UUID `gorm:"type:char(36);primaryKey"`
	Title           string    `gorm:"type:varchar(100);not null"`
	Location        string    `gorm:"type:varchar(100);not null"`
	WorkDescription string    `gorm:"type:text"`
	MinimumExp      string    `gorm:"text"`
	Salary          string    `gorm:"type:text;not null"`
	WorkingDays     string    `gorm:"type:text;not null"`
	WorkingHours    string    `gorm:"type:text;not null"`
	PhotoURL        string    `gorm:"type:text"`
	ArticleURL      string    `gorm:"type:text;not null"`
	CreatedAt       time.Time `gorm:"type:timestamp;autoCreateTime"`
}

func (p Jobpulse) ParseToDTO() dto.ResponseCreateJob {
	return dto.ResponseCreateJob{
		Title:           p.Title,
		Location:        p.Location,
		WorkDescription: p.WorkDescription,
		MinimumExp:      p.MinimumExp,
		Salary:          p.Salary,
		WorkingDays:     p.WorkingDays,
		WorkingHours:    p.WorkingHours,
		PhotoURL:        p.PhotoURL,
		ArticleURL:      p.ArticleURL,
	}
}

func (p Jobpulse) ParseToDTOGet() dto.ResponseGetJob {
	return dto.ResponseGetJob{
		ID:              p.ID,
		Title:           p.Title,
		Location:        p.Location,
		WorkDescription: p.WorkDescription,
		MinimumExp:      p.MinimumExp,
		Salary:          p.Salary,
		WorkingDays:     p.WorkingDays,
		WorkingHours:    p.WorkingHours,
		PhotoURL:        p.PhotoURL,
		ArticleURL:      p.ArticleURL,
	}
}
