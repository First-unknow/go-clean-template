package user

import (
	models "innovasive/go-clean-template/models"

	"github.com/google/uuid"
)

type PsqlUserRepositoryInf interface {
	FetchAll() ([]*models.User, error)
	FindByMail(email string) ([]*models.User, error)
	CreateUser(user *models.User) (uuid.UUID, error)
}
