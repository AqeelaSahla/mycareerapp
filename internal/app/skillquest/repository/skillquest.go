package repository

import(
    "gorm.io/gorm"
)

type SkillquestMySQLItf interface {}

type SkillquestMySQL struct {
    db *gorm.DB
}

func NewSkillquestMySQL(db *gorm.DB) SkillquestMySQLItf {
    return &SkillquestMySQL{db}
}
