package user

import (
	"context"
	"fmt"
	"github.com/go-openapi/strfmt"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgconn"
	"github.com/jackc/pgx/v5/pgxpool"
	"github/user-manager/internal/entity"
	errorpkg "github/user-manager/internal/error"
	"github/user-manager/internal/infrastructure/respository"
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
	return respository.PerformTransaction(ctx, r.master, func(tx pgx.Tx) error {
		query := `
			insert into 
				user_profile (id, first_name, last_name, nickname, password_hash, email, country_id) 
			values ($1, $2, $3, $4, $5, $6, $7)
`

		_, err := tx.Exec(ctx,
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

		query = `
			insert into 
				user_facets (id, tsv)
			values ($1, array_to_tsvector($2))
`
		_, err = tx.Exec(ctx,
			query,
			user.ID,
			user.GetSearchFields(),
		)
		if err != nil {
			return fmt.Errorf("fail save user facets: %w", err)
		}

		return nil
	})
}

func (r *Repository) Delete(ctx context.Context, userID strfmt.UUID) error {
	return respository.PerformTransaction(ctx, r.master, func(tx pgx.Tx) error {
		query := `
			delete from 
				user_profile
			where 
				id = $1
`
		_, err := tx.Exec(ctx, query, userID)
		if err != nil {
			return fmt.Errorf("fail to delete user: %w", err)
		}

		query = `
			delete from 
				user_facets
			where 
				id = $1
`
		_, err = tx.Exec(ctx, query, userID)
		if err != nil {
			return fmt.Errorf("fail to delete user facets: %w", err)
		}

		return nil
	})
}

func (r *Repository) Update(ctx context.Context, user entity.User) error {

	return respository.PerformTransaction(ctx, r.master, func(tx pgx.Tx) error {
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

		query = `
			update 
				user_facets
			set
				tsv = array_to_tsvector($1)
			where
				id = $2
`

		_, err = r.master.Exec(ctx,
			query,
			user.GetSearchFields(),
			user.ID,
		)

		if err != nil {
			return fmt.Errorf("fail to update user facets: %w", err)
		}

		return nil
	})
}

func (r *Repository) GetUsersByFilters(
	ctx context.Context,
	filters map[string]string,
	limit int64,
	next *string,
) ([]entity.User, error) {

	args := make(pgx.NamedArgs)
	args["limit"] = limit

	query := `
		select 
		    u.id,
			u.first_name,
			u.last_name,
			u.nickname,
			u.email,
			u.password_hash,
			c.id,
			c.code 
		from 
			user_profile u
`
	if len(filters) != 0 {
		query += `join user_facets uf on uf.id = u.id `

		for key, value := range filters {
			query += fmt.Sprintf(`AND tsv @@ @%s::tsquery `, key)
			args[key] = fmt.Sprintf("'%s:%s'", key, value)
		}
	}

	query += `join country c on u.country_id = c.id `

	if next != nil {
		query += `where u.id < @id `
		args["id"] = *next
	}

	query += `order by u.id desc limit @limit;
`

	rows, err := r.slave.Query(ctx, query, args)
	if err != nil {
		return nil, fmt.Errorf("fail to get users by filters: %w", err)
	}
	defer rows.Close()

	res := make([]entity.User, 0)
	for rows.Next() {
		user := entity.User{}
		country := entity.Country{}
		if err := rows.Scan(
			&user.ID,
			&user.FirstName,
			&user.LastName,
			&user.Nickname,
			&user.Email,
			&user.PasswordHash,
			&country.ID,
			&country.Code,
		); err != nil {
			return nil, fmt.Errorf("fail to scan user by filters: %w", err)
		}

		user.Country = country

		res = append(res, user)
	}
	if rows.Err() != nil {
		return nil, fmt.Errorf("error in rows while getting users by filters: %w", rows.Err())
	}

	return res, nil
}
