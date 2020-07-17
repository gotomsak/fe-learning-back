package main

import (
	"net/http/httptest"
	"testing"
)

func TestGetQuestion(t *testing.T) {
	e := router()
	req := httptest.NewRequest("GET", "/question?id=6000", nil)

	rec := httptest.NewRecorder()

	e.ServeHTTP(rec, req)

}
