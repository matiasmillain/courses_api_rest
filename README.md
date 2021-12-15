# Course API REST

Levantar un docker con MySQL:

- docker run -p 3306:3306 --name courses -e MYSQL_ROOT_PASSWORD=p4ssCour -e MYSQL_DATABASE=api_rest_courses -d mysql:5.7.25

- go run server.go

