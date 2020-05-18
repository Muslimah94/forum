package handlers

import (
	"fmt"
	"net/http"

	dbase "../dbase"
	models "../models"
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

// LogIn ...
func LogIn(db *dbase.DataBase, w http.ResponseWriter, r *http.Request) {
	//-------DTO----------------------------------------
	var new models.CredDTO
	err := ReceiveJSON(r, &new)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	//--------ENTITY for Credentials table----------------------
	HashedPW, err := bcrypt.GenerateFromPassword([]byte(new.Password), bcrypt.MinCost)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	cred := models.Credentials{
		Email:          new.Email,
		HashedPassword: string(HashedPW),
	}
	exisCr, err := db.SelectUserCredentials(cred)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	if exisCr.ID == 0 {
		SendJSON(w, models.Error{
			Status:      "Failed to login",
			Description: "Email or password is incorrect",
		})
		return
	}
	if cred.HashedPassword != exisCr.HashedPassword {
		SendJSON(w, models.Error{
			Status:      "Failed to login",
			Description: "Email or password is incorrect",
		})
		return
	}
	session := models.Session{UserID: exisCr.ID}
	// Checking is there a session with given UserID
	exisSes, err := db.SelectUserSession(session)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	tx, err := db.DB.Begin()
	// if there's no session, we'll create it and set cookie
	if exisSes.ID == 0 {
		UUID, err := db.CreateSession(session, tx)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		SetCookie(w, r, UUID)
		return
	} else if exisSes.ID != 0 && CheckCookie(r) {
		//db.UpdateSessionDate(exisSes, tx) // ExpDate is updated
	} else if exisSes.ID != 0 && !CheckCookie(r) {
		//db.UpdateSession(exisSes, tx)
	}
	err = SetCookie(w, r, UUID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

}
