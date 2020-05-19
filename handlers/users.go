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
	ID, err := db.InsertUser(user, tx)
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
		ID:             int(ID),
		Email:          new.Email,
		HashedPassword: string(HashedPW),
	}
	err = db.InsertUserCredentials(cred, tx)
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
	session := models.Session{UserID: int(ID)}
	session.UUID, err = db.CreateSession(session, tx)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		tx.Rollback()
		return
	}
	err = SetCookie(w, r, session)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		tx.Rollback()
		return
	}
	tx.Commit()
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

	exisSes, err := db.SelectUserSession(session) // Checking is there a session with the given UserID
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	tx, err := db.DB.Begin() // Starting a transaction in DB
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if exisSes.ID == 0 { // if there's no session, we'll create it and set cookie
		exisSes.UUID, err = db.CreateSession(session, tx)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		SetCookie(w, r, exisSes)
		tx.Commit()
		return
	} else if exisSes.ID != 0 && CheckCookie(r) { // ExpDate need to be updated when user logs in from the same browser
		exisSes.ExpDate = time.Now().Add(time.Hour * 1).Unix()
		err = db.UpdateSessionDate(exisSes, tx)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		tx.Commit()
		err = SetCookie(w, r, exisSes) // Cookie with updated life time should be set
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	} else if exisSes.ID != 0 && !CheckCookie(r) { // UUID & ExpDate need to be updated when user logs in from another browser
		exisSes.ExpDate = time.Now().Add(time.Hour * 1).Unix()
		exisSes.UUID, err = uuid.NewV4()
		err = db.UpdateSession(exisSes, tx)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		tx.Commit()
		err = SetCookie(w, r, exisSes) // Cookie with updated life time and UUID should be set in order to prevent authorization from earlier browser
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	}

}
