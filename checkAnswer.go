package main

import (
	"net/http"

	"github.com/labstack/echo-contrib/session"
	"github.com/labstack/echo/v4"
)

func checkAnswer(c echo.Context) error {
	sess, err := session.Get("session", c)
	if err != nil {
		return c.String(http.StatusInternalServerError, "Error")
	}
	if b, _ := sess.Values["authenticated"]; b != true {
		return c.String(http.StatusUnauthorized, "401")
	}

	answerResult := AnswerResult{}
	question := Question{}

	startTime := c.FormValue("start_time")
	endTime := c.FormValue("end_time")
	otherFocusSecond := c.FormValue("other_focus_second")
	uotherFocusSecond := stringToUint(otherFocusSecond)
	questionID := c.FormValue("question_id")
	uquestionID := stringToUint(questionID)
	userID := c.FormValue("user_id")
	uuserID := stringToUint(userID)
	answerResult.UserAnswer = c.FormValue("user_answer")
	answerResult.QuestionID = uquestionID
	answerResult.MemoLog = c.FormValue("memo_log")
	answerResult.StartTime = stringToTime(startTime)
	answerResult.EndTime = stringToTime(endTime)
	answerResult.OtherFocusSecond = uotherFocusSecond
	answerResult.UserID = uuserID

	db := sqlConnect()
	defer db.Close()

	db.First(&question, questionID)
	result := "incorrect"
	var answer string
	if question.AimgPath != "" {
		answer = question.AimgPath
	} else {
		answer = question.Ans
	}

	if question.AimgPath == answerResult.UserAnswer || question.Ans == answerResult.UserAnswer {
		result = "correct"
	}
	answerResult.AnswerResult = result
	answerResultSend := AnswerResultSend{
		Result: result,
		Answer: answer,
	}

	if c.FormValue("test") == "true" {
		tx := db.Begin()
		err = tx.Create(&answerResult).Error
		tx.Rollback()
		if err != nil {
			return c.NoContent(http.StatusInternalServerError)
		}
		return c.JSON(http.StatusOK, result)
	}

	err = db.Create(&answerResult).Error

	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, answerResultSend)
}
