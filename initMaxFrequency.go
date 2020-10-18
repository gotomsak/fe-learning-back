package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"strconv"

	"github.com/labstack/echo-contrib/session"
	"github.com/labstack/echo/v4"
)

func initMaxFrequency(c echo.Context) error {
	sess, err := session.Get("session", c)
	if err != nil {
		return c.String(http.StatusInternalServerError, "Error")
	}
	if b, _ := sess.Values["authenticated"]; b != true {
		return c.String(http.StatusUnauthorized, "401")
	}

	userID := c.FormValue("user_id")
	maxBlinkNumber := c.FormValue("max_blink_number")
	maxFaceMoveNumber := c.FormValue("max_face_move_number")
	maxFrequencyVideo, err := c.FormFile("max_frequency_video")
	maxBlinkNumberFloat, _ := strconv.ParseFloat(maxBlinkNumber, 64)
	maxFaceMoveNumberFloat, _ := strconv.ParseFloat(maxFaceMoveNumber, 64)
	var maxBlinkFrequency float64 = (maxBlinkNumberFloat / 60) * 5
	var maxFaceMoveFrequency float64 = (maxFaceMoveNumberFloat / 60) * 5

	if err != nil {
		fmt.Println(err)
		return c.NoContent(http.StatusInternalServerError)
	}
	src, err := maxFrequencyVideo.Open()
	if err != nil {
		fmt.Println(err)
		return c.NoContent(http.StatusInternalServerError)
	}
	defer src.Close()
	faceVideoDir := "./data/" + userID
	faceVideoFile := "maxFrequency" + ".mp4"
	if err := os.MkdirAll(faceVideoDir, 0777); err != nil {
		fmt.Println(err)
	}
	dstFile, err := os.Create(faceVideoDir + "/" + faceVideoFile)
	if err != nil {
		fmt.Println(err)
		return c.NoContent(http.StatusInternalServerError)
	}
	defer dstFile.Close()
	if _, err = io.Copy(dstFile, src); err != nil {
		fmt.Println(err)
		return err
	}

	db := sqlConnect()
	defer db.Close()

	var frequency Frequency
	err = db.Where("user_id = ?", userID).First(&frequency).Error
	if err != nil {
		frequency := Frequency{
			UserID:               stringToUint(userID),
			MaxFrequencyVideo:    faceVideoDir + "/" + faceVideoFile,
			MaxFaceMoveNumber:    maxFaceMoveNumberFloat,
			MaxFaceMoveFrequency: maxFaceMoveFrequency,
			MaxBlinkNumber:       maxBlinkNumberFloat,
			MaxBlinkFrequency:    maxBlinkFrequency,
		}
		err = db.Create(&frequency).Error
		if err != nil {
			fmt.Println(err)
			return err
		}
	}
	err = db.Model(&frequency).Updates(Frequency{MaxFrequencyVideo: faceVideoDir + "/" + faceVideoFile,
		MaxFaceMoveNumber:    maxFaceMoveNumberFloat,
		MaxFaceMoveFrequency: maxFaceMoveFrequency,
		MaxBlinkNumber:       maxBlinkNumberFloat,
		MaxBlinkFrequency:    maxBlinkFrequency,
	}).Error
	if err != nil {
		fmt.Println(err)
		return err
	}
	fmt.Println(frequency)

	return c.JSON(http.StatusOK, "ok")
}
