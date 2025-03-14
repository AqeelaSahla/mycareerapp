package usecase

import (
	"mycareerapp/internal/app/beasiswa/repository"
	"mycareerapp/internal/domain/dto"
	"mycareerapp/internal/domain/entity"
)

type BeasiswaUsecaseItf interface {
	GetAllBeasiswa() (*[]dto.ResponseGetBeasiswa, error)
}

type BeasiswaUsecase struct {
	BeasiswaRepository repository.BeasiswaMySQLItf
}

func NewBeasiswaUsecase(beasiswaRepository repository.BeasiswaMySQLItf) BeasiswaUsecaseItf {
	return &BeasiswaUsecase{
		BeasiswaRepository: beasiswaRepository,
	}
}

func (u BeasiswaUsecase) GetAllBeasiswa() (*[]dto.ResponseGetBeasiswa, error) {

	bea := new([]entity.Beasiswa)

	err := u.BeasiswaRepository.GetAllBeasiswa(bea)
	if err != nil {
		return nil, err
	}

	res := make([]dto.ResponseGetBeasiswa, len(*bea))
	for i, beasiswa := range *bea {
		res[i] = beasiswa.ParseToDTOGet()
	}

	return &res, nil
}
