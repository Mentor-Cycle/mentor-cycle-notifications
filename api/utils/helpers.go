package helpers

import "github.com/gin-gonic/gin"

func GenerateCtxError(ctx *gin.Context, message string, errorStatus int) {
	ctx.JSON(errorStatus, gin.H{"error": message})
	return
}
