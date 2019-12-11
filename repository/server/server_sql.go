package server

import (
"../../models"
"context"
"database/sql"
)

func NewSQLServerRepo(Conn *sql.DB) ServerRepo {
	return &sqlServerRepo {
		Conn: Conn,
	}
}

type sqlServerRepo struct {
	Conn *sql.DB
}

func (m *sqlServerRepo) fetch(ctx context.Context, query string, args ...interface{}) ([]*models.Server, error) {
	rows, err := m.Conn.QueryContext(ctx, query, args...)

	if err != nil {
		return nil, err
	}
	defer rows.Close()

	payload := make([]*models.Server, 0)
	for rows.Next() {
		data := new(models.Server)

		err := rows.Scan(
			&data.Address,
			&data.SslGrade,
			&data.Country,
			&data.Owner,
		)
		if err != nil {
			return nil, err
		}
		payload = append(payload, data)
	}
	return payload, nil
}

func (m *sqlServerRepo) GetSByDomainId(ctx context.Context, host string) ([]*models.Server, error) {
	query := "SELECT address, ssl_grade, country, owner FROM servers WHERE domain_id=(SELECT id FROM domains WHERE host=$1)"
	return m.fetch(ctx, query, host)
}
