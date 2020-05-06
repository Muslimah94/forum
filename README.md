# forum\

/api/posts [POST] create new post
/api/posts [GET] get all posts
/api/post?id= [GET] get post by post id
/api/comment [POST] write a new comment to post by post id
/api/comments?post_id= [GET] get all comments to post by post id
/api/categories [GET] get names of categories
/api/reaction [POST] add like/dislike to DB (if like type=1, else type=0)
/api/register [POST] user registration to forum

sessii() get session by id, create session

{get a row where email  = ?
if there's no email return eroor
check password
if true create session
write to cookie's value uuid
MaxAge 3600
