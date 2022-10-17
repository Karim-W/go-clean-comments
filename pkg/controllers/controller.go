package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/karim-w/go-clean-commments/pkg/handlers"
)

func HandleRequests(r *gin.Engine) {
	r.POST(endpoint_COMMENTS, handlers.CreateComment)
	r.GET(endpoint_COMMENT, handlers.FetchComment)
	r.POST(endpoint_REPLIES, handlers.CreateReply)
}
