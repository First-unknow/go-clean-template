package repository

import (
	"fmt"
	"log"

	models "innovasive/go-clean-template/models"
	"innovasive/go-clean-template/orm"
	"innovasive/go-clean-template/service/user"

	"github.com/jmoiron/sqlx"
)

type psqlUserRepository struct {
	db *sqlx.DB
}

func NewPsqlUserRepository(dbcon *sqlx.DB) user.PsqlUserRepositoryInf {
	return &psqlUserRepository{
		db: dbcon,
	}
}

func (p psqlUserRepository) FetchAll() ([]*models.User, error) {
	sql := fmt.Sprintf(`
		SELECT 
			%s
		FROM users
	`,
		models.UserSelector,
	)

	log.Println(sql)

	rows, err := p.db.Queryx(sql)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	users, err := p.orm(rows, nil)
	if err != nil {
		return nil, err
	}

	return users, nil
}

func (p psqlUserRepository) orm(rows *sqlx.Rows, joinField []string) ([]*models.User, error) {
	var users = make([]*models.User, 0)

	for rows.Next() {
		var user = new(models.User)
		user, err := orm.OrmUser(user, rows, joinField)
		if err != nil {
			return nil, err
		}
		if user != nil {
			users = append(users, user)
		}
	}
	return users, nil
}
