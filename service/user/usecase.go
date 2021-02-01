package user

import (
	models "innovasive/go-clean-template/models"

	"github.com/google/uuid"
)

type UserUsecaseInf interface {
	FetchAll() ([]*models.User, error)
	CreateUser(user *models.User) (uuid.UUID, error)
}
