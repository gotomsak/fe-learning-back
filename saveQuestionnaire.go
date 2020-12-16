package main

import (
	"net/http"

	"github.com/labstack/echo-contrib/session"
	"github.com/labstack/echo/v4"
)

func saveQuestionnaire(c echo.Context) error {
	sess, err := session.Get("session", c)
	if err != nil {
		return c.String(http.StatusInternalServerError, "Error")
	}
	if b, _ := sess.Values["authenticated"]; b != true {
		return c.String(http.StatusUnauthorized, "401")
	}
	q := new(Questionnaire)

	if err = c.Bind(q); err != nil {
		return c.String(http.StatusInternalServerError, "The format is different")
	}

	db := sqlConnect()
	defer db.Close()

	if c.FormValue("test") == "true" {
		tx := db.Begin()
		err = tx.Create(&q).Error
		if err != nil {
			return c.String(http.StatusInternalServerError, "500")
		}
		tx.Rollback()
		return c.JSON(http.StatusOK, "testOK")
	}
	err = db.Create(&q).Error

	if err != nil {
		return c.String(http.StatusInternalServerError, "500")
	}

	return c.JSON(http.StatusOK, "200")
}
