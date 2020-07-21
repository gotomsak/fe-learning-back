package main

import (
	"net/http"

	"github.com/gorilla/sessions"
	"github.com/labstack/echo-contrib/session"
	"github.com/labstack/echo/v4"

	"golang.org/x/crypto/bcrypt"
)

func signup(c echo.Context) error {
	db := sqlConnect()
	defer db.Close()
	user := User{}

	password := c.FormValue("password")
	hash := passwordHash(password)
	user.Username = c.FormValue("username")
	user.PasswordDigest = hash
	user.Email = c.FormValue("email")
	if c.FormValue("test") == "true" {
		tx := db.Begin()
		error := tx.Create(&user).Error
		tx.Rollback()
		return error
	}
	error := db.Create(&user).Error
	return error
}

func signin(c echo.Context) error {
	db := sqlConnect()
	defer db.Close()
	user := User{}
	db.Where("email = ?", c.FormValue("email")).Find(&user)
	passDigest := c.FormValue("password")

	passcheck := bcrypt.CompareHashAndPassword([]byte(user.PasswordDigest), []byte(passDigest))

	if passcheck == nil {
		sess, _ := session.Get("session", c)
		sess.Options = &sessions.Options{
			Path:     "/",
			MaxAge:   86400 * 7,
			HttpOnly: true,
		}
		sess.Save(c.Request(), c.Response())
		return c.NoContent(http.StatusOK)
	}
	return passcheck
}

func passwordHash(pw string) string {
	hash, err := bcrypt.GenerateFromPassword([]byte(pw), bcrypt.DefaultCost)
	if err != nil {
		panic(err)
	}
	return string(hash)
}
