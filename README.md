# movie-warehouse v-0.1
REST API go example

**GET** `/films`<br/> 
`curl --user "admin:admin" localhost:8000/films | jq` <br/> 

**GET** `/films/{id}`<br/>
`curl -vH "Connection: close" --user "admin:admin" localhost:8000/films/5` <br/>

 **POST** `/films`<br/>
`curl localhost:8000/films \
      -i -X POST --user "admin:admin" \
      --data '{"name": "movie title",
               "description": "Movie description"}'` <br/>

 **PUT** `/films/{id}`<br/>
`curl localhost:8000/films \
      -i -X PUT --user "admin:admin" \
      --data '{"name": "new title",
               "description": "New description"}'` <br/>

**DELETE** `/films/{id}`<br/>
`curl -X DELETE -i --user "admin:admin" localhost:8000/films/4`
