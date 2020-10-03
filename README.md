# dnd-golang

## Running Locaaly

1. Update postgres values in configuration/config.go (Host, Port, User, Dbname)
2. Set DB_PASSWORD as an environment variable for Postgres
3. You may need to get the Dice package. ```go get github.com/smwest87/dnd_dice```
4. Create the schema and tables using SQL statements in sql/setup.sql. Currently this is manual.
5. `go run main.go`


## Post Request Body

New Character endpoint : `/character/new`

Body :  ```{
	"name": "Kongul",
	"class": "Barbarian"
}```