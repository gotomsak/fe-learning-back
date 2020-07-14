package main

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

// type Question struct {
// 	gorm.Model
// 	Question    string `json:"question"`
// 	QimgPath    string `json:"qimg_path"`
// 	Mistake1    string `json:"mistake1"`
// 	Mistake2    string `json:"mistake2"`
// 	Mistake3    string `json:"mistake3"`
// 	Ans         string `json:"ans"`
// 	MimgPath1   string `json:"mimg_path1"`
// 	MimgPath2   string `json:"mimg_path2"`
// 	MimgPath3   string `json:"mimg_path3"`
// 	AimgPath    string `json:"aimg_path"`
// 	Season      string `json:"season"`
// 	QuestionNum string `json:"question_num"`
// 	Genre       string `json:"genre"`
// }

// type QuestionSend struct {
// 	QuestionID  uint     `json:"question_id"`
// 	Question    string   `json:"question"`
// 	QimgPath    []string `json:"qimg_path"`
// 	AnsList     []string `json:"ans_list"`
// 	AimgList    []string `json:"aimg_list"`
// 	Season      string   `json:"season"`
// 	QuestionNum string   `json:"question_num"`
// 	Genre       string   `json:"genre"`
// }

// func sqlConnect() (database *gorm.DB) {
// 	DBMS := os.Getenv("DBMS")
// 	USER := os.Getenv("USERR")
// 	PASS := os.Getenv("PASS")
// 	PROTOCOL := os.Getenv("PROTOCOL")
// 	DBNAME := os.Getenv("DBNAME")
// 	CONNECT := USER + ":" + PASS + "@" + PROTOCOL + "/" + DBNAME + "?charset=utf8&parseTime=true&loc=Asia%2FTokyo"
// 	db, err := gorm.Open(DBMS, CONNECT)
// 	if err != nil {
// 		panic(err.Error())
// 	}
// 	return db
// }

// func shuffle(a []string) {
// 	rand.Seed(time.Now().UnixNano())
// 	for i := range a {
// 		j := rand.Intn(i + 1)
// 		a[i], a[j] = a[j], a[i]
// 	}
// }

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

func main() {

	eerr := godotenv.Load()
	if eerr != nil {
		panic(eerr.Error())
	}

	e := echo.New()
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.Use(middleware.CORS())
	e.POST("/question_ids", getQuestionIds)
	e.Logger.Fatal(e.Start(":1323"))
}
