package usecase

import (
	"mycareerapp/internal/app/beasiswa/repository"
	"mycareerapp/internal/domain/dto"
	"mycareerapp/internal/domain/entity"

	"github.com/google/uuid"
)

type BeasiswaUsecaseItf interface {
	GetAllBeasiswa() (*[]dto.ResponseGetBeasiswa, error)
	CreateBeasiswa(request dto.RequestCreateBeasiswa) (dto.ResponseCreateBeasiswa, error)
	GetSpecificBeasiswa(BeasiswaID uuid.UUID) (dto.ResponseGetBeasiswa, error)
	UpdateBeasiswa(beasiswaID uuid.UUID, request dto.RequestUpdateBeasiswa) error
	DeleteBeasiswa(beasiswaID uuid.UUID) error
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

func (u BeasiswaUsecase) CreateBeasiswa(request dto.RequestCreateBeasiswa) (dto.ResponseCreateBeasiswa, error) {

	beasiswa := entity.Beasiswa{
		ID:                  uuid.New(),
		Title:               request.Title,
		EducationLevel:      request.EducationLevel,
		EducationLevel2:     request.EducationLevel2,
		BeasiswaDescription: request.BeasiswaDescription,
		Organizer:           request.Organizer,
		Type:                request.Type,
		OpenReg:             request.OpenReg,
		CloseReg:            request.CloseReg,
		PhotoURL:            request.PhotoURL,
		BeasiswaURL:         request.BeasiswaURL,
	}

	err := u.BeasiswaRepository.Create(&beasiswa)
	if err != nil {
		return dto.ResponseCreateBeasiswa{}, err
	}

	return beasiswa.ParseToDTO(), nil
}

func (u BeasiswaUsecase) GetSpecificBeasiswa(BeasiswaID uuid.UUID) (dto.ResponseGetBeasiswa, error) {

	beasiswa := &entity.Beasiswa{
		ID: BeasiswaID,
	}

	err := u.BeasiswaRepository.GetSpecificBeasiswa(beasiswa)
	if err != nil {
		return dto.ResponseGetBeasiswa{}, err
	}

	return beasiswa.ParseToDTOGet(), err
}

func (u BeasiswaUsecase) UpdateBeasiswa(beasiswaID uuid.UUID, request dto.RequestUpdateBeasiswa) error {

	beasiswa := &entity.Beasiswa{
		ID:                  beasiswaID,
		Title:               request.Title,
		EducationLevel:      request.EducationLevel,
		EducationLevel2:     request.EducationLevel2,
		BeasiswaDescription: request.BeasiswaDescription,
		Organizer:           request.Organizer,
		Type:                request.Type,
		OpenReg:             request.OpenReg,
		CloseReg:            request.CloseReg,
		PhotoURL:            request.PhotoURL,
		BeasiswaURL:         request.BeasiswaURL,
	}

	err := u.BeasiswaRepository.Update(beasiswa)
	if err != nil {
		return err
	}

	return nil
}

func (u BeasiswaUsecase) DeleteBeasiswa(beasiswaID uuid.UUID) error {

	beasiswa := &entity.Beasiswa{
		ID: beasiswaID,
	}

	return u.BeasiswaRepository.Delete(beasiswa)
}
