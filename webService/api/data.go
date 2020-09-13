package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetData(ctx *gin.Context) {
	ctx.String(http.StatusOK, "getData OK")
}

func ChgData(ctx *gin.Context) {
	ctx.String(http.StatusOK, "chgData OK")
}
