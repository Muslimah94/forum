# forum\

/api/register [POST] user registration to forum
/api/posts [POST] create new post
/api/posts [GET] get all posts
/api/post?id= [GET] get post by post id
/api/comment [POST] write a new comment to post by post id
/api/comments?post_id= [GET] get all comments to post by post id
/api/categories [GET] get names of categories
/api/reaction [POST] add like/dislike to DB (like type=1, dislike type=0)
/api/reaction [GET] identifying reaction of user on post/comment for coloring icon
/api/login [POST] //nickname, password ()


Will be implemented for ADVANCED FEATURES (optional task):
/api/profile?id= [GET]
/api/profile?id= [PUT]
