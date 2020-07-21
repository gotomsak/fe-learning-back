package main

import (
	"log"
	"net/http"
	"net/http/httptest"
	"net/url"
	"strings"
	"testing"

	"github.com/labstack/echo"
)

func TestGetQuestion(t *testing.T) {
	envLoad()

	e := router()
	req := httptest.NewRequest("GET", "/question?id=6000", nil)

	rec := httptest.NewRecorder()

	e.ServeHTTP(rec, req)
	log.Print(rec.Body)
}

func TestGetQusetionIDs(t *testing.T) {
	envLoad()
	e := router()
	values := url.Values{}
	values.Set("solved_ids", "[]")
	values.Set("question_ids", "[]")
	body := strings.NewReader(values.Encode())
	log.Print(body)
	req := httptest.NewRequest("POST", "/question_ids", body)
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationForm)
	rec := httptest.NewRecorder()
	e.ServeHTTP(rec, req)
	log.Print(rec.Body)
}

func TestSignup(t *testing.T) {
	envLoad()
	e := router()

	values := url.Values{}

	values.Set("username", "hoge")
	values.Set("password", "foobar")
	values.Set("email", "example.com")
	// values.Set("test", "true")

	body := strings.NewReader(values.Encode())
	req := httptest.NewRequest(http.MethodPost, "/signup", body)
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationForm)
	rec := httptest.NewRecorder()
	e.ServeHTTP(rec, req)
}

func TestSignin(t *testing.T) {
	envLoad()
	e := router()
	values := url.Values{}

	values.Set("password", "foobar")
	values.Set("email", "example.com")
	body := strings.NewReader(values.Encode())

	req := httptest.NewRequest(http.MethodPost, "/signin", body)
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationForm)
	rec := httptest.NewRecorder()
	e.ServeHTTP(rec, req)

}
