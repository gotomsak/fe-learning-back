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
	// key := []byte("super-secret-key")
	// store := sessions.NewCookieStore(key)
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	// e.Use(middleware.CORS())
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins:     []string{"http://localhost:3000", "https://localhost:3000", "https://192.168.1.10:3000", "https://fe-learning.gotomsak.work"},
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
	e.POST("/init_max_frequency", initMaxFrequency)
	e.POST("/init_min_frequency", initMinFrequency)
	e.GET("/signout", signout)
	e.GET("/check_session", checkSession)
	e.POST("/question_gym", getQuestionGym)
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
	db := sqlConnect()
	defer db.Close()
	e.Logger.Fatal(e.Start(":1323"))
	// e.Logger.Fatal(e.StartTLS(":1323", "./fullchain.pem", "./privkey.pem"))
}
