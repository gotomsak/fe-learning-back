package main

import (
	"fmt"
	"io"
	"net/http"
	"os"

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

	videoFile, err := c.FormFile("face_video")
	if err != nil {
		return c.NoContent(http.StatusInternalServerError)
	}
	src, err := videoFile.Open()
	if err != nil {
		return c.NoContent(http.StatusInternalServerError)
	}
	defer src.Close()
	faceVideoDir := "./data/" + userID
	faceVideoFile := endTime + ".mp4"
	if err := os.MkdirAll(faceVideoDir, 0777); err != nil {
		fmt.Println(err)
	}
	dstFile, err := os.Create(faceVideoDir + "/" + faceVideoFile)
	if err != nil {
		return c.NoContent(http.StatusInternalServerError)
	}
	defer dstFile.Close()
	if _, err = io.Copy(dstFile, src); err != nil {
		return err
	}

	blink := c.FormValue("blink")
	faceMove := c.FormValue("face_move")
	w := c.FormValue("w")
	c1 := c.FormValue("c1")
	c2 := c.FormValue("c2")
	c3 := c.FormValue("c3")
	fmt.Println(c1)

	fmt.Println(c2)
	fmt.Println(c3)
	fmt.Println(w)

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
		FaceVideoPath:       faceVideoDir + "/" + faceVideoFile,
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
	concentrationData := ConcentrationData{
		UserID:                user.ID,
		AnswerResultSectionID: answerResultSection.ID,
		Blink:                 stringToUint(blink),
		FaceMove:              stringToUint(faceMove),
		W:                     w,
		C1:                    c1,
		C2:                    c2,
		C3:                    c3,
	}
	err = db.Create(&concentrationData).Error
	if err != nil {
		fmt.Println("noneConcentration")
	}
	return c.JSON(http.StatusOK, "ok")
}
