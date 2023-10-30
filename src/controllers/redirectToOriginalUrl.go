package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func (ctlr *controller) RedirectOriginalUrl(ctx *gin.Context) {
	code := ctx.Param("code")

	if code == "" {
		ctx.JSON(http.StatusNotFound, gin.H{
			"error": "Code not found", //post.ErrIdEmpty,
		})
		return
	}

	ctx.Redirect(http.StatusMovedPermanently, "http://www.google.com/")
}
