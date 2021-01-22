package main

import (
	"fmt"
	"net/http"

	"github.com/labstack/echo-contrib/session"
	"github.com/labstack/echo/v4"
)

func getQuestionGym(c echo.Context) error {
	sess, err := session.Get("session", c)
	if err != nil {
		return c.String(http.StatusInternalServerError, "Error")
	}
	if b, _ := sess.Values["authenticated"]; b != true {
		return c.String(http.StatusUnauthorized, "401")
	}
	db := sqlConnect()
	defer db.Close()
	qg := new(GetQuestionGymBind)
	if err = c.Bind(qg); err != nil {
		return c.String(http.StatusInternalServerError, "The format is different")
	}
	fmt.Println(qg.NowLevel)
	// level := qg.NowLevel / 2
	// levelI := int(level)
	questionsSub := QuestionsSub{}

	db.Where("level = ?", qg.NowLevel).Order("rand()").Find(&questionsSub)
	// Res := GetQuestionGymRes{}

	return c.JSON(http.StatusOK, questionsSub)
}
