package repsitory

import (
	"context"

	"github.com/ishvaram/betal-kamailio/models"
)

// SubscriberRepo explain...
type SubscriberRepo interface {
	Fetch(ctx context.Context, num int64) ([]*models.Subscriber, error)
	GetByID(ctx context.Context, id int64) (*models.Subscriber, error)
	Create(ctx context.Context, p *models.Subscriber) (int64, error)
	Update(ctx context.Context, p *models.Subscriber) (*models.Subscriber, error)
	Delete(ctx context.Context, id int64) (bool, error)
}
