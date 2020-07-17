package main

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

// func test(c echo.Context) error {
// 	db := sqlConnect()
// 	defer db.Close()
// 	questionSend := QuestionSend{}
// 	idTest := Question{}
// 	// db.First(&idTest)
// 	// db.First(&idTest, 5312)
// 	db.First(&idTest, 5502)
// 	qimg := regexp.MustCompile(",").Split(idTest.QimgPath, -1)
// 	ans := []string{idTest.Ans, idTest.Mistake1, idTest.Mistake2, idTest.Mistake3}
// 	aimg := []string{idTest.AimgPath, idTest.MimgPath1, idTest.MimgPath2, idTest.MimgPath3}
// 	shuffle(ans)
// 	shuffle(aimg)
// 	questionSend.QuestionID = idTest.ID
// 	questionSend.AnsList = ans
// 	questionSend.AimgList = aimg
// 	questionSend.QimgPath = qimg
// 	questionSend.QuestionNum = idTest.QuestionNum
// 	questionSend.Question = idTest.Question
// 	questionSend.Season = idTest.Season
// 	questionSend.Genre = idTest.Genre
// 	return c.JSON(http.StatusOK, questionSend)
// }

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
