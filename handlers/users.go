package handlers

import (
	"fmt"
	"net/http"
	"time"

	dbase "../dbase"
	models "../models"
	uuid "github.com/satori/go.uuid"
	"golang.org/x/crypto/bcrypt"
)

// RegisterLogin ...
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
		RoleID:   3, // role:"user"
	}
	tx, err := db.DB.Begin()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		fmt.Println("Cannot start transaction")
		return
	}
	ID, err := db.CreateUser(user, tx)
	if err != nil && err.Error()[:6] == "UNIQUE" {
		SendJSON(w, models.Error{
			Status:      "Failed",
			Description: "User with such a nickname already exists, please try another one",
		})
		tx.Rollback()
		return
	} else if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		tx.Rollback()
		return
	}
	HashedPW, err := bcrypt.GenerateFromPassword([]byte(new.Password), bcrypt.MinCost)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		tx.Rollback()
		return
	}
	//---------ENTITY for Credentials table---------------
	cred := models.Credentials{
		ID:             ID,
		Email:          new.Email,
		HashedPassword: string(HashedPW),
	}
	err = db.CreateUserCredentials(cred, tx)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		tx.Rollback()
		return
	}
	session := models.Session{UserID: ID}
	UUID, err := db.CreateSession(session, tx)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		tx.Rollback()
		return
	}
	err = SetCookie(w, r, UUID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		tx.Rollback()
		return
	}
	return
}

// SetCookie ...
func SetCookie(w http.ResponseWriter, r *http.Request, UUID uuid.UUID) error {
	cookie, err := r.Cookie("logged-in_forum")
	if err == http.ErrNoCookie {
		cookie = &http.Cookie{
			Name:     "logged-in_forum",
			Value:    UUID.String(),
			Expires:  time.Now().Add(time.Hour * 1),
			Secure:   true,
			HttpOnly: true,
		}
	}
	http.SetCookie(w, cookie)
	return nil
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
