# forum\

/api/register [POST] user registration to forum
/api/login [POST] //nickname, password ()
/api/addpost [POST] create new post
/api/posts?liked=0&created=0 [GET] get all posts
/api/posts?liked=1&created=0 [GET] get liked posts
/api/posts?liked=0&created=1 [GET] get created posts
/api/post?id= [GET] get post by post id
/api/comment [POST] write a new comment to post by post id
/api/comments?post_id= [GET] get all comments to post by post id
/api/categories [GET] get names of categories
/api/reaction [POST] add like/dislike to DB (like type=1, dislike type=0)



Will be implemented for ADVANCED FEATURES (optional task):
/api/profile?id= [GET]
/api/profile?id= [PUT]
/api/post?id= [PUT] edit post
/api/comment?id= [PUT] edit comment
/api/post?id= [DELETE] delete post
/api/comment?id= [DELETE] delete comment