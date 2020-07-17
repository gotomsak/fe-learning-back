package main

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

func router() *echo.Echo {
	e := echo.New()
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.Use(middleware.CORS())
	e.POST("/question_ids", getQuestionIds)
	e.GET("/question", getQuestion)
	return e
}

func envLoad() {
	err := godotenv.Load()
	if err != nil {
		panic(err.Error())
	}
}

func main() {

	envLoad()
	e := router()

	e.Logger.Fatal(e.Start(":1323"))
}
