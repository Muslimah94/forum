package dbase

import (
	"fmt"

	models "../models"
)

func (db *DataBase) SelectCategories() ([]models.PostCategories, error) {
	rows, err := db.DB.Query(`SELECT PostsCategories.PostID, CategoryID, Categories.Name FROM PostsCategories INNER JOIN
	Categories ON PostsCategories.CategoryID = Categories.ID`)
	defer rows.Close()
	if err != nil {
		fmt.Println("SelectCategories Query:", err)
		return nil, err
	}
	var pc []models.PostCategories
	for rows.Next() {
		var p models.PostCategories
		err = rows.Scan(&p.PostID, &p.CategoryID, &p.CategoryName)
		if err != nil {
			fmt.Println("SelectCategories rows.Scan:", err)
			continue
		}
		pc = append(pc, p)
	}
	if err = rows.Err(); err != nil {
		fmt.Println("SelectCategories rows:", err)
		return nil, err
	}
	return pc, nil
}

func (db *DataBase) SelectCategoriesByPostID(id int) ([]models.PostCategories, error) {
	rows, err := db.DB.Query(`SELECT PostsCategories.PostID, CategoryID, Categories.Name FROM PostsCategories INNER JOIN
	Categories ON PostsCategories.CategoryID = Categories.ID WHERE PostsCategories.PostID = ?`, id)
	defer rows.Close()
	if err != nil {
		fmt.Println("SelectCategories Query:", err)
		return nil, err
	}
	var pc []models.PostCategories
	for rows.Next() {
		var p models.PostCategories
		err = rows.Scan(&p.PostID, &p.CategoryID, &p.CategoryName)
		if err != nil {
			fmt.Println("SelectCategories rows.Scan:", err)
			continue
		}
		pc = append(pc, p)
	}
	if err = rows.Err(); err != nil {
		fmt.Println("SelectCategories rows:", err)
		return nil, err
	}
	return pc, nil
}

func (db *DataBase) ReturnCategories() ([]string, error) {
	rows, err := db.DB.Query(`SELECT Name FROM Categories`)
	defer rows.Close()
	if err != nil {
		fmt.Println("ReturnCategories Query:", err)
		return nil, err
	}
	cat := []string{}
	for rows.Next() {
		var a string
		err = rows.Scan(&a)
		if err != nil {
			fmt.Println("ReturnCategories rows.Scan:", err)
			continue
		}
		cat = append(cat, a)
	}
	if err = rows.Err(); err != nil {
		fmt.Println("SelectCategories rows:", err)
		return nil, err
	}
	return cat, nil
}

func (db *DataBase) AssociateCategory(pID, cID int) error {

	st, err := db.DB.Prepare(`INSERT INTO PostsCategories (PostID, CategoryID) VALUES (?,?)`)
	defer st.Close()
	if err != nil {
		fmt.Println("CreatePost Prepare", err)
		return err
	}
	_, err = st.Exec(pID, cID)
	if err != nil {
		fmt.Println("CreatePost Exec", err)
		return err
	}
	return nil
}

// func AddNewCategory(db *sql.DB, w http.ResponseWriter, r *http.Request) {
// 	var cat *models.Categories
// 	handlers.ReceiveJSON(r, &cat)

// 	st, err1 := db.Prepare(`INSERT INTO Categories (Name) VALUES (?)`)
// 	if err1 != nil {
// 		fmt.Println("AddNewCategory db.Prepare", err1)
// 		http.Error(w, err1.Error(), http.StatusInternalServerError)
// 		return
// 	}
// 	_, err3 := st.Exec(cat.Name)
// 	if err3 != nil {
// 		fmt.Println("AddNewCategory st.Exec", err3)
// 		http.Error(w, err3.Error(), http.StatusInternalServerError)
// 		return
// 	}
// }

// func AddPostCategories(db *sql.DB, w http.ResponseWriter, r *http.Request) {
// 	var cat *models.PostsCategories
// 	handlers.ReceiveJSON(r, &cat)

// 	st, err1 := db.Prepare(`INSERT INTO PostsCategories (PostID,CategoryID) VALUES (?,?)`)
// 	if err1 != nil {
// 		fmt.Println("AddNewPostCategory db.Prepare", err1)
// 		http.Error(w, err1.Error(), http.StatusInternalServerError)
// 		return
// 	}
// 	_, err3 := st.Exec(cat.PostID, cat.CategoryID)
// 	if err3 != nil {
// 		fmt.Println("AddNewPostCategory st.Exec", err3)
// 		http.Error(w, err3.Error(), http.StatusInternalServerError)
// 		return
// 	}
// }
