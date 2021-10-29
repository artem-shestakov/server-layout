# Golang HTTP server layout

## Database migrations
> Use [golang-migration](https://github.com/golang-migrate/migrate) tool
* Create migrations in `db/migrations folder`. There are several migrations exist.
* Run `ud`/`down` migration. Use database URL as environment variable for convenience:
```shell
# Create environment variable
$ export POSTGRESQL_URL='postgres://api:api@localhost:5432/api?sslmode=disable'
#  Up
$ migrate -database ${POSTGRESQL_URL} -path db/migrations up 
# Down
$ migrate -database ${POSTGRESQL_URL} -path db/migrations udown
```
