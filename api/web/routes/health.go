package routes

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type ServerStatus struct {
	ServerStatus string
}

func (r *Router) SetHealthChecker() {
	r.Engine.GET("/healthcheck", healthCheck)
}

func healthCheck(ctx *gin.Context) {
	ctx.IndentedJSON(http.StatusOK, ServerStatus{ServerStatus: "Good"})
}
