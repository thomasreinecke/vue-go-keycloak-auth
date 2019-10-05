
# go-rest - RESTful service exposing a basic user management

my very first go project, trying to learn the language, packaging structures and best practices



## cmd/sushibar
REST service bound to `::5000` thats providing a menu of sushi rolls on `http://localhost:5000/sushi`

### Installation
```
git clone ...
go run cmd/sushibar/sushibar.go 
```

the **sushibar** service exposes the following REST interface:
* GET `/sushi` returns full list of all sushi rolls on the menu
* GET `/sushi/{id}` returns a specific roll from the menu
* POST `/sushi` adds a roll to the sushi menue (structure is `{"name":"my roll","ingredients":"good stuff"}`
* POST `/sushi/{id}` updates an existing sushi roll on the given id or throws a HTTP:403 if the id is not known
* DELETE `/sushi/{id}` deletes an existing roll on the given id or throws a HTTP:403 if the id is not known

