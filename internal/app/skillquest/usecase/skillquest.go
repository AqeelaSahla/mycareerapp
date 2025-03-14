package usecase

type SkillquestUsecaseItf interface {}

type SkillquestUsecase struct {}

func NewSkillquestUsecase() SkillquestUsecaseItf {
    return &SkillquestUsecase{}
}
