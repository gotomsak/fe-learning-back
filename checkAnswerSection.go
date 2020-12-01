package main

import (
	"context"
	"fmt"
	"net/http"

	"github.com/labstack/echo-contrib/session"
	"github.com/labstack/echo/v4"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func checkAnswerSection(c echo.Context) error {
	sess, err := session.Get("session", c)
	if err != nil {
		return c.String(http.StatusInternalServerError, "Please sign in")
	}
	if b, _ := sess.Values["authenticated"]; b != true {
		return c.String(http.StatusUnauthorized, "401")
	}

	cas := new(CheckAnswerSection)
	if err = c.Bind(cas); err != nil {
		return c.String(http.StatusInternalServerError, "The format is different")
	}
	fmt.Println(cas)

	db := sqlConnect()
	defer db.Close()
	mc, ctx := mongoConnect()
	defer mc.Disconnect(ctx)

	results := mc.Database("fe-concentration").Collection("answer_result_sectoin_ids")
	res, err := results.InsertOne(context.Background(), Results{ResultIDs: cas.AnswerResultIDs})
	var resID string
	if oid, ok := res.InsertedID.(primitive.ObjectID); ok {
		resID = oid.Hex()

	} else {

		return c.JSON(http.StatusInternalServerError, "Not objectid.ObjectID, do what you want")
	}
	answerResultSection := AnswerResultSection{
		UserID:              cas.UserID,
		AnswerResultIDs:     resID,
		CorrectAnswerNumber: cas.CorrectAnswerNumber,
		OtherFocusSecond:    cas.OtherFocusSecond,
		FaceImagePath:       cas.FaceImagePath,
		StartTime:           cas.StartTime,
		EndTime:             cas.EndTime,
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

	if cas.Method1 == true {
		col1 := mc.Database("fe-concentration").Collection("concentration")
		col1.InsertOne(context.Background(), ConcentrationData{
			UserID:                cas.UserID,
			AnswerResultSectionID: answerResultSection.ID,
			FaceImagePath:         cas.FaceImagePath,
			Blink:                 cas.Blink,
			FaceMove:              cas.FaceMove,
			Angle:                 cas.Angle,
			W:                     cas.W,
			C1:                    cas.C1,
			C2:                    cas.C2,
			C3:                    cas.C3,
		})
	}
	if cas.Method2 == true {
		col2 := mc.Database("fe-concentration").Collection("son-concentration")
		col2.InsertOne(context.Background(), SonConcentrationData{
			UserID:                cas.UserID,
			AnswerResultSectionID: answerResultSection.ID,
			FaceImagePath:         cas.FaceImagePath,
			Concentration:         cas.Concentration,
		})
	}
	answerResultSectionIDSend := AnswerResultSectionIDSend{}
	answerResultSectionIDSend.AnswerResultSectionID = answerResultSection.ID

	return c.JSON(http.StatusOK, answerResultSectionIDSend)
}
