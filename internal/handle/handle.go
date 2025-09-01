package handle

import (
	"log/slog"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/kandreydev/effective-mobile/internal/repository"
)

type Handle struct {
	subscriptionsRepo repository.SubscriptionsProvider
	log               slog.Logger
}

func New(subscriptionsRepo repository.SubscriptionsProvider, log *slog.Logger) *Handle {
	return &Handle{
		subscriptionsRepo: subscriptionsRepo,
		log:               *log,
	}
}

func (h *Handle) ListSubscriptions(c *gin.Context) {
	h.log.Error("ListSubscriptions not implemented")
	c.JSON(http.StatusNotImplemented, gin.H{"error": "not implemented"})
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
