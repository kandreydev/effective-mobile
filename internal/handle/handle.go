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
	log               slog.Logger
}

func New(repo repository.SubscriptionsProvider, log *slog.Logger) *Handle {
	return &Handle{
		repo: repo,
		log:               *log,
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
	h.log.Error("CreateSubscription not implemented")
	c.JSON(http.StatusNotImplemented, gin.H{"error": "not implemented"})
}

func (h *Handle) GetSubscription(c *gin.Context) {
	h.log.Error("GetSubscription not implemented")
	c.JSON(http.StatusNotImplemented, gin.H{"error": "not implemented"})
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
