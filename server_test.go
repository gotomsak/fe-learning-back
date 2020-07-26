package main

import (
	"log"
	"net/http"
	"net/http/httptest"
	"net/url"
	"strings"
	"testing"
	"time"

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

func TestCheckAnswer(t *testing.T) {
	envLoad()
	e := router()

	nowTime := time.Now()
	nowTimeString := nowTime.Format(layout)
	log.Print(nowTimeString)
	otherFocusSecond := "26"
	questionID := "6000"
	userID := "66"
	memoLog := "wakaranai"
	userAnswer := "ディスプレイに映像，文字などの情報を表示する電子看板"

	values := url.Values{}
	values.Set("start_time", nowTimeString)
	values.Set("end_time", nowTimeString)
	values.Set("other_focus_second", otherFocusSecond)
	values.Set("question_id", questionID)
	values.Set("user_id", userID)
	values.Set("memo_log", memoLog)
	values.Set("user_answer", userAnswer)
	values.Set("test", "true")
	body := strings.NewReader(values.Encode())
	req := httptest.NewRequest("POST", "/check_answer", body)
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationForm)
	req.Header.Add("Cookie", cookie.Name+"="+cookie.Value)
	rec := httptest.NewRecorder()
	e.ServeHTTP(rec, req)
	log.Print(rec.Body)
}
