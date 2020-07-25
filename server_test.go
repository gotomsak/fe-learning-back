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

var cookie *http.Cookie

func parseCookies(value string) map[string]*http.Cookie {
	m := map[string]*http.Cookie{}
	for _, c := range (&http.Request{Header: http.Header{"Cookie": {value}}}).Cookies() {
		m[c.Name] = c
	}
	return m
}

func TestSignup(t *testing.T) {
	envLoad()
	e := router()

	values := url.Values{}

	values.Set("username", "hoge")
	values.Set("password", "foobar")
	values.Set("email", "example.com")

	body := strings.NewReader(values.Encode())
	req := httptest.NewRequest(http.MethodPost, "/signup", body)
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationForm)
	rec := httptest.NewRecorder()
	e.ServeHTTP(rec, req)

	log.Print(rec.Body)
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
	cookie = parseCookies(rec.Header().Get("Set-Cookie"))["session"]

	log.Print(rec.HeaderMap)
}

func TestGetQuestion(t *testing.T) {
	envLoad()

	e := router()
	req := httptest.NewRequest("GET", "/question?id=6000", nil)

	rec := httptest.NewRecorder()
	req.Header.Add("Cookie", cookie.Name+"="+cookie.Value)

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
	req.Header.Add("Cookie", cookie.Name+"="+cookie.Value)
	rec := httptest.NewRecorder()
	e.ServeHTTP(rec, req)
	log.Print(rec.Body)
}
