package user

import (
	"context"
	"fmt"
	"github.com/go-openapi/strfmt"
	"github.com/jackc/pgx/v5/pgconn"
	"github.com/jackc/pgx/v5/pgxpool"
	errorpkg "github/user-manager/internal/error"

	"github/user-manager/internal/entity"
)

const (
	duplicateKeyErrorCode = "23505"
)

type Repository struct {
	master *pgxpool.Pool
	slave  *pgxpool.Pool
}

func NewRepository(master *pgxpool.Pool, slave *pgxpool.Pool) *Repository {
	return &Repository{
		master: master,
		slave:  slave,
	}
}

func (r *Repository) Save(ctx context.Context, user entity.User) error {
	query := `
	insert into 
	    user_profile (id, first_name, last_name, nickname, password_hash, email, country_id) 
	values ($1, $2, $3, $4, $5, $6, $7)
`

	_, err := r.master.Exec(ctx,
		query,
		user.ID,
		user.FirstName,
		user.LastName,
		user.Nickname,
		user.PasswordHash,
		user.Email,
		user.Country.ID,
	)
	if err != nil {
		if pgerr, ok := err.(*pgconn.PgError); ok {
			if pgerr.Code == duplicateKeyErrorCode {
				return errorpkg.UserAlreadyExists
			}
		}

		return fmt.Errorf("fail save user: %w", err)
	}

	return nil
}

func (r *Repository) Delete(ctx context.Context, userID strfmt.UUID) error {
	query := `
	delete from 
	    user_profile
	where 
	    id = $1
`
	_, err := r.master.Exec(ctx, query, userID)
	if err != nil {
		return fmt.Errorf("fail to delete user: %w", err)
	}

	return nil
}

func (r *Repository) Update(ctx context.Context, user entity.User) error {
	query := `
	update 
		user_profile
	set
	    first_name = $1,
		last_name = $2,
		nickname = $3,
	    password_hash = $4,
	    email = $5,
	    country_id = $6
	where
	    id = $7
`

	res, err := r.master.Exec(ctx,
		query,
		user.FirstName,
		user.LastName,
		user.Nickname,
		user.PasswordHash,
		user.Email,
		user.Country.ID,
		user.ID,
	)

	if err != nil {
		if pgerr, ok := err.(*pgconn.PgError); ok {
			if pgerr.Code == duplicateKeyErrorCode {
				return errorpkg.UserAlreadyExists
			}
		}

		return fmt.Errorf("fail to update user: %w", err)
	}

	if res.RowsAffected() == 0 {
		return errorpkg.UserNotFound
	}

	return nil
}
