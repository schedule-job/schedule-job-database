package pg

import (
	"context"
	"errors"
	"log"

	"github.com/jackc/pgx/v5"
	"github.com/schedule-job/schedule-job-database/core"
)

func (p *PostgresSQL) InsertAction(job_id, name string, payload map[string]interface{}) error {
	_, err := p.usePostgresSQL(func(client *pgx.Conn, ctx context.Context) (result interface{}, err error) {
		_, errExec := client.Exec(ctx, "INSERT INTO action (job_id, name, payload) VALUES ($1, $2, $3)", job_id, name, payload)
		if errExec != nil {
			return nil, errExec
		}

		return nil, nil
	})
	if err != nil {
		log.Fatalln(err.Error())
		return err
	}
	return nil
}

func (p *PostgresSQL) UpdateAction(job_id, name string, payload map[string]interface{}) error {
	return p.InsertAction(job_id, name, payload)
}

func (p *PostgresSQL) DeleteAction(job_id string) error {
	_, err := p.usePostgresSQL(func(client *pgx.Conn, ctx context.Context) (result interface{}, err error) {
		_, errExec := client.Exec(ctx, "DELETE FROM action WHERE job_id = $1", job_id)
		if errExec != nil {
			return nil, errExec
		}
		return nil, nil
	})
	if err != nil {
		log.Fatalln(err.Error())
		return err
	}
	return nil
}

func (p *PostgresSQL) SelectAction(job_id string) (*core.FullAction, error) {
	info := core.FullAction{}
	_, err := p.usePostgresSQL(func(client *pgx.Conn, ctx context.Context) (result interface{}, err error) {
		errQuery := client.QueryRow(ctx, "SELECT job_id, name, payload FROM action WHERE job_id = $1 ORDER BY created_at desc", job_id).Scan(
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
		log.Fatalln(err.Error())
		return nil, err
	}

	return &info, nil
}

func (p *PostgresSQL) SelectIdsByAction() ([]string, error) {
	data, err := p.usePostgresSQL(func(client *pgx.Conn, ctx context.Context) (result interface{}, err error) {
		rows, queryErr := client.Query(ctx, "SELECT job_id FROM action ORDER BY created_at desc")
		if queryErr != nil {
			return nil, queryErr
		}
		ids := []string{}
		for rows.Next() {
			id := ""
			scanErr := rows.Scan(&id)
			ids = append(ids, id)
			if scanErr != nil {
				continue
			}
		}
		return ids, nil
	})

	if err != nil {
		return nil, err
	}

	ids, check := data.([]string)

	if !check {
		return nil, errors.New("failed")
	}

	return ids, nil
}

func (p *PostgresSQL) SelectActions() (*[]core.FullAction, error) {
	data, err := p.usePostgresSQL(func(client *pgx.Conn, ctx context.Context) (result interface{}, err error) {
		rows, queryErr := client.Query(ctx, "SELECT job_id, name, payload FROM action ORDER BY created_at desc")
		if queryErr != nil {
			return nil, queryErr
		}
		requests := []core.FullAction{}
		for rows.Next() {
			request := core.FullAction{}
			scanErr := rows.Scan(&request.JobId,
				&request.Name,
				&request.Payload)
			requests = append(requests, request)
			if scanErr != nil {
				continue
			}
		}
		return requests, nil
	})

	if err != nil {
		return nil, err
	}

	requests, check := data.([]core.FullAction)

	if !check {
		return nil, errors.New("failed")
	}

	return &requests, nil
}
