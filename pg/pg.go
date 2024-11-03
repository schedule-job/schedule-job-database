package pg

import (
	"context"

	"github.com/jackc/pgx/v5"
	"github.com/schedule-job/schedule-job-database/core"
)

type PostgresSQL struct {
	core.Database
	Dsn string
}

func (p *PostgresSQL) usePostgresSQL(callback func(client *pgx.Conn, ctx context.Context) (result interface{}, err error)) (result interface{}, err error) {
	ctx := context.Background()
	conn, err := pgx.Connect(ctx, p.Dsn)

	if err != nil {
		panic(err)
	}

	defer conn.Close(ctx)

	return callback(conn, ctx)
}
