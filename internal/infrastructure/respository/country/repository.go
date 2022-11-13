package country

import (
	"context"
	"errors"
	"fmt"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"

	"github/user-manager/internal/entity"
	errorpkg "github/user-manager/internal/error"
)

type Repository struct {
	slave *pgxpool.Pool
}

func NewRepository(slave *pgxpool.Pool) *Repository {
	return &Repository{
		slave: slave,
	}
}

func (r *Repository) GetCountryByCode(ctx context.Context, code string) (entity.Country, error) {
	query := `
	select 
	    id, code 
	from
	    country
	where code = $1
`

	res := r.slave.QueryRow(ctx, query, code)

	country := entity.Country{}
	err := res.Scan(&country.ID, &country.Code)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return country, fmt.Errorf("country not found: %w", errorpkg.CountryNotFound)
		}

		return country, fmt.Errorf("fail to get country: %w", err)
	}

	return country, nil
}
