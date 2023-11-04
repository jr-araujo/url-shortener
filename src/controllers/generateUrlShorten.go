package controllers

import (
	"crypto/sha256"
	"fmt"
	"math/big"
	"net/http"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/itchyny/base58-go"
	"urlshortener.com/devgym/jr/models"
)

func (ctlr *controller) GenerateUrlShorten(ctx *gin.Context) {
	var input models.InputUrl

	if err := ctx.ShouldBindJSON(&input); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if input.Url == "" {
		ctx.JSON(http.StatusLengthRequired, gin.H{"error": "Please URL needs to be filled"})
		return
	}

	if input.Url[0:4] != "http" && input.Url[0:5] != "https" {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Mal-formed URL"})
		return
	}

	urlHashBytes := sha256Of(input.Url + time.Now().String())
	generatedNumber := new(big.Int).SetBytes(urlHashBytes).Uint64()
	code := base58Encoded([]byte(fmt.Sprintf("%d", generatedNumber)))[:6]

	shortenUrlToPersist := models.ShortenUrl{Code: code, Original: input.Url, Access_number: 0}

	// Insert generated url
	if result := ctlr.DB.Create(&shortenUrlToPersist); result.Error != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": result.Error})
		return
	}

	var shortenUrlOutput models.ShortenUrlOutput
	shortenUrlOutput.Code = code
	shortenUrlOutput.ShortenUrl = "https://jr.com/" + code
	ctx.JSON(http.StatusCreated, shortenUrlOutput)
}

func sha256Of(input string) []byte {
	algorithm := sha256.New()
	algorithm.Write([]byte(input))
	return algorithm.Sum(nil)
}

func base58Encoded(bytes []byte) string {
	encoding := base58.BitcoinEncoding
	encoded, err := encoding.Encode(bytes)
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}
	return string(encoded)
}
