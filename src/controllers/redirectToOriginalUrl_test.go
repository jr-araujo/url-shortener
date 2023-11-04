package controllers

import (
	"io"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"

	database "urlshortener.com/devgym/jr/repository"
)

func TestRedirectToOriginal_ShouldFail_WhenNoPassingCode(t *testing.T) {
	mockResponse := `{"error":"Code is required."}`
	r := SetUpRouter()
	DB := database.Init()
	ctlr := New(DB)
	r.GET("/shorteners/:code", ctlr.RedirectOriginalUrl)

	req, _ := http.NewRequest("GET", "/shorteners/:code", nil) //bytes.NewBuffer(jsonValue)
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	responseData, _ := io.ReadAll(w.Body)
	assert.Equal(t, mockResponse, string(responseData))
	assert.Equal(t, http.StatusLengthRequired, w.Code)
}

func TestRedirectToOriginal_ShouldFail_WhenNoExistingCode(t *testing.T) {
	mockResponse := `{"error":"Code not found."}`
	r := SetUpRouter()
	DB := database.Init()
	ctlr := New(DB)
	r.GET("/shorteners/:code", ctlr.RedirectOriginalUrl)

	// inputUrl := models.InputUrl{
	// 	Url: "https://www.test.com",
	// }
	// jsonValue, _ := json.Marshal(inputUrl)
	req, _ := http.NewRequest("GET", "/shorteners/123", nil) //bytes.NewBuffer(jsonValue)
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	responseData, _ := io.ReadAll(w.Body)
	assert.Equal(t, mockResponse, string(responseData))
	assert.Equal(t, http.StatusNotFound, w.Code)
}

func TestRedirectToOriginal_ShouldBeSuccessful_WhenPassingValidCode(t *testing.T) {
	r := SetUpRouter()
	DB := database.Init()
	ctlr := New(DB)
	r.GET("/shorteners/:code", ctlr.RedirectOriginalUrl)

	req, _ := http.NewRequest("GET", "/shorteners/H7528o", nil) //bytes.NewBuffer(jsonValue)
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusPermanentRedirect, w.Code)
}
