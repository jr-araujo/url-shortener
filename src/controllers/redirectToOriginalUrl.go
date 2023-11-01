package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"urlshortener.com/devgym/jr/models"
)

func (ctlr *controller) RedirectOriginalUrl(ctx *gin.Context) {
	code := ctx.Param("code")

	if code == "" {
		ctx.JSON(http.StatusNotFound, gin.H{
			"error": "Code not found",
		})
		return
	}

	var shortenUrl models.ShortenUrl

	if err := ctlr.DB.Where("code = ?", code).First(&shortenUrl).Error; err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{
			"error": "Code not found",
		})
		return
	}

	ctx.Redirect(http.StatusMovedPermanently, shortenUrl.Original)
}
