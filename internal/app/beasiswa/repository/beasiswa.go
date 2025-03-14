package repository

import(
    "gorm.io/gorm"
)

type BeasiswaMySQLItf interface {}

type BeasiswaMySQL struct {
    db *gorm.DB
}

func NewBeasiswaMySQL(db *gorm.DB) BeasiswaMySQLItf {
    return &BeasiswaMySQL{db}
}
