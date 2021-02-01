package repository

import (
	"fmt"
	"log"

	models "innovasive/go-clean-template/models"
	"innovasive/go-clean-template/orm"
	"innovasive/go-clean-template/service/user"

	"github.com/google/uuid"
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

func (p psqlUserRepository) FindByMail(email string) ([]*models.User, error) {
	sql := fmt.Sprintf(`
		SELECT 
			%s
		FROM users
		WHERE users.email='%s'
	`,
		models.UserSelector,
		email,
	)

	log.Println(sql)

	rows, err := p.db.Queryx(sql)
	log.Println(rows)
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

func (p psqlUserRepository) CreateUser(user *models.User) (uuid.UUID, error) {
	var userID uuid.UUID
	tx, err := p.db.Begin()
	if err != nil {
		return uuid.Nil, err
	}
	sql := `
	INSERT INTO users(email,first_name,last_name)
	VALUES ($1::text, $2::text, $3::text)
	RETURNING users.id;
	`
	log.Println(sql)
	stmt, err := tx.Prepare(sql)
	if err != nil {
		return uuid.Nil, err
	}
	defer stmt.Close()
	err = stmt.QueryRow(user.Email, user.Firstname, user.Lastname).Scan(&userID)

	if err != nil {
		return uuid.Nil, err
	}
	return userID, nil
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
