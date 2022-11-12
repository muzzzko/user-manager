package postgres

import (
	"context"
	"github.com/go-openapi/strfmt"
	"github/user-manager/internal/entity"
	"log"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/vrischmann/envconfig"

	"github/user-manager/config"
)

var conn *pgxpool.Pool

func init() {
	cfg := config.Service{}
	if err := envconfig.InitWithPrefix(&cfg, "user_manager"); err != nil {
		panic(err)
	}

	pgxConfig, err := pgxpool.ParseConfig(cfg.Postgres.Master)
	if err != nil {
		log.Fatalf("fail to parse postgres master config: %s", err.Error())
	}
	conn, err = pgxpool.New(context.Background(), pgxConfig.ConnString())
	if err != nil {
		log.Fatalf("fail to connect to postgres master: %s", err.Error())
	}
	err = conn.Ping(context.Background())
	if err != nil {
		log.Fatalf("fail to ping postgres master: %s", err.Error())
	}
}

func TruncateUserProfile() error {
	query := `
	truncate user_profile
`

	_, err := conn.Exec(context.Background(), query)

	return err
}

func GetUserByID(id strfmt.UUID) (entity.User, error) {
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
	join 
		country c on u.country_id = c.id
	where u.id = $1
`

	user := entity.User{}
	country := entity.Country{}
	res := conn.QueryRow(context.Background(), query, id)
	err := res.Scan(
		&user.ID,
		&user.FirstName,
		&user.LastName,
		&user.Nickname,
		&user.Email,
		&user.PasswordHash,
		&country.ID,
		&country.Code,
	)

	user.Country = country

	return user, err
}
