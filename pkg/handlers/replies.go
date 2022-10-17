package handlers

import (
	"github.com/gin-gonic/gin"
	"github.com/karim-w/go-clean-commments/pkg/contracts"
	"github.com/karim-w/go-clean-commments/pkg/usecases"
)

func CreateReply(ctx *gin.Context) {
	var req contracts.CreateReplyContract
	if err := ctx.BindJSON(&req); err != nil {
		ctx.JSON(400, map[string]interface{}{
			"error": err.Error(),
		})
		return
	}
	status, response := usecases.ReplyToComment(ctx, ctx.Param("id"), req.Body, req.UserId)
	ctx.JSON(status, response)
}
