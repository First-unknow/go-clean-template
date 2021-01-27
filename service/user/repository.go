package user

import model "innovasive/go-clean-template/models"

type PsqlUserRepositoryInf interface {
	FetchAll() ([]*model.User, error)
}
