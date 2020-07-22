package main

import (
	"math/rand"
	"net/http"
	"time"

	"github.com/labstack/echo/v4"
)

func getQuestionIds(c echo.Context) error {
	db := sqlConnect()
	defer db.Close()
	questions := Question{}
	db.Last(&questions)
	var count int
	db.Table("questions").Count(&count)
	firstID := int(questions.ID) - int(count) + 1

	sIdsStr := c.FormValue("solved_ids")
	qIdsStr := c.FormValue("question_ids")
	newQuestionList := []int{}
	solveList := strToIntList(sIdsStr)
	questionList := strToIntList(qIdsStr)
	solveList = append(solveList, questionList...)

	for {
		rand.Seed(time.Now().UnixNano())
		random := rand.Intn(count)
		if true == searchIDs(solveList, random) {
			continue
		}
		if true == searchIDs(newQuestionList, random) {
			continue
		}
		newQuestionList = append(newQuestionList, firstID+random)
		if len(newQuestionList) == 10 {
			break
		}
	}
	gqi := GetQuestionIDs{
		QuestionIDs: newQuestionList,
		SolvedIDs:   solveList,
	}
	return c.JSON(http.StatusOK, gqi)
}
