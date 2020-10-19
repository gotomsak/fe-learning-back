package main

import (
	"bytes"
	"io"
	"math/rand"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var layout = "2006-01-02 15:04:05"

func sqlConnect() (database *gorm.DB) {
	DBMS := os.Getenv("DBMS")
	USER := os.Getenv("USERR")
	PASS := os.Getenv("PASS")
	PROTOCOL := os.Getenv("PROTOCOL")
	DBNAME := os.Getenv("DBNAME")
	CONNECT := USER + ":" + PASS + "@" + PROTOCOL + "/" + DBNAME + "?charset=utf8mb4&parseTime=true&loc=Asia%2FTokyo"
	db, err := gorm.Open(DBMS, CONNECT)
	db.Set("gorm:table_options", "ENGINE=InnoDB").AutoMigrate(&User{})
	db.Set("gorm:table_options", "ENGINE=InnoDB").AutoMigrate(&AnswerResult{})
	db.Set("gorm:table_options", "ENGINE=InnoDB").AutoMigrate(&AnswerResultSection{})
	db.Set("gorm:table_options", "ENGINE=InnoDB").AutoMigrate(&Questionnaire{})
	db.Set("gorm:table_options", "ENGINE=InnoDB").AutoMigrate(&Frequency{})
	db.Set("gorm:table_options", "ENGINE=InnoDB").AutoMigrate(&ConcentrationData{})
	if err != nil {
		panic(err.Error())
	}
	return db
}

// intsにsearchがあったらそれを削除してリストを返す
func remove(ints []int, search int) []int {
	result := []int{}
	for _, v := range ints {
		if v != search {
			result = append(result, v)
		}
	}
	return result
}

// intsの中にsearchがあったらtrueを返す
func searchIDs(ints []uint, search uint) bool {
	for _, v := range ints {
		if v == search {
			return true
		}
	}
	return false
}

// string型で受け取った数値のリストをInt型のリストにして返す
func strToUIntList(str string) []uint {
	intList := []uint{}
	str = strings.Trim(str, "[]")
	strList := strings.Split(str, ",")
	if str != "" {
		for i := 0; i < len(strList); i++ {
			n, _ := strconv.ParseUint(strList[i], 10, 32)
			un := uint(n)
			intList = append(intList, un)
		}
	}
	return intList
}

// string型のリストをバラバラの順番にして返す
func shuffle(a []string) {
	rand.Seed(time.Now().UnixNano())
	for i := range a {
		j := rand.Intn(i + 1)
		a[i], a[j] = a[j], a[i]
	}
}

func stringToTime(str string) time.Time {
	jst, _ := time.LoadLocation("Asia/Tokyo")
	t, _ := time.ParseInLocation(layout, str, jst)
	return t
}

func stringToUint(str string) uint {
	Uint32, _ := strconv.ParseUint(str, 10, 32)
	Uint := uint(Uint32)
	return Uint
}

// io.Readerをbyteのスライスに変換
func streamToByte(stream io.Reader) []byte {
	buf := new(bytes.Buffer)
	buf.ReadFrom(stream)
	return buf.Bytes()
}
