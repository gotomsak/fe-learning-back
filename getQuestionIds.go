package main

import (
	"math/rand"
	"net/http"
	"time"

	"github.com/labstack/echo-contrib/session"
	"github.com/labstack/echo/v4"
)

func getQuestionIds(c echo.Context) error {
	sess, err := session.Get("session", c)
	if err != nil {
		return c.String(http.StatusInternalServerError, "Error")
	}
	if b, _ := sess.Values["authenticated"]; b != true {
		return c.String(http.StatusUnauthorized, "401")
	}

	db := sqlConnect()
	defer db.Close()
	questions := Question{}
	db.Last(&questions)
	var count int
	db.Table("questions").Count(&count)
	firstID := int(questions.ID) - int(count) + 1
	firstIDU := uint(firstID)
	sIdsStr := c.FormValue("solved_ids")
	qIdsStr := c.FormValue("question_ids")
	newQuestionList := []uint{}
	solveList := strToUIntList(sIdsStr)
	questionList := strToUIntList(qIdsStr)
	solveList = append(solveList, questionList...)

	for {
		rand.Seed(time.Now().UnixNano())
		random := rand.Intn(count)
		urandom := uint(random)
		if true == searchIDs(solveList, urandom) {
			continue
		}
		if true == searchIDs(newQuestionList, urandom) {
			continue
		}
		newQuestionList = append(newQuestionList, firstIDU+urandom)
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
