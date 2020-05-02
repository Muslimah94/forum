package mplx

import (
	"fmt"
	"net/http"
	"strconv"

	dbase "../dbase"
	handlers "../handlers"
)

// Multiplexer ...
func Multiplexer(db *dbase.DataBase, w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
	p := r.URL.Path
	m := r.Method

	//----------POST---------------------------------------
	if p == "/api/post" {
		if m == "GET" {
			handlers.GetAllPosts(db, w, r)
		} else if m == "POST" {
			handlers.NewPost(db, w, r)
		}
	}

	//-----------COMMENTS----------------------------------
	if len(p) > 13 && p[0:13] == "/api/comment/" {
		postID, err := strconv.Atoi(p[13:])
		if err != nil {
			fmt.Println("Atoi /api/comment:", err.Error())
		}
		if m == "GET" {
			handlers.GetCommentsByPostID(db, w, r, postID)
		}
	} else if m == "POST" && p == "/api/comment" {
		handlers.NewComment(db, w, r)
	}

	//-----------RACTIONS----------------------------------
	if p == "/api/reaction" && m == "POST" {
		handlers.NewReaction(db, w, r)
	}

	//-----------Categories----------------------------------
	if p == "/api/categories" && m == "GET" {
		handlers.GetCategories(db, w, r)
	}
}

func Middleware(db *dbase.DataBase, w http.ResponseWriter, r *http.Request) {

}

// } else if len(p) > 22 && p[0:22] == "/api/post/categoryid/" && m == "GET" {
// 	// categoryID, err := strconv.Atoi(p[22:])
// 	// if err != nil {
// 	// 	fmt.Println("USER Atoi ERROR:", err.Error())
// 	// }
// 	//dbase.GetPostsByCategoryID(db, w, r, categoryID)
// } else if len(p) > 10 && p[0:10] == "/api/post/" {
// 	//postID, err := strconv.Atoi(p[10:])
// 	// if err != nil {
// 	// 	fmt.Println("USER Atoi ERROR:", err.Error())
// 	// }

//else if m == "PUT" {
// 		//dbase.EditPostByID(db, w, r, postID)
// 	} else if m == "DELETE" {
// 		//dbase.DeletePostByID(db, w, r, postID)
// 	}
// }
// //----------COMMENT------------------------------------
// if p == "/api/comment" {
// 	if m == "GET" {
// 		dbase.GetAllComments(db, w, r)
// 	} else if m == "POST" {
// 		dbase.AddNewComment(db, w, r)
// 	}
// } else if len(p) > 13 && p[0:13] == "/api/comment/" {
// 	commentID, err := strconv.Atoi(p[13:])
// 	if err != nil {
// 		fmt.Println("Comment Atoi Error:", err.Error())
// 	}
// 	if m == "GET" {
// 		dbase.GetCommentByID(db, w, r, commentID)
// 	} else if m == "PUT" {
// 		dbase.EditCommentByID(db, w, r, commentID)
// 	} else if m == "DELETE" {
// 		dbase.DeleteCommentByID(db, w, r, commentID)
// 	}
// }

// if p == "/api/category" && m == "POST" {
// 	dbase.AddNewCategory(db, w, r)
// }

// if p == "/api/postscategory" && m == "POST" {
// 	dbase.AddPostCategories(db, w, r)
// }

// if p == "/api/like" {
// 	// 	//DBinteraction.AddReaction(db, w, r)
// 	// } else if p[0:14] == "/api/reaction/" && len(p) > 14 {
// 	// 	if p[0:19] == "/api/reaction/post/" {
// 	// 		postID, err := strconv.Atoi(p[19:])
// 	// 		if err != nil {
// 	// 			fmt.Println("USER Atoi ERROR:", err.Error())
// 	// 		}
// 	// 		if m == "GET" {
// 	// 			fmt.Println("postID:", postID)
// 	// 			dbase.GetAllReactionsToPost(db, w, r, postID)
// 	// 		}
// 	// 	} else if p[0:22] == "/api/reaction/comment/" {

// 	// 	} else if p[0:14] == "/api/reaction/" {
// 	// 	}
// 	// }

// }

// //----------USER---------------------------------------
// if p == "/api/user" {
// 	if m == "GET" {
// 		dbase.GetAllUsers(db, w, r)
// 	} else if m == "POST" {
// 		dbase.AddNewUser(db, w, r)
// 	}
// } else if len(p) > 17 && p[0:17] == "/api/user/roleid/" && m == "GET" {
// 	roleID, err := strconv.Atoi(p[17:])
// 	if err != nil {
// 		fmt.Println("USER Atoi ERROR:", err.Error())
// 	}
// 	dbase.GetUsersByRoleID(db, w, r, roleID)
// } else if len(p) > 10 && p[0:10] == "/api/user/" {
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
// }
// //----------ROLE---------------------------------------
// if p == "/api/role" {
// 	if m == "GET" {
// 		dbase.GetAllRoles(db, w, r)
// 	} else if m == "POST" {
// 		dbase.AddNewRole(db, w, r)
// 	}
// } else if len(p) > 10 && p[0:10] == "/api/role/" {
// 	roleID, err := strconv.Atoi(p[10:])
// 	if err != nil {
// 		fmt.Println("USER Atoi ERROR:", err.Error())
// 	}
// 	if m == "PUT" {
// 		dbase.EditRoleByID(db, w, r, roleID)
// 	} else if m == "DELETE" {
// 		dbase.DeleteRoleByID(db, w, r, roleID)
// 	}
// }
