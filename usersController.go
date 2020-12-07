package main

import (
	"log"
	"net/http"

	"github.com/gorilla/sessions"
	"github.com/labstack/echo-contrib/session"
	"github.com/labstack/echo/v4"

	"golang.org/x/crypto/bcrypt"
)

func signup(c echo.Context) error {
	db := sqlConnect()
	defer db.Close()

	u := new(User)
	if err := c.Bind(u); err != nil {
		return c.JSON(http.StatusInternalServerError, "Faild Bind")
	}

	u.PasswordDigest = passwordHash(u.PasswordDigest)

	if c.FormValue("test") == "true" {
		tx := db.Begin()
		err := tx.Create(&u).Error
		tx.Rollback()
		if err != nil {
			return c.NoContent(http.StatusInternalServerError)
		}
		return c.JSON(http.StatusOK, "testok")
	}
	err := db.Create(&u).Error
	if err != nil {
		return err
	}
	return c.JSON(http.StatusOK, "ok")
}

func signin(c echo.Context) error {
	db := sqlConnect()
	defer db.Close()
	user := User{}
	u := new(UserSignin)
	if err := c.Bind(u); err != nil {
		return c.JSON(http.StatusInternalServerError, "Faild Bind")
	}

	db.Where("email = ?", u.Email).Find(&user)

	passcheck := bcrypt.CompareHashAndPassword([]byte(user.PasswordDigest), []byte(u.Password))

	if passcheck == nil {
		sess, _ := session.Get("session", c)
		sess.Options = &sessions.Options{
			Path:     "/",
			MaxAge:   86400 * 7,
			HttpOnly: true,
			SameSite: http.SameSiteNoneMode,
			Secure:   true,
		}
		sess.Values["authenticated"] = true
		if err := sess.Save(c.Request(), c.Response()); err != nil {
			return c.NoContent(http.StatusInternalServerError)
		}

		return c.JSON(http.StatusOK, UserSend{UserID: user.ID, Username: user.Username})
	}
	return passcheck
}

func signout(c echo.Context) error {
	sess, _ := session.Get("session", c)
	sess.Options = &sessions.Options{MaxAge: -1, Path: "/"}
	if err := sess.Save(c.Request(), c.Response()); err != nil {
		return c.NoContent(http.StatusInternalServerError)
	}
	return c.NoContent(http.StatusOK)
}

func checkSession(c echo.Context) error {
	sess, _ := session.Get("session", c)
	log.Print(sess.Values["authenticated"])
	if b, _ := sess.Values["authenticated"]; b != true {
		return c.String(http.StatusUnauthorized, "401")
	}
	return c.JSON(http.StatusOK, "200")

}

func passwordHash(pw string) string {
	hash, err := bcrypt.GenerateFromPassword([]byte(pw), bcrypt.DefaultCost)
	if err != nil {
		panic(err)
	}
	return string(hash)
}
