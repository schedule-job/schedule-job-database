package pg

import (
	"context"
	"time"

	"github.com/jackc/pgx/v5"
	"github.com/schedule-job/schedule-job-database/core"
)

func (p *PostgresSQL) InsertRequestLog(jobID string, payload interface{}) error {
	data := payload.(core.RequestTypePayload)
	_, err := p.usePostgresSQL(func(client *pgx.Conn, ctx context.Context) (result interface{}, err error) {
		return client.Exec(ctx, "INSERT INTO request_logs (job_id, status, request_url, request_method, request_headers, request_body, response_headers, response_body, response_status_code) VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9)", jobID, data.Status, data.RequestUrl, data.RequestMethod, data.RequestHeaders, data.RequestBody, data.ResponseHeaders, data.ResponseBody, data.ResponseStatusCode)
	})
	return err
}

func (p *PostgresSQL) GetRequestLogs(jobId, lastId string, limit int) ([]core.Log, error) {
	var lastCreated = time.Date(1990, 1, 1, 1, 0, 0, 0, time.UTC)
	if lastId != "" {
		log, err := p.GetRequestLogDetail(lastId, jobId)
		if err == nil {
			lastCreated = log.CreatedAt
		}
	}

	data, dbErr := p.usePostgresSQL(func(client *pgx.Conn, ctx context.Context) (interface{}, error) {
		logs := []core.Log{}
		rows, queryErr := client.Query(ctx, "SELECT id, job_id, status, request_url, request_method, response_status_code, created_at FROM request_logs WHERE created_at > $1 AND job_id = $2 ORDER BY created_at LIMIT $3", lastCreated, jobId, limit)

		if queryErr != nil {
			return nil, queryErr
		}

		for rows.Next() {
			log := core.Log{}
			scanErr := rows.Scan(&log.Id, &log.JobId, &log.Status, &log.RequestUrl, &log.RequestMethod, &log.ResponseStatusCode, &log.CreatedAt)
			if scanErr != nil {
				continue
			}
			logs = append(logs, log)
		}

		return logs, nil
	})

	if dbErr != nil {
		return nil, dbErr
	}

	return data.([]core.Log), nil
}

func (p *PostgresSQL) GetRequestLogDetail(id, jobId string) (*core.FullLog, error) {
	log := core.FullLog{}
	_, dbErr := p.usePostgresSQL(func(client *pgx.Conn, ctx context.Context) (interface{}, error) {
		queryErr := client.QueryRow(ctx, "SELECT id, job_id, status, request_url, request_method, request_headers, request_body, response_headers, response_body, response_status_code, created_at FROM request_logs WHERE id = $1 AND job_id = $2", id, jobId).Scan(
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
		return nil, dbErr
	}

	return &log, nil
}
