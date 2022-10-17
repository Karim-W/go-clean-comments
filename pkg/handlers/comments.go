package handlers

import (
	"github.com/gin-gonic/gin"
	"github.com/karim-w/go-clean-commments/pkg/contracts"
	"github.com/karim-w/go-clean-commments/pkg/usecases"
)

func CreateComment(ctx *gin.Context) {
	var body contracts.CreateCommentContract
	if err := ctx.BindJSON(&body); err != nil {
		ctx.JSON(400, map[string]interface{}{
			"error": err.Error(),
		})
		return
	}
	status, response := usecases.CreateComment(ctx, body.Body, body.UserId)
	ctx.JSON(status, response)
}

func FetchComment(ctx *gin.Context) {
	status, response := usecases.FetchComment(ctx, ctx.Param("id"))
	ctx.JSON(status, response)
}
