package api

import (
	"net/http"
	models "webService/models"

	"github.com/gin-gonic/gin"
)

func GetCfg(ctx *gin.Context) {

	resp := models.CustResp{
		ReturnCd:  0,
		ReturnMsg: "getCfg OK",
	}

	ctx.JSON(http.StatusOK, resp)
}

func ChgCfg(ctx *gin.Context) {

	resp := models.CustResp{
		ReturnCd:  0,
		ReturnMsg: "chgCfg OK",
	}

	ctx.JSON(http.StatusOK, resp)
}
