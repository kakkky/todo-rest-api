// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.27.0
// source: user.sql

package sqlc

import (
	"context"
)

const deleteUser = `-- name: DeleteUser :exec
delete from users
where id=?
`

func (q *Queries) DeleteUser(ctx context.Context, id string) error {
	_, err := q.db.ExecContext(ctx, deleteUser, id)
	return err
}

const fetchAllUser = `-- name: FetchAllUser :many
select id,email,name,hashed_password
from users
`

type FetchAllUserRow struct {
	ID             string
	Email          string
	Name           string
	HashedPassword string
}

func (q *Queries) FetchAllUser(ctx context.Context) ([]FetchAllUserRow, error) {
	rows, err := q.db.QueryContext(ctx, fetchAllUser)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []FetchAllUserRow{}
	for rows.Next() {
		var i FetchAllUserRow
		if err := rows.Scan(
			&i.ID,
			&i.Email,
			&i.Name,
			&i.HashedPassword,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const findUserByEmail = `-- name: FindUserByEmail :one
select id,email,name,hashed_password
from users
where email = ?
`

type FindUserByEmailRow struct {
	ID             string
	Email          string
	Name           string
	HashedPassword string
}

func (q *Queries) FindUserByEmail(ctx context.Context, email string) (FindUserByEmailRow, error) {
	row := q.db.QueryRowContext(ctx, findUserByEmail, email)
	var i FindUserByEmailRow
	err := row.Scan(
		&i.ID,
		&i.Email,
		&i.Name,
		&i.HashedPassword,
	)
	return i, err
}

const insertUser = `-- name: InsertUser :exec
insert into users (
    id,
    name,
    email,
    hashed_password
) values (
    ?,
    ?,
    ?,
    ?
)
`

type InsertUserParams struct {
	ID             string
	Name           string
	Email          string
	HashedPassword string
}

func (q *Queries) InsertUser(ctx context.Context, arg InsertUserParams) error {
	_, err := q.db.ExecContext(ctx, insertUser,
		arg.ID,
		arg.Name,
		arg.Email,
		arg.HashedPassword,
	)
	return err
}

const updateUser = `-- name: UpdateUser :exec
update users
set name=?,email=?
where id=?
`

type UpdateUserParams struct {
	Name  string
	Email string
	ID    string
}

func (q *Queries) UpdateUser(ctx context.Context, arg UpdateUserParams) error {
	_, err := q.db.ExecContext(ctx, updateUser, arg.Name, arg.Email, arg.ID)
	return err
}
