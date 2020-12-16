package main

import (
	"context"
	"fmt"
	"net/http"

	"github.com/labstack/echo-contrib/session"
	"github.com/labstack/echo/v4"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func checkAnswer(c echo.Context) error {
	sess, err := session.Get("session", c)
	if err != nil {
		return c.String(http.StatusInternalServerError, "Error")
	}
	if b, _ := sess.Values["authenticated"]; b != true {
		return c.String(http.StatusUnauthorized, "401")
	}

	question := Question{}
	ca := new(CheckAnswer)
	if err = c.Bind(ca); err != nil {
		return c.String(http.StatusInternalServerError, "The format is different")
	}
	fmt.Println("kusoga")
	mc, ctx := mongoConnect()
	defer mc.Disconnect(ctx)

	results := mc.Database("fe-concentration").Collection("concentration")
	concData := ConcentrationData{
		ConcentrationData: ca.ConcentrationData,
	}
	res, err := results.InsertOne(context.Background(), concData)
	var resID string
	if oid, ok := res.InsertedID.(primitive.ObjectID); ok {
		resID = oid.Hex()
	} else {
		return c.JSON(http.StatusInternalServerError, "Not objectid.ObjectID, do what you want")
	}

	db := sqlConnect()
	defer db.Close()

	db.First(&question, ca.QuestionID)
	result := "incorrect"
	var answer string
	if question.AimgPath != "" {
		answer = question.AimgPath
	} else {
		answer = question.Ans
	}

	if question.AimgPath == ca.UserAnswer || question.Ans == ca.UserAnswer {
		result = "correct"
	}
	answerResult := AnswerResult{
		UserID:            ca.UserID,
		UserAnswer:        ca.UserAnswer,
		AnswerResult:      result,
		ConcentrationData: resID,
		MemoLog:           ca.MemoLog,
		OtherFocusSecond:  ca.OtherFocusSecond,
		QuestionID:        ca.QuestionID,
		StartTime:         stringToTime(ca.StartTime),
		EndTime:           stringToTime(ca.EndTime),
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

	answerResultSend := AnswerResultSend{
		AnswerResultID: answerResult.ID,
		Result:         result,
		Answer:         answer,
	}

	return c.JSON(http.StatusOK, answerResultSend)
}
