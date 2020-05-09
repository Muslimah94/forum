package handlers

import (
	"fmt"
	"net/http"
	"time"

	dbase "../dbase"
	models "../models"
	"golang.org/x/crypto/bcrypt"
)

func RegisterLogin(db *dbase.DataBase, w http.ResponseWriter, r *http.Request) {
	//-------DTO----------------------------------------
	var new models.RegisterUser
	err := ReceiveJSON(r, &new)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	//--------ENTITY for Users table----------------------
	user := models.User{
		Nickname: new.Nickname,
		RoleID:   3,
	}
	ID, err := db.CreateUser(user)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	//---------ENTITY for Credentials table---------------
	HashedPW, err := bcrypt.GenerateFromPassword([]byte(new.Password), bcrypt.MinCost)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	cred := models.Credentials{
		ID:             ID,
		Email:          new.Email,
		HashedPassword: string(HashedPW),
	}
	err = db.CreateUserCredentials(cred)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	session := models.Session{UserID: ID}
	UUID, err := db.CreateSession(session)
	fmt.Println("Last created session's UUID:", UUID)
	err = SetCookie(w, r)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func SetCookie(w http.ResponseWriter, r *http.Request) error {
	cookie, err := r.Cookie("logged-in_forum")
	if err == http.ErrNoCookie {
		cookie = &http.Cookie{
			Name:     "logged-in_forum",
			Value:    "1",
			Expires:  time.Now().Add(time.Hour * 1),
			Secure:   true,
			HttpOnly: true,
		}
	}
	http.SetCookie(w, cookie)
	return nil
}

func DeleteCookie(w http.ResponseWriter, r *http.Request) error {
	cookie, err := r.Cookie("logged-in_forum")
	if err != nil {
		fmt.Println("DeleteCookie error:")
		return err
	}
	cookie = &http.Cookie{
		Name:     "logged-in_forum",
		Value:    "1",
		MaxAge:   -1,
		Secure:   true,
		HttpOnly: true,
	}
	http.SetCookie(w, cookie)
	return nil
}
