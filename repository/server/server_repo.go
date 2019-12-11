package server

import (
 	"../../models"
	"context"
)

type ServerRepo interface {
	GetSByDomainId(ctx context.Context, host string) ([]*models.Server, error)
}