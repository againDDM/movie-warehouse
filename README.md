# movie-warehouse v-0.1
REST API go example

**GET** `/films`   
`curl --user "admin:admin" localhost:8000/films | jq`    

**GET** `/films/{id}`  
`curl --user "admin:admin" localhost:8000/films/5`   

 **POST** `/films`  
`curl localhost:8000/films -i -X POST --user "admin:admin" --data '{"name": "movie title", "description": "Movie description"}'`  

 **PUT** `/films/{id}`  
`curl localhost:8000/films/7 -i -X PUT --user "admin:admin" --data '{"name": "new title", "description": "New description"}'`   
`curl localhost:8000/films/9 -i -X PUT --user "admin:admin" --data '{"name": "new title"}'`   
`curl localhost:8000/films/8 -i -X PUT --user "admin:admin" --data '{"description": "New description"}'`   

**DELETE** `/films/{id}`  
`curl -X DELETE -i --user "admin:admin" localhost:8000/films/7`
