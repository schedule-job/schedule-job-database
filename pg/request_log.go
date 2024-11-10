package pg

import (
	"context"
	"time"

	"github.com/jackc/pgx/v5"
	"github.com/schedule-job/schedule-job-database/core"
	schedule_errors "github.com/schedule-job/schedule-job-errors"
)

func (p *PostgresSQL) InsertRequestLog(job_id string, payload interface{}) error {
	data := payload.(core.RequestTypePayload)
	_, err := p.usePostgresSQL(func(client *pgx.Conn, ctx context.Context) (result interface{}, err error) {
		return client.Exec(ctx, "INSERT INTO request_log (job_id, status, request_url, request_method, request_headers, request_body, response_headers, response_body, response_status_code) VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9)", job_id, data.Status, data.RequestUrl, data.RequestMethod, data.RequestHeaders, data.RequestBody, data.ResponseHeaders, data.ResponseBody, data.ResponseStatusCode)
	})
	if err != nil {
		return &schedule_errors.QueryError{Err: err}
	}
	return nil
}

func (p *PostgresSQL) SelectRequestLogs(job_id, lastId string, limit int) ([]core.RequestLog, error) {
	var lastCreated = time.Date(1990, 1, 1, 1, 0, 0, 0, time.UTC)
	if lastId != "" {
		log, err := p.SelectRequestLog(lastId, job_id)
		if err == nil {
			lastCreated = log.CreatedAt
		}
	}

	logs := []core.RequestLog{}
	_, dbErr := p.usePostgresSQL(func(client *pgx.Conn, ctx context.Context) (interface{}, error) {
		rows, queryErr := client.Query(ctx, "SELECT id, job_id, status, request_url, request_method, response_status_code, created_at FROM request_logs WHERE created_at > $1 AND job_id = $2 ORDER BY created_at LIMIT $3", lastCreated, job_id, limit)

		if queryErr != nil {
			return nil, queryErr
		}

		for rows.Next() {
			log := core.RequestLog{}
			scanErr := rows.Scan(&log.Id, &log.JobId, &log.Status, &log.RequestUrl, &log.RequestMethod, &log.ResponseStatusCode, &log.CreatedAt)
			if scanErr != nil {
				continue
			}
			logs = append(logs, log)
		}

		return logs, nil
	})

	if dbErr != nil {
		return nil, &schedule_errors.QueryError{Err: dbErr}
	}

	return logs, nil
}

func (p *PostgresSQL) SelectRequestLog(id, job_id string) (*core.FullRequestLog, error) {
	log := core.FullRequestLog{}
	_, dbErr := p.usePostgresSQL(func(client *pgx.Conn, ctx context.Context) (interface{}, error) {
		queryErr := client.QueryRow(ctx, "SELECT id, job_id, status, request_url, request_method, request_headers, request_body, response_headers, response_body, response_status_code, created_at FROM request_logs WHERE id = $1 AND job_id = $2", id, job_id).Scan(
			&log.Id,
			&log.JobId,
			&log.Status,
			&log.RequestUrl,
			&log.RequestMethod,
			&log.RequestHeaders,
			&log.RequestBody,
			&log.ResponseHeaders,
			&log.ResponseBody,
			&log.ResponseStatusCode,
			&log.CreatedAt,
		)

		if queryErr != nil {
			return nil, queryErr
		}

		return log, nil
	})

	if dbErr != nil {
		return nil, &schedule_errors.QueryError{Err: dbErr}
	}

	return &log, nil
}
