package config

type Service struct {
	Server   Server
	Postgres Postgres
	Kafka    Kafka

	Environment string
}
