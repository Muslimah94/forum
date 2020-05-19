package handlers

import (
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
func CheckCookie(r *http.Request) bool {
	cookie, err := r.Cookie("logged-in_forum")
	if err == http.ErrNoCookie || cookie.Value == "" {
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
