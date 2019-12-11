package domain

import (
	"../../models"
	"context"
	"database/sql"
)

func NewSQLDomainRepo(Conn *sql.DB) DomainRepo {
	return &sqlDomainRepo {
		Conn: Conn,
	}
}

type sqlDomainRepo struct {
	Conn *sql.DB
}

func (m *sqlDomainRepo) fetchDomain(ctx context.Context, query string, args ...interface{}) (
	[]*models.Domain, error) {

	rows, err := m.Conn.QueryContext(ctx, query, args...)

	if err != nil {
		return nil, err
	}
	defer rows.Close()

	payload := make([]*models.Domain, 0)
	for rows.Next() {
		data := new(models.Domain)

		err := rows.Scan(
			&data.Title,
		)
		if err != nil {
			return nil, err
		}
		payload = append(payload, data)
	}
	return payload, nil
}

func (m *sqlDomainRepo) fetchDomainServers(ctx context.Context, query string, args ...interface{}) (
	[]*models.DomainServers, error) {

	rows, err := m.Conn.QueryContext(ctx, query, args...)

	if err != nil {
		return nil, err
	}
	defer rows.Close()

	payload := make([]*models.DomainServers, 0)
	for rows.Next() {
		data := new(models.DomainServers)

		err := rows.Scan(
			&data.ServersChanged,
			&data.SslGrade,
			&data.PreviousSslGrade,
			&data.Logo,
			&data.Title,
			&data.IsDown,
		)
		if err != nil {
			return nil, err
		}
		payload = append(payload, data)
	}
	return payload, nil
}

func (m *sqlDomainRepo) GetS(ctx context.Context) ([]*models.Domain, error) {
	query := "SELECT title FROM domains;"
	return m.fetchDomain(ctx, query)
}

func (m *sqlDomainRepo) GetByHost(ctx context.Context, host string) (*models.DomainServers, error) {
	query := "SELECT servers_changed, ssl_grade, previous_ssl_grade, logo, title, is_down FROM domains WHERE host=$1"

	rows, err := m.fetchDomainServers(ctx, query, host)
	if err != nil {
		return nil, err
	}

	payload := &models.DomainServers{}
	if len(rows) > 0 {
		payload = rows[0]
	} else {
		return nil, models.ErrNotFound
	}

	return payload, nil
}
