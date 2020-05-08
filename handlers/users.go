package handlers

import (
	"fmt"
	"net/http"

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

}
