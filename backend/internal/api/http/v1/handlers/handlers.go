package handlers

import (
	"backend/internal/usecase"
	"github.com/gin-gonic/gin"
	"strconv"
)

type Handler struct {
	uc usecase.UseCase
}

func NewHandler(uc usecase.UseCase) *Handler {
	return &Handler{uc: uc}
}

func (h *Handler) GetPoll(ctx *gin.Context) {
	id := ctx.Param("id")
	intID, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		ctx.JSON(500, gin.H{"status": "error"})
		return
	}

	poll, err := h.uc.GetPollById(int(intID))
	if err != nil {
		ctx.JSON(500, gin.H{"status": "error"})
		return
	}

	ctx.JSON(200, poll)
}

func (h *Handler) VoteForPoll(ctx *gin.Context) {
	id := ctx.Param("id")
	intID, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		ctx.JSON(500, gin.H{"status": "error"})
		return
	}

	err = h.uc.VoteForOption(int(intID))
	if err != nil {
		ctx.JSON(500, gin.H{"status": "error"})
		return
	}

	ctx.JSON(200, gin.H{"status": "ok"})
}
