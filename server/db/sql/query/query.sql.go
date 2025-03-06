// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.27.0
// source: query.sql

package query

import (
	"context"
)

const createUser = `-- name: CreateUser :one
INSERT INTO users (
  user_id, mail, name, hashed_password
) VALUES (
  $1, $2, $3, $4
)
RETURNING user_id, mail, name, belong, hashed_password
`

type CreateUserParams struct {
	UserID         string
	Mail           string
	Name           string
	HashedPassword string
}

func (q *Queries) CreateUser(ctx context.Context, arg CreateUserParams) (User, error) {
	row := q.db.QueryRow(ctx, createUser,
		arg.UserID,
		arg.Mail,
		arg.Name,
		arg.HashedPassword,
	)
	var i User
	err := row.Scan(
		&i.UserID,
		&i.Mail,
		&i.Name,
		&i.Belong,
		&i.HashedPassword,
	)
	return i, err
}

const getAllMessage = `-- name: GetAllMessage :many
SELECT message_id, school, x, y, message, created_at, float_time, likes from messages
`

func (q *Queries) GetAllMessage(ctx context.Context) ([]Message, error) {
	rows, err := q.db.Query(ctx, getAllMessage)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Message
	for rows.Next() {
		var i Message
		if err := rows.Scan(
			&i.MessageID,
			&i.School,
			&i.X,
			&i.Y,
			&i.Message,
			&i.CreatedAt,
			&i.FloatTime,
			&i.Likes,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const getUserByEmail = `-- name: GetUserByEmail :one
SELECT user_id, mail, name, hashed_password FROM users
WHERE mail = $1 LIMIT 1
`

type GetUserByEmailRow struct {
	UserID         string
	Mail           string
	Name           string
	HashedPassword string
}

func (q *Queries) GetUserByEmail(ctx context.Context, mail string) (GetUserByEmailRow, error) {
	row := q.db.QueryRow(ctx, getUserByEmail, mail)
	var i GetUserByEmailRow
	err := row.Scan(
		&i.UserID,
		&i.Mail,
		&i.Name,
		&i.HashedPassword,
	)
	return i, err
}

const getUserByID = `-- name: GetUserByID :one
SELECT user_id, mail, name, hashed_password FROM users
WHERE user_id = $1 LIMIT 1
`

type GetUserByIDRow struct {
	UserID         string
	Mail           string
	Name           string
	HashedPassword string
}

func (q *Queries) GetUserByID(ctx context.Context, userID string) (GetUserByIDRow, error) {
	row := q.db.QueryRow(ctx, getUserByID, userID)
	var i GetUserByIDRow
	err := row.Scan(
		&i.UserID,
		&i.Mail,
		&i.Name,
		&i.HashedPassword,
	)
	return i, err
}

const incrementLikes = `-- name: IncrementLikes :exec
UPDATE messages
  set likes = likes + 1
WHERE message_id = $1
`

func (q *Queries) IncrementLikes(ctx context.Context, messageID string) error {
	_, err := q.db.Exec(ctx, incrementLikes, messageID)
	return err
}

const updatePassword = `-- name: UpdatePassword :exec
UPDATE users
  set hashed_password = $2
WHERE user_id = $1
`

type UpdatePasswordParams struct {
	UserID         string
	HashedPassword string
}

func (q *Queries) UpdatePassword(ctx context.Context, arg UpdatePasswordParams) error {
	_, err := q.db.Exec(ctx, updatePassword, arg.UserID, arg.HashedPassword)
	return err
}

const updateUser = `-- name: UpdateUser :exec
UPDATE users
  set mail = $2,
  name = $3,
  hashed_password = $4
WHERE user_id = $1
`

type UpdateUserParams struct {
	UserID         string
	Mail           string
	Name           string
	HashedPassword string
}

func (q *Queries) UpdateUser(ctx context.Context, arg UpdateUserParams) error {
	_, err := q.db.Exec(ctx, updateUser,
		arg.UserID,
		arg.Mail,
		arg.Name,
		arg.HashedPassword,
	)
	return err
}
