package main

import (
	"math/rand"
	"net/http"
	"time"

	"github.com/labstack/echo"
)

func getQuestionIds(c echo.Context) error {
	db := sqlConnect()
	defer db.Close()
	questions := Question{}
	db.Last(&questions)
	var count int
	db.Table("questions").Count(&count)
	firstId := int(questions.ID) - int(count) + 1

	s_ids_str := c.FormValue("solved_ids")
	new_question_list := []int{}
	solve_list := strToIntList(s_ids_str)

	for {
		rand.Seed(time.Now().UnixNano())
		random := rand.Intn(count)
		if true == searchIDs(solve_list, random) {
			continue
		}
		if true == searchIDs(new_question_list, random) {
			continue
		}
		new_question_list = append(new_question_list, firstId+random)
		if len(new_question_list) == 10 {
			break
		}
	}
	gqi := GetQuestionIDs{
		QuestionIDs: new_question_list,
		SolvedIDs:   solve_list,
	}
	return c.JSON(http.StatusOK, gqi)
}
