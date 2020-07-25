package main

import (
	"net/http"
	"regexp"
	"strconv"

	"github.com/labstack/echo-contrib/session"
	"github.com/labstack/echo/v4"
)

func getQuestion(c echo.Context) error {
	sess, err := session.Get("session", c)
	if err != nil {
		return c.String(http.StatusInternalServerError, "Error")
	}
	if b, _ := sess.Values["authenticated"]; b != true {
		return c.String(http.StatusUnauthorized, "401")
	}
	db := sqlConnect()
	defer db.Close()
	question := Question{}
	questionSend := QuestionSend{}
	questionID := c.QueryParam("id")
	questionIDI, _ := strconv.Atoi(questionID)
	db.First(&question, questionIDI)
	qimg := regexp.MustCompile(",").Split(question.QimgPath, -1)
	ans := []string{question.Ans, question.Mistake1, question.Mistake2, question.Mistake3}
	aimg := []string{question.AimgPath, question.MimgPath1, question.MimgPath2, question.MimgPath3}
	shuffle(ans)
	shuffle(aimg)
	questionSend.QuestionID = question.ID
	questionSend.AnsList = ans
	questionSend.AimgList = aimg
	questionSend.QimgPath = qimg
	questionSend.QuestionNum = question.QuestionNum
	questionSend.Question = question.Question
	questionSend.Season = question.Season
	questionSend.Genre = question.Genre
	return c.JSON(http.StatusOK, questionSend)
}
