package repository

import (
	"mycareerapp/internal/domain/entity"

	"gorm.io/gorm"
)

type BeasiswaMySQLItf interface {
    GetAllBeasiswa(beasiswa *[]entity.Beasiswa) error
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
