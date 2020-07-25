package main

import (
	"math/rand"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

func sqlConnect() (database *gorm.DB) {
	DBMS := os.Getenv("DBMS")
	USER := os.Getenv("USERR")
	PASS := os.Getenv("PASS")
	PROTOCOL := os.Getenv("PROTOCOL")
	DBNAME := os.Getenv("DBNAME")
	CONNECT := USER + ":" + PASS + "@" + PROTOCOL + "/" + DBNAME + "?charset=utf8mb4&parseTime=true&loc=Asia%2FTokyo"
	db, err := gorm.Open(DBMS, CONNECT)
	db.Set("gorm:table_options", "ENGINE=InnoDB").AutoMigrate(&User{})
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
func searchIDs(ints []int, search int) bool {
	for _, v := range ints {
		if v == search {
			return true
		}
	}
	return false
}

// string型で受け取った数値のリストをInt型のリストにして返す
func strToIntList(str string) []int {
	intList := []int{}
	str = strings.Trim(str, "[]")
	strList := strings.Split(str, ",")
	if str != "" {
		for i := 0; i < len(strList); i++ {
			n, _ := strconv.Atoi(strList[i])
			intList = append(intList, n)
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
	t, _ := time.Parse(layout, str)
	return t
}

func stringToUint(str string) uint {
	Uint32, _ := strconv.ParseUint(str, 10, 32)
	Uint := uint(Uint32)
	return Uint
}
