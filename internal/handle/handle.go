package handle

import (
	"log/slog"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/kandreydev/effective-mobile/internal/models"
	"github.com/kandreydev/effective-mobile/internal/repository"
)

type Handle struct {
	repo repository.SubscriptionsProvider
	log  slog.Logger
}

func New(repo repository.SubscriptionsProvider, log *slog.Logger) *Handle {
	return &Handle{
		repo: repo,
		log:  *log,
	}
}

func (h *Handle) ListSubscriptions(c *gin.Context) {
	subscriptions, err := h.repo.ListSubscription(c.Request.Context())
	if err != nil {
		h.log.Error("failed to list subscriptions", slog.String("error", err.Error()))
		Error := models.Error{
			Code:    http.StatusInternalServerError,
			Message: "failed to list subscriptions",
		}
		c.JSON(http.StatusInternalServerError, Error)

		return
	}

	h.log.Info("got list")
	c.JSON(http.StatusOK, subscriptions)
}

func (h *Handle) CreateSubscription(c *gin.Context) {
	var input models.SubscriptionInput
	if err := c.ShouldBindJSON(&input); err != nil {
		h.log.Error("failed to bind JSON", slog.String("error", err.Error()))
		Error := models.Error{
			Code:    http.StatusBadRequest,
			Message: "failed to bind JSON",
		}
		c.JSON(http.StatusBadRequest, Error)

		return
	}

	h.log.Info("input", "input", input)

	subscription, err := h.repo.CreateSubscription(c.Request.Context(), input)
	if err != nil {
		h.log.Error("failed to create subscription", slog.String("error", err.Error()))
		Error := models.Error{
			Code:    http.StatusInternalServerError,
			Message: "failed to create subscription",
		}
		c.JSON(http.StatusInternalServerError, Error)

		return
	}

	h.log.Info("created subscription", slog.String("id", subscription.ID))
	c.JSON(http.StatusCreated, subscription)
}

func (h *Handle) GetSubscription(c *gin.Context) {
	id := c.Param("id")

	subscription, err := h.repo.GetSubscription(c.Request.Context(), id)
	if err != nil {
		h.log.Error("failed to get subscription", slog.String("error", err.Error()))
		Error := models.Error{
			Code:    http.StatusInternalServerError,
			Message: "failed to get subscription",
		}
		c.JSON(http.StatusInternalServerError, Error)

		return
	}

	h.log.Info("got subscription", slog.String("id", subscription.ID))
	c.JSON(http.StatusOK, subscription)
}

func (h *Handle) UpdateSubscription(c *gin.Context) {
	h.log.Error("UpdateSubscription not implemented")
	c.JSON(http.StatusNotImplemented, gin.H{"error": "not implemented"})
}

func (h *Handle) DeleteSubscription(c *gin.Context) {
	h.log.Error("DeleteSubscription not implemented")
	c.JSON(http.StatusNotImplemented, gin.H{"error": "not implemented"})
}

func (h *Handle) CalculateTotalCost(c *gin.Context) {
	h.log.Error("CalculateTotalCost not implemented")
	c.JSON(http.StatusNotImplemented, gin.H{"error": "not implemented"})
}
