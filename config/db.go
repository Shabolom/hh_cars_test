package config

import (
	"context"
	"fmt"
	"github.com/Masterminds/squirrel"
	"github.com/jackc/pgx/v4/pgxpool"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	_ "github.com/lib/pq"
)

var Sq squirrel.StatementBuilderType
var Pool *pgxpool.Pool

func InitPgSQL() error {

	// создание строки подключения
	connectionString := fmt.Sprintf("postgresql://%s:%s@%s:%s/%s?sslmode=disable",
		Env.DbUser,
		Env.DbPassword,
		Env.DbHost,
		Env.DbPort,
		Env.DbName,
	)

	pool, err := pgxpool.Connect(context.Background(), connectionString)
	if err != nil {
		return err
	}

	sqlBuilder := squirrel.StatementBuilder.PlaceholderFormat(squirrel.Dollar)
	if err != nil {
		return err
	}

	Pool = pool
	Sq = sqlBuilder

	return nil
}
