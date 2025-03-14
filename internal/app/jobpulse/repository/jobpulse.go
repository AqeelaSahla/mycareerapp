package repository

import (
	"mycareerapp/internal/domain/entity"
	"net/http"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

type JobpulseMySQLItf interface {
	Create(job *entity.Jobpulse) error
	GetAllJobs(jobs *[]entity.Jobpulse) error
	GetSpecificJob(job *entity.Jobpulse) error
	Update(job *entity.Jobpulse) error
	Delete(job *entity.Jobpulse) error
}

type JobpulseMySQL struct {
	db *gorm.DB
}

func NewJobpulseMySQL(db *gorm.DB) JobpulseMySQLItf {
	return &JobpulseMySQL{db}
}

func (r JobpulseMySQL) GetAllJobs(jobs *[]entity.Jobpulse) error {

	return r.db.Debug().Find(jobs).Error
}

func (r JobpulseMySQL) Create(job *entity.Jobpulse) error {

	return r.db.Debug().Create(job).Error
}

func (r JobpulseMySQL) GetSpecificJob(job *entity.Jobpulse) error {

	return r.db.Debug().First(job).Error
}

func (r JobpulseMySQL) Update(job *entity.Jobpulse) error {

	return r.db.Debug().Updates(job).Error
}

func (r JobpulseMySQL) Delete(job *entity.Jobpulse) error {

	q := r.db.Debug().Delete(job).RowsAffected

	if q == 0 {
		return fiber.NewError(http.StatusNotFound, "job not found")
	}

	return nil
}
