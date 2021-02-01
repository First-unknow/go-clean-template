package usecase

import (
	models "innovasive/go-clean-template/models"
	"innovasive/go-clean-template/service/user"

	"github.com/google/uuid"
)

type userUsecase struct {
	psqlUserRepo user.PsqlUserRepositoryInf
}

func NewUserUsecase(uRepo user.PsqlUserRepositoryInf) user.UserUsecaseInf {
	return &userUsecase{
		psqlUserRepo: uRepo,
	}
}

func (u userUsecase) FetchAll() ([]*models.User, error) {
	return u.psqlUserRepo.FetchAll()
}

func (u userUsecase) CreateUser(user *models.User) (uuid.UUID, error) {
	return u.psqlUserRepo.CreateUser(user)
}
