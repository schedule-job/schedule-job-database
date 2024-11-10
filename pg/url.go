package pg

import (
	"context"

	"github.com/jackc/pgx/v5"
	schedule_errors "github.com/schedule-job/schedule-job-errors"
)

func (p *PostgresSQL) selectUrls(category string) ([]string, error) {
	urls := []string{}
	_, err := p.usePostgresSQL(func(client *pgx.Conn, ctx context.Context) (result interface{}, err error) {
		rows, queryErr := client.Query(ctx, "SELECT url FROM urls WHERE category = $1 ORDER BY created_at desc", category)
		if queryErr != nil {
			return nil, queryErr
		}

		for rows.Next() {
			var url = ""
			scanErr := rows.Scan(&url)
			urls = append(urls, url)
			if scanErr != nil {
				continue
			}
		}
		return nil, nil
	})

	if err != nil {
		return nil, &schedule_errors.QueryError{Err: err}
	}

	return urls, nil
}

func (p *PostgresSQL) SelectAgentUrls() ([]string, error) {
	return p.selectUrls("agent")
}

func (p *PostgresSQL) SelectBatchUrls() ([]string, error) {
	return p.selectUrls("batch")
}
