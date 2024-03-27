# golang-projects
A collection of small Go projects all in one repo

## http-server
A simple blog post platform with CRUD operations for blog posts.
- GET /v1/blogposts will retrieve a list of all blogposts as json.
- POST /v1/blogposts will create a new blogpost and append it to the list
- GET /v1/blogposts/{id} will get a blogpost from the ID
- PUT /v1/blogposts/{id} will update a blogpost with the given ID
- DELETE /v1/blogposts/{id} will delete a blogpost with the given ID
