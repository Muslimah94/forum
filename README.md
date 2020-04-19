# forum

/api/user:
GET: get all users
POST: add new user
----------------------------
/api/user/id:
GET: get user by ID
PUT: edit user by ID
DELETE: delete user by ID
----------------------------
/api/user/roleid/id:
GET: get users by role ID
--------------------------------------------------------
/api/role:
GET: get all roles
POST: add new role
----------------------------
/api/role/id:
PUT: edit role by ID
DELETE: delete role by ID
--------------------------------------------------------
/api/post:
GET: get all posts
POST: add new post
----------------------------
/api/post/id:
GET: get post by ID
PUT: edit post by ID
DELETE: delete post by ID 
--------------------------------------------------------
/api/comment/postID/id:
GET: get all comments to post
POST: add new comment to post
----------------------------
/api/comment/commentID/id:
GET: get comment by ID
PUT: edit comment by ID
DELETE: delete comment by ID 
--------------------------------------------------------
/api/reaction:
POST: add new reaction to post/comment | depends on fields of json
/api/reaction/id:
DELETE: delete reaction by ID
----------------------------
/api/reaction/post/id:
GET: get all reactions to Post
----------------------------
/api/reaction/comment/id:
GET: get all reactions to comment
--------------------------------------------------------
/api/login
{get a row where email  = ?
if there's no email return eroor
check password
if true create session
write to cookie's value uuid
MaxAge 3600
