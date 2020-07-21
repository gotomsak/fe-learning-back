package main

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/sessions"

	"github.com/joho/godotenv"
	"github.com/labstack/echo-contrib/session"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func router() *echo.Echo {
	e := echo.New()
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.Use(middleware.CORS())
	e.Use(session.Middleware(sessions.NewCookieStore([]byte("secret"))))

	e.POST("/signin", signin)
	e.POST("/question_ids", getQuestionIds)
	e.GET("/question", getQuestion)
	e.POST("/signup", signup)
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
