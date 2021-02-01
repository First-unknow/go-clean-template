package model

import "github.com/google/uuid"

const UserSelector = `
		users.id			"users.id",
		users.email			"users.email",
		users.first_name		"users.first_name",
		users.last_name 		"users.last_name"
`

type User struct {
	TableName struct{}  `json:"-" db:"users"`
	ID        uuid.UUID `json:"id" db:"id" type:"uuid"`
	Email     string    `json:"email" db:"email" type:"string"`
	Firstname string    `json:"first_name" db:"first_name" type:"string"`
	Lastname  string    `json:"last_name" db:"last_name" type:"string"`
}

type Users []*User
