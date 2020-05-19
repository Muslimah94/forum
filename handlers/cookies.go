package handlers

import (
	"fmt"
	"net/http"
	"time"

	models "../models"
)

// SetCookie ...
func SetCookie(w http.ResponseWriter, r *http.Request, s models.Session) error {
	cookie, err := r.Cookie("logged-in_forum")
	if err != nil {
		cookie = &http.Cookie{
			Name:     "logged-in_forum",
			Value:    s.UUID.String(),
			Expires:  time.Now().Add(time.Hour * 1),
			Secure:   true,
			HttpOnly: true,
		}
	}
	cookie = &http.Cookie{
		Name:     "logged-in_forum",
		Value:    s.UUID.String(),
		Expires:  time.Now().Add(time.Hour * 1),
		Secure:   true,
		HttpOnly: true,
	}
	http.SetCookie(w, cookie)
	return nil
}

// CheckCookie ...
func CheckCookie(r *http.Request, exisSes models.Session) bool {
	cookie, err := r.Cookie("logged-in_forum")
	if err == http.ErrNoCookie {
		fmt.Println("no cookie")
		return false
	}
	if cookie.Value == "" {
		fmt.Println("empty value cookie")
		return false
	}
	if cookie.Value != exisSes.UUID.String() {
		fmt.Println("doesn't match uuid")
		return false
	}
	return true
}

// DeleteCookie ...
// func DeleteCookie(w http.ResponseWriter, r *http.Request) error {
// 	cookie, err := r.Cookie("logged-in_forum")
// 	if err != nil {
// 		fmt.Println("DeleteCookie:", err)
// 		return err
// 	}
// 	cookie = &http.Cookie{
// 		MaxAge: -1,
// 	}
// 	http.SetCookie(w, cookie)
// 	return nil
// }
