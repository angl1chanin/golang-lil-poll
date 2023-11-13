package routes

import (
	"backend/internal/api/http/v1/handlers"
	"github.com/gin-gonic/gin"
)

func Setup(router *gin.Engine, h *handlers.Handler) {
	v1 := router.Group("/v1")
	r := v1.Group("/poll")

	// /v1/poll/
	r.GET("/:id", h.GetPoll)
	r.POST("/option/vote/:id", h.VoteForPoll)
}
