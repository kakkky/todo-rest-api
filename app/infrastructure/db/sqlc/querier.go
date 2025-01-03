// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.27.0

package sqlc

import (
	"context"
)

type Querier interface {
	DeleteTask(ctx context.Context, id string) error
	DeleteUser(ctx context.Context, id string) error
	FetchAllUser(ctx context.Context) ([]FetchAllUserRow, error)
	FetchTaskById(ctx context.Context, id string) (FetchTaskByIdRow, error)
	FetchTasks(ctx context.Context) ([]FetchTasksRow, error)
	FetchUserTasks(ctx context.Context, userID string) ([]FetchUserTasksRow, error)
	FindTaskById(ctx context.Context, id string) (FindTaskByIdRow, error)
	FindUserByEmail(ctx context.Context, email string) (FindUserByEmailRow, error)
	FindUserById(ctx context.Context, id string) (FindUserByIdRow, error)
	InsertTask(ctx context.Context, arg InsertTaskParams) error
	InsertUser(ctx context.Context, arg InsertUserParams) error
	UpdateTask(ctx context.Context, arg UpdateTaskParams) error
	UpdateUser(ctx context.Context, arg UpdateUserParams) error
}

var _ Querier = (*Queries)(nil)
