package handlers

import (
	"fmt"
	"net/http"
	"time"

	dbase "../dbase"
	models "../models"
)

// SetCookie ...
func SetCookie(w http.ResponseWriter, r *http.Request, s models.Session) error {
	cookie, _ := r.Cookie("logged-in_forum")
	cookie = &http.Cookie{
		Name:    "logged-in_forum",
		Value:   s.UUID.String(),
		Expires: time.Now().Add(time.Hour * 1),
		Path:    "/",
		//Secure:   true,
		//HttpOnly: true,
	}
	http.SetCookie(w, cookie)
	return nil
}

// CheckCookie ...
func CheckCookie(r *http.Request, exisSes models.Session) bool {
	cookie, err := r.Cookie("logged-in_forum")
	if err == http.ErrNoCookie {
		fmt.Println("Check cookie", err)
		return false
	}
	if cookie.Value == "" {
		fmt.Println("Check cookie: cookie.Value is empty")
		return false
	}
	if cookie.Value != exisSes.UUID.String() {
		fmt.Println("Check cookie: UUID doesn't match")
		return false
	}
	return true
}

// GetUserIDBySession ...
func GetUserIDBySession(db *dbase.DataBase, r *http.Request) (int, error) {
	cookie, err := r.Cookie("logged-in_forum")
	if err == http.ErrNoCookie {
		return 0, err
	}
	UUID := cookie.Value
	id, err := db.SelectUserIDBySession(UUID)
	if err != nil {
		return 0, err
	}

	return id, nil
}

// DeleteCookie ...
func DeleteCookie(w http.ResponseWriter, r *http.Request) error {
	cookie, err := r.Cookie("logged-in_forum")
	if err != nil {
		fmt.Println("DeleteCookie:", err)
		return err
	}
	cookie = &http.Cookie{
		MaxAge: -1,
	}
	http.SetCookie(w, cookie)
	return nil
}
