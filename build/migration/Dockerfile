FROM golang:1.19

RUN go install github.com/pressly/goose/v3/cmd/goose@v3.7.0

ENV USER_MANAGER_POSTGRES_MASTER="user=user_manager password=user_manager host=pgdb port=5432 dbname=user_manager sslmode=disable"

COPY . .

ENTRYPOINT ["make", "migrate"]
