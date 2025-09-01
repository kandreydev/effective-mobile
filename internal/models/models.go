package models

import "time"

// Subscription represents a user's subscription record as it exists in the system.
type Subscription struct {
	ID          string     `json:"id"`
	ServiceName string     `json:"service_name"`
	Price       int        `json:"price"`
	UserID      string     `json:"user_id"`
	StartDate   *time.Time `json:"start_date"`
	EndDate     *time.Time `json:"end_date,omitempty"`
}

// SubscriptionInput is the model used for creating or updating a subscription.
// It omits server-generated fields like ID.
type SubscriptionInput struct {
	ServiceName string     `json:"service_name"`
	Price       int        `json:"price"`
	UserID      string     `json:"user_id"`
	StartDate   *time.Time `json:"start_date"`
	EndDate     *time.Time `json:"end_date,omitempty"`
}

// TotalCost represents the response body for the cost calculation endpoint.
type TotalCost struct {
	UserID    string     `json:"user_id"`
	Total     int        `json:"total"`
	StartDate *time.Time `json:"start_date"`
	EndDate   *time.Time `json:"end_date"`
}

// Error represents a standard API error response.
type Error struct {
	Code    int32  `json:"code"`
	Message string `json:"message"`
}
