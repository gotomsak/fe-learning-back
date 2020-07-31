package main

import (
	"net/http"
	"strconv"

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

	db := sqlConnect()
	defer db.Close()
	concentration, _ := strconv.Atoi(c.FormValue("concentration"))
	whiledoing, _ := strconv.ParseBool(c.FormValue("while_doing"))
	cheating, _ := strconv.ParseBool(c.FormValue("cheating"))
	nonsense, _ := strconv.ParseBool(c.FormValue("nonsense"))
	questionnaire := Questionnaire{
		AnswerResultSectionID: stringToUint(c.FormValue("answer_result_section_id")),
		UserID:                stringToUint(c.FormValue("user_id")),
		Concentration:         concentration,
		WhileDoing:            whiledoing,
		Cheating:              cheating,
		Nonsense:              nonsense,
	}

	if c.FormValue("test") == "true" {
		tx := db.Begin()
		err = tx.Create(&questionnaire).Error
		if err != nil {
			return c.String(http.StatusInternalServerError, "500")
		}
		tx.Rollback()
		return c.JSON(http.StatusOK, "testOK")
	}
	err = db.Create(&questionnaire).Error

	if err != nil {
		return c.String(http.StatusInternalServerError, "500")
	}

	return c.JSON(http.StatusOK, "200")
}
