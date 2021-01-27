package user

import models "innovasive/go-clean-template/models"

type UserUsecaseInf interface {
	FetchAll() ([]*models.User, error)
}
