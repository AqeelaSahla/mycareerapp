package repository

import (
	"mycareerapp/internal/domain/entity"
	"net/http"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

type BeasiswaMySQLItf interface {
	GetAllBeasiswa(bea *[]entity.Beasiswa) error
	Create(beasiswa *entity.Beasiswa) error
	GetSpecificBeasiswa(beasiswa *entity.Beasiswa) error
	Update(beasiswa *entity.Beasiswa) error
	Delete(beasiswa *entity.Beasiswa) error
}

type BeasiswaMySQL struct {
	db *gorm.DB
}

func NewBeasiswaMySQL(db *gorm.DB) BeasiswaMySQLItf {
	return &BeasiswaMySQL{db}
}

func (r BeasiswaMySQL) GetAllBeasiswa(bea *[]entity.Beasiswa) error {

	return r.db.Debug().Find(bea).Error
}

func (r BeasiswaMySQL) Create(beasiswa *entity.Beasiswa) error {

	return r.db.Debug().Create(beasiswa).Error
}

func (r BeasiswaMySQL) GetSpecificBeasiswa(beasiswa *entity.Beasiswa) error {

	return r.db.Debug().First(beasiswa).Error
}

func (r BeasiswaMySQL) Update(beasiswa *entity.Beasiswa) error {

	return r.db.Debug().Updates(beasiswa).Error
}

func (r BeasiswaMySQL) Delete(beasiswa *entity.Beasiswa) error {

	q := r.db.Debug().Delete(beasiswa).RowsAffected

	if q == 0 {
		return fiber.NewError(http.StatusNotFound, "Beasiswa not found")
	}

	return nil
}
