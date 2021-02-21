package dbase

import (
	"fmt"

	"github.com/Muslimah94/forum-back/models"
)

// CountReactionsToPost ...
func (db *DataBase) CountReactionsToPost(t int, postID int) (int, error) {
	num := 0
	rows, err := db.DB.Query(`SELECT COUNT(*) FROM Reactions WHERE Type = ? AND PostID = ?`, t, postID)
	if err != nil {
		fmt.Println("CountReactionsToPost Query:", err)
		return 0, err
	}
	defer rows.Close()
	if rows.Next() {
		err = rows.Scan(&num)
		if err != nil {
			fmt.Println("CountReactionsToPost rows.Scan:", err)
		}
	}
	if err = rows.Err(); err != nil {
		fmt.Println("CountReactionsToPost rows:", err)
		return 0, err
	}
	return num, nil
}

// CountReactionsToComment ...
func (db *DataBase) CountReactionsToComment(t int, commentID int) (int, error) {
	num := 0
	rows, err := db.DB.Query(`SELECT COUNT(*) FROM Reactions WHERE Type = ? AND CommentID = ?`, t, commentID)
	if err != nil {
		fmt.Println("CountReactionsToComment Query:", err)
		return 0, err
	}
	defer rows.Close()
	if rows.Next() {
		err = rows.Scan(&num)
		if err != nil {
			fmt.Println("CountReactionsToComment rows.Scan:", err)
		}
	}
	if err = rows.Err(); err != nil {
		fmt.Println("CountReactionsToComment rows:", err)
		return 0, err
	}
	return num, nil
}

// InsertReaction ...
func (db *DataBase) InsertReaction(new models.Reaction) error {

	if new.PostID == 0 {
		st, err := db.DB.Prepare(`INSERT INTO Reactions (AuthorID, Type, CommentID) VALUES (?,?,?)`)
		if err != nil {
			fmt.Println("InsertReaction[comment] Prepare", err)
			return err
		}
		defer st.Close()
		_, err = st.Exec(new.AuthorID, new.Type, new.CommentID)
		if err != nil {
			fmt.Println("InsertReaction[comment] Exec", err)
			return err
		}
	} else {
		st, err := db.DB.Prepare(`INSERT INTO Reactions (AuthorID, Type, PostID) VALUES (?,?,?)`)
		if err != nil {
			fmt.Println("InsertReaction[post] Prepare", err)
			return err
		}
		defer st.Close()
		_, err = st.Exec(new.AuthorID, new.Type, new.PostID)
		if err != nil {
			fmt.Println("InsertReaction[post] Exec", err)
			return err
		}
	}
	return nil
}

// SelectReaction ...
func (db *DataBase) SelectReaction(new models.Reaction) (models.Reaction, error) {
	var existing models.Reaction
	if new.PostID == 0 {
		rows, err := db.DB.Query(`SELECT ID, Type, AuthorID, CommentID FROM Reactions WHERE AuthorID = ? AND CommentID = ?`, new.AuthorID, new.CommentID)
		if err != nil {
			fmt.Println("SelectReaction Query[comment]:", err)
			return existing, err
		}
		defer rows.Close()
		if rows.Next() {
			err = rows.Scan(&existing.ID, &existing.Type, &existing.AuthorID, &existing.CommentID)
			if err != nil {
				fmt.Println("SelectReaction[comment] rows.Scan:", err)
			}
		}
		if err = rows.Err(); err != nil {
			fmt.Println("SelectReaction[comment] rows:", err)
			return existing, err
		}
	} else {
		rows, err := db.DB.Query(`SELECT ID, Type, AuthorID, PostID FROM Reactions WHERE AuthorID = ? AND PostID = ?`, new.AuthorID, new.PostID)
		if err != nil {
			fmt.Println("SelectReaction Query[post]:", err)
			return existing, err
		}
		defer rows.Close()
		if rows.Next() {
			err = rows.Scan(&existing.ID, &existing.Type, &existing.AuthorID, &existing.PostID)
			if err != nil {
				fmt.Println("SelectReaction[post] rows.Scan:", err)
			}
		}
		if err = rows.Err(); err != nil {
			fmt.Println("SelectReaction[post] rows:", err)
			return existing, err
		}
	}
	return existing, nil
}

// UpdateReaction ...
func (db *DataBase) UpdateReaction(new models.Reaction) error {

	if new.PostID == 0 {
		stmt, err := db.DB.Prepare(`UPDATE Reactions SET type = ? WHERE AuthorID = ? AND CommentID = ?`)
		if err != nil {
			fmt.Println("UpdateReaction Prepare[comment]", err)
			return err
		}
		_, err = stmt.Exec(new.Type, new.AuthorID, new.CommentID)
		if err != nil {
			fmt.Println("UpdateReaction Exec[comment]", err)
			return err
		}
		defer stmt.Close()
	} else {
		stmt, err := db.DB.Prepare(`UPDATE Reactions SET type = ? WHERE AuthorID = ? AND PostID = ?`)
		if err != nil {
			fmt.Println("UpdateReaction Prepare[post]", err)
			return err
		}
		_, err = stmt.Exec(new.Type, new.AuthorID, new.PostID)
		if err != nil {
			fmt.Println("UpdateReaction Exec[post]", err)
			return err
		}
		defer stmt.Close()
	}
	return nil
}

// DeleteReaction ...
func (db *DataBase) DeleteReaction(new models.Reaction) error {

	if new.PostID == 0 {
		st, err := db.DB.Prepare(`DELETE FROM Reactions WHERE AuthorID = ? AND Type = ? AND CommentID = ?`)
		if err != nil {
			fmt.Println("DeleteReaction Prepare", err)
			return err
		}
		defer st.Close()
		_, err = st.Exec(new.AuthorID, new.Type, new.CommentID)
		if err != nil {
			fmt.Println("DeleteReaction Exec", err)
			return err
		}
	} else {
		st, err := db.DB.Prepare(`DELETE FROM Reactions WHERE AuthorID = ? AND Type = ? AND PostID = ?`)
		if err != nil {
			fmt.Println("DeleteReaction Prepare", err)
			return err
		}
		defer st.Close()
		_, err = st.Exec(new.AuthorID, new.Type, new.PostID)
		if err != nil {
			fmt.Println("DeleteReaction Exec", err)
			return err
		}
	}
	return nil
}

// SelectLikedPostsIDs ...
func (db *DataBase) SelectLikedPostsIDs(id int) ([]int, error) {
	var ar []int
	rows, err := db.DB.Query(`SELECT PostID FROM Reactions WHERE AuthorID = ? AND Type = ?`, id, 1)
	if err != nil {
		fmt.Println("SelectLikedPostsID Query:", err)
		return ar, err
	}
	defer rows.Close()
	for rows.Next() {
		var a int
		err = rows.Scan(&a)
		if err != nil {
			fmt.Println("SelectLikedPostsID rows.Scan:", err)
			continue
		}
		ar = append(ar, a)
	}
	if err = rows.Err(); err != nil {
		fmt.Println("SelectLikedPostsID rows:", err)
		return ar, err
	}
	return ar, nil
}
