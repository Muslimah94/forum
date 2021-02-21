package handlers

import (
	"fmt"
	"net/http"
	"time"

	"github.com/Muslimah94/forum-back/dbase"
	"github.com/Muslimah94/forum-back/models"
	uuid "github.com/satori/go.uuid"
	"golang.org/x/crypto/bcrypt"
)

// RegisterLogin ...
func RegisterLogin(db *dbase.DataBase, w http.ResponseWriter, r *http.Request) {
	var new models.RegisterUser
	err := ReceiveJSON(r, &new)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	if !validNick(new.Nickname) {
		SendJSON(w, models.Error{
			Status:      "Failed",
			Description: "Allowed nickname length is 5-20 symbols. Nickname should contain the Latin alphabet, numbers and given special characters only: `-=~!@#$%^&*()_+\\|/? {}[]",
		})
		return
	}
	if !validPass(new.Password) {
		SendJSON(w, models.Error{
			Status:      "Failed",
			Description: "Allowed password length is 5-20 symbols. Password should contain at least 1 lowercase letter, 1 uppercase letter, 1 number and 1 special character",
		})
		return
	}
	if !validEmail(new.Email) {
		SendJSON(w, models.Error{
			Status:      "Failed",
			Description: "Invaid email address, try another one or check yours for mistakes",
		})
		return
	}
	//--------ENTITY for Users table----------------------
	user := models.User{
		Nickname: new.Nickname,
		RoleID:   3, // role:"user"
	}
	HashedPW, err := bcrypt.GenerateFromPassword([]byte(new.Password), bcrypt.MinCost)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
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
			Description: "User with such an email already exists, please try another one",
		})
		tx.Rollback()
		return
	} else if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		tx.Rollback()
		return
	}
	session := models.Session{UserID: int(ID)}
	session.UUID, err = db.InsertSession(session, tx)
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
	cred := models.Credentials{
		Email:          new.Email,
		HashedPassword: new.Password,
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
	err = bcrypt.CompareHashAndPassword([]byte(exisCr.HashedPassword), []byte(new.Password))
	if err != nil {
		fmt.Println(err)
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
	fmt.Println(exisSes)
	if exisSes.ID == 0 { // if there's no session, we'll create it and set cookie
		exisSes.UUID, err = db.InsertSession(session, tx)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		tx.Commit()
		SetCookie(w, r, exisSes)
		return
	}
	exisSes.ExpDate = time.Now().Add(time.Hour * 1).Unix()
	if CheckCookie(r, exisSes) { //if browser is the same ExpDate need to be updated only
		err = db.UpdateSessionDate(exisSes, tx)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	} else { // if browser isn't the same, session need to be updated totally
		exisSes.UUID = uuid.NewV4()
		err = db.UpdateSession(exisSes, tx)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	}
	tx.Commit()
	err = SetCookie(w, r, exisSes) // Cookie with updated life time and/or new UUID should be set
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

// LogOut ...
func LogOut(db *dbase.DataBase, w http.ResponseWriter, r *http.Request) {
	cookie, err := r.Cookie("logged-in_forum")
	if err != nil {
		fmt.Println("LogOut, cookie:", err)
		return
	}
	err = db.DeleteSession(cookie.Value)
	if err != nil {
		fmt.Println("LogOut, cannot delete session:", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	return
}
