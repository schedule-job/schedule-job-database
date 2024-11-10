package pg

import (
	"context"

	"github.com/jackc/pgx/v5"
	"github.com/schedule-job/schedule-job-database/core"
	schedule_errors "github.com/schedule-job/schedule-job-errors"
)

func (p *PostgresSQL) InsertAuthorization(name string, payload interface{}) error {
	_, err := p.usePostgresSQL(func(client *pgx.Conn, ctx context.Context) (result interface{}, err error) {
		_, errExec := client.Exec(ctx, "INSERT INTO authorization (name, payload) VALUES ($1, $2)", name, payload)
		if errExec != nil {
			return nil, errExec
		}

		return nil, nil
	})

	if err != nil {
		return &schedule_errors.QueryError{Err: err}
	}

	return nil
}

func (p *PostgresSQL) UpdateAuthorization(name string, payload interface{}) error {
	return p.InsertAuthorization(name, payload)
}

func (p *PostgresSQL) DeleteAuthorization(name string) error {
	_, err := p.usePostgresSQL(func(client *pgx.Conn, ctx context.Context) (result interface{}, err error) {
		_, errExec := client.Exec(ctx, "DELETE FROM authorizations WHERE name = $1", name)
		if errExec != nil {
			return nil, errExec
		}

		return nil, nil
	})

	if err != nil {
		return &schedule_errors.QueryError{Err: err}
	}

	return nil
}

func (p *PostgresSQL) SelectAuthorizations() ([]core.FullAuthorization, error) {
	authorizations := []core.FullAuthorization{}
	_, err := p.usePostgresSQL(func(client *pgx.Conn, ctx context.Context) (result interface{}, err error) {
		rows, queryErr := client.Query(ctx, "SELECT name, payload FROM authorizations ORDER BY created_at desc")
		if queryErr != nil {
			return nil, queryErr
		}

		for rows.Next() {
			authorization := core.FullAuthorization{}
			scanErr := rows.Scan(&authorization.Name, &authorization.Payload)
			authorizations = append(authorizations, authorization)
			if scanErr != nil {
				continue
			}
		}
		return nil, nil
	})

	if err != nil {
		return nil, &schedule_errors.QueryError{Err: err}
	}

	return authorizations, nil
}
