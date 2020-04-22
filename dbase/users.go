package dbase

// import (
// 	"database/sql"
// 	"fmt"
// 	"net/http"

// 	models "../models"

// 	handlers "../handlers"
// )

// // GetAllUsers ...
// func GetAllUsers(db *sql.DB, w http.ResponseWriter, r *http.Request) {
// 	rows, err1 := db.Query(`SELECT * FROM Users`)
// 	if err1 != nil {
// 		fmt.Println("GetAllUsers db.Query ERROR:", err1)
// 		http.Error(w, err1.Error(), http.StatusInternalServerError)
// 		return
// 	}
// 	defer rows.Close()
// 	var AllUsers []models.Users
// 	for rows.Next() {
// 		var u models.Users
// 		err2 := rows.Scan(&u.ID, &u.Email, &u.Nickname, &u.Password, &u.RoleID)
// 		if err2 != nil {
// 			fmt.Println("GetAllUsers rows.Scan ERROR:", err2)
// 			continue
// 		}
// 		AllUsers = append(AllUsers, u)
// 	}
// 	if err3 := rows.Err(); err3 != nil {
// 		fmt.Println("GetAllUsers rows ERROR:", err3)
// 		http.Error(w, err3.Error(), http.StatusInternalServerError)
// 		return
// 	}
// 	handlers.SendJSON(w, AllUsers)
// }

// // AddNewUser ...
// func AddNewUser(db *sql.DB, w http.ResponseWriter, r *http.Request) {

// 	//check session (true false) { est' ili net, timestamp < time.Now {
// 	// 	delete session
// 	// } else {
// 	// update current+3600
// 	// set cookie MaxAge 3600
// 	// *******
// 	// 	}
// 	// est' li net , redirect

// 	var user *models.Users
// 	handlers.ReceiveJSON(r, &user)
// 	st, err2 := db.Prepare(`INSERT INTO Users (Email, Nickname, Password, RoleID) VALUES (?,?,?,?)`)
// 	if err2 != nil {
// 		fmt.Println("AddNewUser db.Prepare", err2)
// 		http.Error(w, err2.Error(), http.StatusInternalServerError)
// 		return
// 	}
// 	_, err3 := st.Exec(user.Email, user.Nickname, user.Password, user.RoleID)
// 	if err3 != nil {
// 		fmt.Println("AddNewUser st.Exec", err3)
// 		http.Error(w, err3.Error(), http.StatusInternalServerError)
// 		return
// 	}
// }

// // GetUserByID ...
// func GetUserByID(db *sql.DB, w http.ResponseWriter, r *http.Request, userID int) {
// 	var user models.Users
// 	rows := db.QueryRow(`SELECT * FROM Users WHERE ID = $1`, userID)
// 	err := rows.Scan(&user.ID, &user.Email, &user.Nickname, &user.Password, &user.RoleID)
// 	if err != nil {
// 		fmt.Println("GetUserByI rows.Scan ERROR:", err)
// 		http.Error(w, err.Error(), http.StatusInternalServerError)
// 		return
// 	}
// 	handlers.SendJSON(w, user)
// }

// // EditUserByID ...
// func EditUserByID(db *sql.DB, w http.ResponseWriter, r *http.Request, userID int) {
// 	var new *models.Users
// 	handlers.ReceiveJSON(r, &new)
// 	st, err2 := db.Prepare(`UPDATE Users SET Email = ?, Nickname = ?, Password = ?, RoleID = ? where ID = ?`)
// 	if err2 != nil {
// 		fmt.Println("EditUserByID db.Prepare:", err2)
// 		http.Error(w, err2.Error(), http.StatusInternalServerError)
// 		return
// 	}
// 	_, err3 := st.Exec(new.Email, new.Nickname, new.Password, new.RoleID, userID)
// 	if err3 != nil {
// 		fmt.Println("EditUserByID st.Exec:", err3)
// 		http.Error(w, err3.Error(), http.StatusInternalServerError)
// 		return
// 	}
// }

// // DeleteUserByID ...
// func DeleteUserByID(db *sql.DB, w http.ResponseWriter, r *http.Request, userID int) {

// 	st, err1 := db.Prepare(`DELETE FROM Users WHERE ID = ?`)
// 	if err1 != nil {
// 		fmt.Println("DeleteUserByID db.Prepare:", err1)
// 		http.Error(w, err1.Error(), http.StatusInternalServerError)
// 		return
// 	}
// 	_, err2 := st.Exec(userID)
// 	if err2 != nil {
// 		fmt.Println("DeleteUserByID st.Exec:", err2)
// 		http.Error(w, err2.Error(), http.StatusInternalServerError)
// 		return
// 	}
// }

// // GetUsersByRoleID ...
// func GetUsersByRoleID(db *sql.DB, w http.ResponseWriter, r *http.Request, roleID int) {
// 	rows, err1 := db.Query(`SELECT * FROM Users WHERE RoleID = ?`, roleID)
// 	if err1 != nil {
// 		fmt.Println("GetUsersByRoleID db.Query ERROR:", err1)
// 		http.Error(w, err1.Error(), http.StatusInternalServerError)
// 		return
// 	}
// 	defer rows.Close()
// 	var AllUsers []models.Users
// 	for rows.Next() {
// 		var u models.Users
// 		err2 := rows.Scan(&u.ID, &u.Email, &u.Nickname, &u.Password, &u.RoleID)
// 		if err2 != nil {
// 			fmt.Println("GetUsersByRoleID rows.Scan ERROR:", err2)
// 			continue
// 		}
// 		AllUsers = append(AllUsers, u)
// 	}
// 	if err3 := rows.Err(); err3 != nil {
// 		fmt.Println("GetUsersByRoleID rows ERROR:", err3)
// 		http.Error(w, err3.Error(), http.StatusInternalServerError)
// 		return
// 	}
// 	handlers.SendJSON(w, AllUsers)
// }
