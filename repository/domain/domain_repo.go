package domain

import (
	"../../models"
	"context"
)

type DomainRepo interface {
	GetS(ctx context.Context) ([]*models.Domain, error)
	GetByHost(ctx context.Context, host string) (*models.DomainServers, error)
}