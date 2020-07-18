package main

import (
	"log"
	"net/http/httptest"
	"net/url"
	"strings"
	"testing"
)

func TestGetQuestion(t *testing.T) {
	envLoad()
	db := sqlConnect()
	tx := db.Begin()

	e := router()
	req := httptest.NewRequest("GET", "/question?id=6000", nil)

	rec := httptest.NewRecorder()

	e.ServeHTTP(rec, req)
	log.Print(rec.Body)
	tx.Rollback()
}

func TestGetQusetionIDs(t *testing.T) {
	envLoad()
	db := sqlConnect()
	tx := db.Begin()
	e := router()
	values := url.Values{}
	values.Set("question", "[]")
	values.Set("question_ids", "[]")
	body := strings.NewReader(values.Encode())
	req := httptest.NewRequest("POST", "/question_ids", body)
	rec := httptest.NewRecorder()

	e.ServeHTTP(rec, req)
	log.Print(rec.Body)
	tx.Rollback()
}
