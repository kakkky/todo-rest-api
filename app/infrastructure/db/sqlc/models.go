// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.27.0

package sqlc

import (
	"database/sql"
)

type User struct {
	ID             string
	Email          string
	Name           string
	HashedPassword string
	CreatedAt      sql.NullTime
	UpdatedAt      sql.NullTime
}