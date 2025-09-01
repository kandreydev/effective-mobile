package repository

import (
	"context"
	"time"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/kandreydev/effective-mobile/internal/models"
	"github.com/pkg/errors"
)

type SubscriptionsRepo struct {
	pool *pgxpool.Pool
}

func NewSubscriptionsRepo(pool *pgxpool.Pool) *SubscriptionsRepo {
	return &SubscriptionsRepo{pool: pool}
}

func (r *SubscriptionsRepo) ListSubscription(ctx context.Context) ([]models.Subscription, error) {
	rows, err := r.pool.Query(ctx,
		`SELECT id, service_name, price, user_id, start_date, end_date FROM subscriptions`,
	)
	if err != nil {
		return nil, errors.Wrap(err, "failed to query ListSubscription")
	}

	defer rows.Close()

	subscriptions := []models.Subscription{}

	for rows.Next() {
		subscription := models.Subscription{}

		err := rows.Scan(&subscription.ID, &subscription.ServiceName, &subscription.Price, &subscription.UserID, &subscription.StartDate, &subscription.EndDate)
		if err != nil {
			return nil, errors.Wrap(err, "failed to scan row in ListSubscription")
		}

		if subscription.StartDate != nil {
			subscription.StartDateString = subscription.StartDate.Format("01-2006")
		}

		if subscription.EndDate != nil {
			subscription.EndDateString = subscription.EndDate.Format("01-2006")
		}

		subscriptions = append(subscriptions, subscription)
	}

	if err := rows.Err(); err != nil {
		return nil, errors.Wrap(err, "rows iteration error in ListSubscription")
	}

	return subscriptions, nil
}

func (r *SubscriptionsRepo) GetSubscription(ctx context.Context, id string) (*models.Subscription, error) {
	var sub models.Subscription

	query := `SELECT id, service_name, price, user_id, start_date, end_date FROM subscriptions WHERE id = $1`

	err := r.pool.QueryRow(ctx, query, id).Scan(
		&sub.ID, &sub.ServiceName, &sub.Price, &sub.UserID, &sub.StartDate, &sub.EndDate,
	)
	if err != nil {
		return nil, errors.Wrap(err, "failed to get subscription from storagr")
	}

	if sub.StartDate != nil {
		sub.StartDateString = sub.StartDate.Format("01-2006")
	}

	if sub.EndDate != nil {
		sub.EndDateString = sub.EndDate.Format("01-2006")
	}

	return &sub, nil
}

func (r *SubscriptionsRepo) CreateSubscription(ctx context.Context, input models.SubscriptionInput) (*models.Subscription, error) {
	var sub models.Subscription

	query := `
		INSERT INTO subscriptions (service_name, price, user_id, start_date, end_date)
		VALUES ($1, $2, $3, $4, $5)
		RETURNING id, service_name, price, user_id, start_date, end_date
	`

	startDate, err := time.Parse("01-2006", input.StartDateString)
	if err != nil {
		return nil, errors.Wrap(err, "failed to parse start date")
	} else {
		input.StartDate = &startDate
	}

	if input.EndDateString == "" {
		input.EndDate = nil
	} else {
		endDate, err := time.Parse("01-2006", input.EndDateString)
		if err != nil {
			return nil, errors.Wrap(err, "failed to parse end date")
		} else {
			input.EndDate = &endDate
		}
	}

	err = r.pool.QueryRow(ctx, query,
		input.ServiceName,
		input.Price,
		input.UserID,
		input.StartDate,
		input.EndDate,
	).Scan(&sub.ID, &sub.ServiceName, &sub.Price, &sub.UserID, &sub.StartDate, &sub.EndDate)
	if err != nil {
		return nil, errors.Wrap(err, "failed to create subscription")
	}

	if sub.StartDate != nil {
		sub.StartDateString = sub.StartDate.Format("01-2006")
	}

	if sub.EndDate != nil {
		sub.EndDateString = sub.EndDate.Format("01-2006")
	}

	return &sub, nil
}

func (r *SubscriptionsRepo) UpdateSubscription(ctx context.Context, id string, input models.SubscriptionInput) (*models.Subscription, error) {
	var sub models.Subscription

	query := `
		UPDATE subscriptions
		SET service_name = $1, price = $2, user_id = $3, start_date = $4, end_date = $5, updated_at = now()
		WHERE id = $6
		RETURNING id, service_name, price, user_id, start_date, end_date
	`

	err := r.pool.QueryRow(ctx, query,
		input.ServiceName,
		input.Price,
		input.UserID,
		input.StartDate,
		input.EndDate,
		id,
	).Scan(&sub.ID, &sub.ServiceName, &sub.Price, &sub.UserID, &sub.StartDate, &sub.EndDate)
	if err != nil {
		return nil, errors.Wrap(err, "failed to update subscription")
	}

	return &sub, nil
}

func (r *SubscriptionsRepo) DeleteSubscription(ctx context.Context, id string) error {
	_, err := r.pool.Exec(ctx, `DELETE FROM subscriptions WHERE id = $1`, id)
	if err != nil {
		return errors.Wrap(err, "failed to delete subscription")
	}

	return nil
}

func (r *SubscriptionsRepo) CalculateTotalCost(ctx context.Context, userID, startDate, endDate string) (int, error) {
	var totalCost int

	query := `
		SELECT SUM(price) FROM subscriptions
		WHERE user_id = $1 AND start_date >= $2 AND end_date <= $3
	`

	err := r.pool.QueryRow(ctx, query, userID, startDate, endDate).Scan(&totalCost)
	if err != nil {
		return 0, errors.Wrap(err, "failed to calculate total cost")
	}

	return totalCost, nil
}
