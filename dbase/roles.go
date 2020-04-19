package dbase

import (
	"database/sql"
	"fmt"
	"net/http"

	models "../models"
)

// GetAllRoles ...
func GetAllRoles(db *sql.DB, w http.ResponseWriter, r *http.Request) {
	rows, err1 := db.Query(`SELECT * FROM Roles`)
	if err1 != nil {
		fmt.Println("GetAllRoles db.Query ERROR:", err1)
		http.Error(w, err1.Error(), http.StatusInternalServerError)
		return
	}
	defer rows.Close()
	var AllRoles []models.Roles
	for rows.Next() {
		var u models.Roles
		err2 := rows.Scan(&u.ID, &u.Name)
		if err2 != nil {
			fmt.Println("GetAllUsers rows.Scan ERROR:", err2)
			continue
		}
		AllRoles = append(AllRoles, u)
	}
	if err3 := rows.Err(); err3 != nil {
		fmt.Println("GetAllUsers rows ERROR:", err3)
		http.Error(w, err3.Error(), http.StatusInternalServerError)
		return
	}
	SendJSON(w, &AllRoles)
}

// AddNewRole ...
func AddNewRole(db *sql.DB, w http.ResponseWriter, r *http.Request) {
	//check session (true false) { est' ili net, timestamp < time.Now {
	// 	delete session
	// } else {
	// update current+3600
	// set cookie MaxAge 3600
	// *******
	// 	}
	// est' li net , redirect

	var role *models.Roles
	ReceiveJSON(r, &role)
	st, err2 := db.Prepare(`INSERT INTO Roles (Name) VALUES (?)`)
	if err2 != nil {
		fmt.Println("AddNewRole db.Prepare", err2)
		http.Error(w, err2.Error(), http.StatusInternalServerError)
		return
	}
	_, err3 := st.Exec(role.Name)
	if err3 != nil {
		fmt.Println("AddNewRole st.Exec", err3)
		http.Error(w, err3.Error(), http.StatusInternalServerError)
		return
	}
}

// EditRoleByID ...
func EditRoleByID(db *sql.DB, w http.ResponseWriter, r *http.Request, roleID int) {
	var new *models.Roles
	ReceiveJSON(r, &new)
	st, err2 := db.Prepare(`UPDATE Roles SET Name = ? where ID = ?`)
	if err2 != nil {
		fmt.Println("EditRoleByID db.Prepare:", err2)
		http.Error(w, err2.Error(), http.StatusInternalServerError)
		return
	}
	_, err3 := st.Exec(new.Name, roleID)
	if err3 != nil {
		fmt.Println("EditRoleByID st.Exec:", err3)
		http.Error(w, err3.Error(), http.StatusInternalServerError)
		return
	}
}

// DeleteRoleByID ...
func DeleteRoleByID(db *sql.DB, w http.ResponseWriter, r *http.Request, roleID int) {

	st, err1 := db.Prepare(`DELETE FROM Roles WHERE ID = ?`)
	if err1 != nil {
		fmt.Println("DeleteRoleByID db.Prepare:", err1)
		http.Error(w, err1.Error(), http.StatusInternalServerError)
		return
	}
	_, err2 := st.Exec(roleID)
	if err2 != nil {
		fmt.Println("DeleteRoleByID st.Exec:", err2)
		http.Error(w, err2.Error(), http.StatusInternalServerError)
		return
	}
}
