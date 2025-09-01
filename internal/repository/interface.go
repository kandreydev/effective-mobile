package repository

import (
	"context"

	"github.com/kandreydev/effective-mobile/internal/models"
)

type SubscriptionsProvider interface {
	ListSubscription(ctx context.Context) ([]models.Subscription, error)
	GetSubscription(ctx context.Context, id string) (*models.Subscription, error)
	CreateSubscription(ctx context.Context, input models.SubscriptionInput) (*models.Subscription, error)
	UpdateSubscription(ctx context.Context, id string, input models.SubscriptionInput) (*models.Subscription, error)
	DeleteSubscription(ctx context.Context, id string) error
	CalculateTotalCost(ctx context.Context, userID, startDate, endDate string) (int, error)
}
