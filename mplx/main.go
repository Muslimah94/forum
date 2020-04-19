package mplx

import (
	"fmt"
	"net/http"

	dbase "../dbase"
)

func Multiplexer(w http.ResponseWriter, r *http.Request) {
	fmt.Println("multi")
	db, err := dbase.Create("forumDB")
	defer db.Close()
	if err != nil {
		fmt.Println("sql.Open ERROR:", err)
	}

	// p := r.URL.Path
	// m := r.Method

	// if p == "/api/user" {
	// 	if m == "GET" {
	// 		dbase.GetAllUsers(db, w, r)
	// 	} else if m == "POST" {
	// 		dbase.AddNewUser(db, w, r)
	// 	}
	// } else if p[0:9] == "/api/user" && len(p) > 9 {
	// 	userID, err := strconv.Atoi(p[10:])
	// 	if err != nil {
	// 		fmt.Println("USER Atoi ERROR:", err.Error())
	// 	}
	// 	if m == "GET" {
	// 		dbase.GetUserByID(db, w, r, userID)
	// 	} else if m == "PUT" {
	// 		dbase.EditUserByID(db, w, r, userID)
	// 	} else if m == "DELETE" {
	// 		dbase.DeleteUserByID(db, w, r, userID)
	// 	}
	// } else if p == "/api/post" {
	// 	if m == "GET" {
	// 		dbase.GetAllPosts(db, w, r)
	// 	} else if m == "POST" {
	// 		dbase.AddNewPost(db, w, r)
	// 	}
	// } else if p[0:9] == "/api/post" && len(p) > 9 {
	// 	postID, err := strconv.Atoi(p[10:])
	// 	if err != nil {
	// 		fmt.Println("USER Atoi ERROR:", err.Error())
	// 	}
	// 	if m == "GET" {
	// 		dbase.GetPostByID(db, w, r, postID)
	// 	} else if m == "PUT" {
	// 		dbase.EditPostByID(db, w, r, postID)
	// 	} else if m == "DELETE" {
	// 		dbase.DeletePostByID(db, w, r, postID)
	// 	}
	// } else if p[0:12] == "/api/comment" && len(p) > 12 {
	// 	if p[0:19] == "/api/comment/postID" {
	// 		postID, err := strconv.Atoi(p[20:])
	// 		if err != nil {
	// 			fmt.Println("Comment Atoi Error:", err.Error())
	// 		}
	// 		if m == "GET" {
	// 			dbase.GetAllCommentsToPost(db, w, r, postID)
	// 		} else if m == "POST" {
	// 			dbase.AddNewCommentToPost(db, w, r, postID)
	// 		}
	// 	} else if p[0:22] == "/api/comment/commentID" {
	// 		commentID, err := strconv.Atoi(p[23:])
	// 		if err != nil {
	// 			fmt.Println("Comment Atoi Error:", err.Error())
	// 		}
	// 		if m == "GET" {
	// 			dbase.GetCommentByID(db, w, r, commentID)
	// 		} else if m == "PUT" {
	// 			dbase.EditCommentByID(db, w, r, commentID)
	// 		} else if m == "DELETE" {
	// 			dbase.DeleteCommentByID(db, w, r, commentID)
	// 		}
	// 	}

	// } else if p == "/api/reaction" {
	// 	//DBinteraction.AddReaction(db, w, r)
	// } else if p[0:14] == "/api/reaction/" && len(p) > 14 {
	// 	if p[0:19] == "/api/reaction/post/" {
	// 		postID, err := strconv.Atoi(p[19:])
	// 		if err != nil {
	// 			fmt.Println("USER Atoi ERROR:", err.Error())
	// 		}
	// 		if m == "GET" {
	// 			fmt.Println("postID:", postID)
	// 			dbase.GetAllReactionsToPost(db, w, r, postID)
	// 		}
	// 	} else if p[0:22] == "/api/reaction/comment/" {

	// 	} else if p[0:14] == "/api/reaction/" {
	// 	}
	// }

}
