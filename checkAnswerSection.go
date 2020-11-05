package main

import (
	"context"
	"net/http"

	"github.com/labstack/echo-contrib/session"
	"github.com/labstack/echo/v4"
)

func checkAnswerSection(c echo.Context) error {
	sess, err := session.Get("session", c)
	if err != nil {
		return c.String(http.StatusInternalServerError, "Error")
	}
	if b, _ := sess.Values["authenticated"]; b != true {
		return c.String(http.StatusUnauthorized, "401")
	}
	userID := c.FormValue("user_id")
	startTime := c.FormValue("start_time")
	endTime := c.FormValue("end_time")
	otherFocusSecond := c.FormValue("other_focus_second")
	answerResultIDs := c.FormValue("answer_result_ids")
	correctAnswerNumber := c.FormValue("correct_answer_number")

	blink := c.FormValue("blink")
	faceMove := c.FormValue("face_move")
	angle := c.FormValue("angle")
	w := c.FormValue("w")
	c1 := c.FormValue("c1")
	c2 := c.FormValue("c2")
	c3 := c.FormValue("c3")
	method1 := c.FormValue("method1")
	concentration := c.FormValue("concentration")
	method2 := c.FormValue("method2")
	faceImagePath := c.FormValue("face_image_path")

	db := sqlConnect()
	defer db.Close()

	var user User
	userIDInt := stringToUint(userID)
	db.First(&user, userIDInt)

	answerResultSection := AnswerResultSection{
		UserID:              stringToUint(userID),
		AnswerResultIDs:     answerResultIDs,
		CorrectAnswerNumber: stringToUint(correctAnswerNumber),
		OtherFocusSecond:    stringToUint(otherFocusSecond),
		FaceImagePath:       faceImagePath,
		StartTime:           stringToTime(startTime),
		EndTime:             stringToTime(endTime),
	}

	if c.FormValue("test") == "true" {
		tx := db.Begin()
		err := tx.Create(&answerResultSection)
		tx.Rollback()
		if err != nil {
			return c.NoContent(http.StatusInternalServerError)
		}
		return c.JSON(http.StatusOK, "testok")
	}
	err = db.Create(&answerResultSection).Error
	if err != nil {
		return err
	}
	mc, ctx := mongoConnect()
	defer mc.Disconnect(ctx)

	if method1 == "true" {
		col1 := mc.Database("fe-concentration").Collection("concentration")
		col1.InsertOne(context.Background(), ConcentrationData{
			UserID:                user.ID,
			AnswerResultSectionID: answerResultSection.ID,
			FaceImagePath:         faceImagePath,
			Blink:                 blink,
			FaceMove:              faceMove,
			Angle:                 angle,
			W:                     w,
			C1:                    c1,
			C2:                    c2,
			C3:                    c3,
		})
	}
	if method2 == "true" {
		col2 := mc.Database("fe-concentration").Collection("son-concentration")
		col2.InsertOne(context.Background(), SonConcentrationData{
			UserID:                user.ID,
			AnswerResultSectionID: answerResultSection.ID,
			FaceImagePath:         faceImagePath,
			Concentration:         concentration,
		})
	}
	answerResultSectionIDSend := AnswerResultSectionIDSend{}
	answerResultSectionIDSend.AnswerResultSectionID = answerResultSection.ID

	return c.JSON(http.StatusOK, answerResultSectionIDSend)
}
