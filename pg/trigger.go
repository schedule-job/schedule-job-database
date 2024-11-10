package pg

import (
	"context"

	"github.com/jackc/pgx/v5"
	"github.com/schedule-job/schedule-job-database/core"
	schedule_errors "github.com/schedule-job/schedule-job-errors"
)

func (p *PostgresSQL) InsertTrigger(job_id, name string, payload map[string]string) error {
	_, err := p.usePostgresSQL(func(client *pgx.Conn, ctx context.Context) (result interface{}, err error) {
		_, errExec := client.Exec(ctx, "INSERT INTO trigger (job_id, name, payload) VALUES ($1, $2, $3)", job_id, name, payload)
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

func (p *PostgresSQL) UpdateTrigger(job_id, name string, payload map[string]string) error {
	return p.InsertTrigger(job_id, name, payload)
}

func (p *PostgresSQL) DeleteTrigger(job_id string) error {
	_, err := p.usePostgresSQL(func(client *pgx.Conn, ctx context.Context) (result interface{}, err error) {
		_, errExec := client.Exec(ctx, "DELETE FROM triggers WHERE job_id = $1", job_id)
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

func (p *PostgresSQL) SelectTrigger(job_id string) (*core.FullTrigger, error) {
	info := core.FullTrigger{}
	_, err := p.usePostgresSQL(func(client *pgx.Conn, ctx context.Context) (result interface{}, err error) {
		errQuery := client.QueryRow(ctx, "SELECT job_id, name, payload FROM triggers WHERE job_id = $1 ORDER BY created_at desc", job_id).Scan(
			&info.JobId,
			&info.Name,
			&info.Payload,
		)

		if errQuery != nil {
			return nil, errQuery
		}

		return nil, nil
	})

	if err != nil {
		return nil, &schedule_errors.QueryError{Err: err}
	}

	return &info, nil
}
