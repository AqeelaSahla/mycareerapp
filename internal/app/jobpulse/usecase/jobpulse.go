package usecase

import (
	"mycareerapp/internal/app/jobpulse/repository"
	"mycareerapp/internal/domain/dto"
	"mycareerapp/internal/domain/entity"

	"github.com/google/uuid"
)

type JobpulseUsecaseItf interface {
	GetAllJobs() (*[]dto.ResponseGetJob, error)
	GetSpecificJob(jobID uuid.UUID) (dto.ResponseGetJob, error)
	CreateJob(request dto.RequestCreateJob) (dto.ResponseCreateJob, error)
	UpdateJob(jobID uuid.UUID, request dto.RequestUpdateJob) error
	DeleteJob(jobID uuid.UUID) error
}

type JobpulseUsecase struct {
	JobpulseRepository repository.JobpulseMySQLItf
}

func NewJobpulseUsecase(jobpulseRepository repository.JobpulseMySQLItf) JobpulseUsecaseItf {
	return &JobpulseUsecase{
		JobpulseRepository: jobpulseRepository,
	}
}

func (u JobpulseUsecase) GetAllJobs() (*[]dto.ResponseGetJob, error) {

	jobs := new([]entity.Jobpulse)

	err := u.JobpulseRepository.GetAllJobs(jobs)
	if err != nil {
		return nil, err
	}

	res := make([]dto.ResponseGetJob, len(*jobs))
	for i, job := range *jobs {
		res[i] = job.ParseToDTOGet()
	}

	return &res, nil
}

func (u JobpulseUsecase) CreateJob(request dto.RequestCreateJob) (dto.ResponseCreateJob, error) {

	job := entity.Jobpulse{
		ID:              uuid.New(),
		Title:           request.Title,
		Location:        request.Location,
		WorkDescription: request.WorkDescription,
		MinimumExp:      request.MinimumExp,
		Salary:          request.Salary,
		WorkingDays:     request.WorkingDays,
		WorkingHours:    request.WorkingHours,
		PhotoURL:        request.PhotoURL,
		ArticleURL:      request.ArticleURL,
	}

	err := u.JobpulseRepository.Create(&job)
	if err != nil {
		return dto.ResponseCreateJob{}, err
	}

	return job.ParseToDTO(), nil
}

func (u JobpulseUsecase) GetSpecificJob(jobID uuid.UUID) (dto.ResponseGetJob, error) {

	job := &entity.Jobpulse{
		ID: jobID,
	}

	err := u.JobpulseRepository.GetSpecificJob(job)
	if err != nil {
		return dto.ResponseGetJob{}, err
	}

	return job.ParseToDTOGet(), err
}

func (u JobpulseUsecase) UpdateJob(jobID uuid.UUID, request dto.RequestUpdateJob) error {

	job := &entity.Jobpulse{
		ID:              jobID,
		Title:           request.Title,
		Location:        request.Location,
		WorkDescription: request.WorkDescription,
		MinimumExp:      request.MinimumExp,
		Salary:          request.Salary,
		WorkingDays:     request.WorkingDays,
		WorkingHours:    request.WorkingHours,
		PhotoURL:        request.PhotoURL,
		ArticleURL:      request.ArticleURL,
	}

	err := u.JobpulseRepository.Update(job)
	if err != nil {
		return err
	}

	return nil
}

func (u JobpulseUsecase) DeleteJob(jobID uuid.UUID) error {

	job := &entity.Jobpulse{
		ID: jobID,
	}

	return u.JobpulseRepository.Delete(job)
}
