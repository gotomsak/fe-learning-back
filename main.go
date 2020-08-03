package main

import (
	"net/http"

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
	// e.Use(middleware.CORS())
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins:     []string{"http://localhost:3000"},
		AllowCredentials: true,
		AllowMethods:     []string{http.MethodGet, http.MethodPut, http.MethodPost, http.MethodDelete},
	}))
	e.Use(session.Middleware(sessions.NewCookieStore([]byte("secret"))))

	e.POST("/signin", signin)
	e.POST("/question_ids", getQuestionIds)
	e.GET("/question", getQuestion)
	e.POST("/check_answer", checkAnswer)
	e.POST("/check_answer_section", checkAnswerSection)
	e.POST("/save_questionnaire", saveQuestionnaire)
	e.POST("/signup", signup)
	e.GET("/signout", signout)
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
