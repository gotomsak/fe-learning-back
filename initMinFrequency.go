package main

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/labstack/echo-contrib/session"
	"github.com/labstack/echo/v4"
)

func initMinFrequency(c echo.Context) error {
	sess, err := session.Get("session", c)
	if err != nil {
		return c.String(http.StatusInternalServerError, "Error")
	}
	if b, _ := sess.Values["authenticated"]; b != true {
		return c.String(http.StatusUnauthorized, "401")
	}

	userID := c.FormValue("user_id")
	minBlinkNumber := c.FormValue("min_blink_number")
	minFaceMoveNumber := c.FormValue("min_face_move_number")
	minBlinkNumberFloat, _ := strconv.ParseFloat(minBlinkNumber, 64)
	minFaceMoveNumberFloat, _ := strconv.ParseFloat(minFaceMoveNumber, 64)
	var minBlinkFrequency float64 = (minBlinkNumberFloat / 60) * 5
	var minFaceMoveFrequency float64 = (minFaceMoveNumberFloat / 60) * 5

	db := sqlConnect()
	defer db.Close()
	var frequency Frequency
	err = db.Where("user_id = ?", userID).First(&frequency).Error
	if err != nil {
		frequency := Frequency{
			UserID: stringToUint(userID),

			MinFaceMoveNumber:    minFaceMoveNumberFloat,
			MinFaceMoveFrequency: minFaceMoveFrequency,
			MinBlinkNumber:       minBlinkNumberFloat,
			MinBlinkFrequency:    minBlinkFrequency,
		}
		err = db.Create(&frequency).Error
		if err != nil {
			fmt.Println(err)
			return err
		}
	}
	db.Model(&frequency).Updates(Frequency{

		MinFaceMoveNumber:    minFaceMoveNumberFloat,
		MinFaceMoveFrequency: minFaceMoveFrequency,
		MinBlinkNumber:       minBlinkNumberFloat,
		MinBlinkFrequency:    minBlinkFrequency,
	})
	fmt.Println(frequency)

	return c.JSON(http.StatusOK, "ok")

}
