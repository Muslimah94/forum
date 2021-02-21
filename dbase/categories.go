package dbase

import (
	"database/sql"
	"fmt"

	"github.com/Muslimah94/forum-back/models"
)

// SelectCategories ...
func (db *DataBase) SelectCategories() ([]models.PostCat, error) {
	rows, err := db.DB.Query(`SELECT PostCats.PostID, CategoryID, Categories.Name FROM PostCats INNER JOIN
	Categories ON PostCats.CategoryID = Categories.ID`)
	if err != nil {
		fmt.Println("SelectCategories Query:", err)
		return nil, err
	}
	defer rows.Close()
	var pc []models.PostCat
	for rows.Next() {
		var p models.PostCat
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

// SelectCategoriesByPostID ...
func (db *DataBase) SelectCategoriesByPostID(id int) ([]models.PostCat, error) {
	rows, err := db.DB.Query(`SELECT PostCats.PostID, CategoryID, Categories.Name FROM PostCats INNER JOIN
	Categories ON PostCats.CategoryID = Categories.ID WHERE PostCats.PostID = ?`, id)
	if err != nil {
		fmt.Println("SelectCategoriesByPostID Query:", err)
		return nil, err
	}
	defer rows.Close()
	var pc []models.PostCat
	for rows.Next() {
		var p models.PostCat
		err = rows.Scan(&p.PostID, &p.CategoryID, &p.CategoryName)
		if err != nil {
			fmt.Println("SelectCategoriesByPostID rows.Scan:", err)
			continue
		}
		pc = append(pc, p)
	}
	if err = rows.Err(); err != nil {
		fmt.Println("SelectCategoriesByPostID rows:", err)
		return nil, err
	}
	return pc, nil
}

// SelectPostIDsByCategory ...
func (db *DataBase) SelectPostIDsByCategory(category string) ([]models.PostCat, error) {
	rows, err := db.DB.Query(`SELECT PostCats.PostID, CategoryID, Categories.Name FROM PostCats INNER JOIN
	Categories ON PostCats.CategoryID = Categories.ID WHERE Categories.Name = ?`, category)
	if err != nil {
		fmt.Println("SelectPostIDsByCategory Query:", err)
		return nil, err
	}
	defer rows.Close()
	var pc []models.PostCat
	for rows.Next() {
		var p models.PostCat
		err = rows.Scan(&p.PostID, &p.CategoryID, &p.CategoryName)
		if err != nil {
			fmt.Println("SelectPostIDsByCategory rows.Scan:", err)
			continue
		}
		pc = append(pc, p)
	}
	if err = rows.Err(); err != nil {
		fmt.Println("SelectPostIDsByCategory rows:", err)
		return nil, err
	}
	return pc, nil
}

// ReturnCategories ...
func (db *DataBase) ReturnCategories() ([]string, error) {
	rows, err := db.DB.Query(`SELECT Name FROM Categories`)
	if err != nil {
		fmt.Println("ReturnCategories Query:", err)
		return nil, err
	}
	defer rows.Close()
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
		fmt.Println("ReturnCategories rows:", err)
		return nil, err
	}
	return cat, nil
}

// AssociateCategory ...
func (db *DataBase) AssociateCategory(pID, cID int, tx *sql.Tx) error {

	st, err := tx.Prepare(`INSERT INTO PostCats (PostID, CategoryID) VALUES (?,?)`)
	if err != nil {
		fmt.Println("AssociateCategory Prepare", err)
		return err
	}
	defer st.Close()
	_, err = st.Exec(pID, cID)
	if err != nil {
		fmt.Println("AssociateCategory Exec", err)
		return err
	}
	return nil
}

// func (db *DataBase) AddNewCategory(cat models.CategoriesDTO) error {
// 	st, err := db.Prepare(`INSERT INTO Categories (Name) VALUES (?)`)
// 	if err != nil {
// 		fmt.Println("AddNewCategory db.Prepare", err)
// 		http.Error(w, err.Error(), http.StatusInternalServerError)
// 		return err
// 	}
// 	_, err = st.Exec(cat.NewCategoryName)
// 	if err != nil {
// 		fmt.Println("AddNewCategory st.Exec", err)
// 		http.Error(w, err.Error(), http.StatusInternalServerError)
// 		return err
// 	}
//    return nil
// }
