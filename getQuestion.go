package main

import (
	"fmt"
	"net/http"
	"regexp"
	"strconv"

	"github.com/labstack/echo"
)

func getQuestion(c echo.Context) error {
	fmt.Println("nyan")
	db := sqlConnect()
	defer db.Close()
	question := Question{}
	questionSend := QuestionSend{}
	questionId := c.QueryParam("id")
	questionIdI, _ := strconv.Atoi(questionId)
	db.First(&question, questionIdI)
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
