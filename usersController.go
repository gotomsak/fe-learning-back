package main

import (
	"github.com/labstack/echo"

	"golang.org/x/crypto/bcrypt"
)

func signup(c echo.Context) error {
	db := sqlConnect()
	defer db.Close()
	user := User{}

	username := c.FormValue("username")
	password := c.FormValue("password")
	email := c.FormValue("email")
	hash := passwordHash(password)
	user.Username = string(username)
	user.PasswordDigest = hash
	user.Email = string(email)
	error := db.Create(&user).Error
	return error
}

func passwordHash(pw string) string {
	hash, err := bcrypt.GenerateFromPassword([]byte(pw), bcrypt.DefaultCost)
	if err != nil {
		panic(err)
	}
	return string(hash)
}
