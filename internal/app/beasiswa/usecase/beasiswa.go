package usecase

type BeasiswaUsecaseItf interface {}

type BeasiswaUsecase struct {}

func NewBeasiswaUsecase() BeasiswaUsecaseItf {
    return &BeasiswaUsecase{}
}
