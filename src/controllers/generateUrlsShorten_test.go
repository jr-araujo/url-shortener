package controllers

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/gin-gonic/gin"
	"urlshortener.com/devgym/jr/models"
	database "urlshortener.com/devgym/jr/repository"
)

func SetUpRouter() *gin.Engine {
	router := gin.Default()
	return router
}

func TestGenerateUrlShorten_ShouldFail_WhenParsingInput(t *testing.T) {
	mockResponse := `{"error":"EOF"}`
	r := SetUpRouter()
	DB := database.Init()
	ctlr := New(DB)
	r.POST("/shorteners", ctlr.GenerateUrlShorten)

	// inputUrl := models.InputUrl{}
	// jsonValue, _ := json.Marshal(inputUrl)
	req, _ := http.NewRequest("POST", "/shorteners", bytes.NewBuffer(nil))

	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	responseData, _ := io.ReadAll(w.Body)
	assert.Equal(t, mockResponse, string(responseData))
	assert.Equal(t, http.StatusBadRequest, w.Code)
}

func TestGenerateUrlShorten_ShouldFail_WhenPassingEmptyUrl(t *testing.T) {
	mockResponse := `{"error":"Please URL needs to be filled"}`
	r := SetUpRouter()
	DB := database.Init()
	ctlr := New(DB)
	r.POST("/shorteners", ctlr.GenerateUrlShorten)

	inputUrl := models.InputUrl{}
	jsonValue, _ := json.Marshal(inputUrl)
	req, _ := http.NewRequest("POST", "/shorteners", bytes.NewBuffer(jsonValue))

	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	responseData, _ := io.ReadAll(w.Body)
	assert.Equal(t, mockResponse, string(responseData))
	assert.Equal(t, http.StatusLengthRequired, w.Code)
}

func TestGenerateUrlShorten_ShouldFail_WhenInputUrlDoestHaveSchema(t *testing.T) {
	mockResponse := `{"error":"Mal-formed URL"}`
	r := SetUpRouter()
	DB := database.Init()
	ctlr := New(DB)
	r.POST("/shorteners", ctlr.GenerateUrlShorten)

	inputUrl := models.InputUrl{
		Url: "www.test.com",
	}
	jsonValue, _ := json.Marshal(inputUrl)
	req, _ := http.NewRequest("POST", "/shorteners", bytes.NewBuffer(jsonValue))

	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	responseData, _ := io.ReadAll(w.Body)
	assert.Equal(t, mockResponse, string(responseData))
	assert.Equal(t, http.StatusBadRequest, w.Code)
}

func TestGenerateUrlShorten_ShouldBeSuccessful_WhenInputGoodFormedUrl(t *testing.T) {
	r := SetUpRouter()
	DB := database.Init()
	ctlr := New(DB)
	r.POST("/shorteners", ctlr.GenerateUrlShorten)

	inputUrl := models.InputUrl{
		Url: "https://www.test.com",
	}
	jsonValue, _ := json.Marshal(inputUrl)
	req, _ := http.NewRequest("POST", "/shorteners", bytes.NewBuffer(jsonValue))

	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusCreated, w.Code)
}
