asfasfasfv0002
##############TTTTTTTTTTTaaaaaa
## Run docker

* `docker-compose build`
* `docker-compose up -d`

## Hot reload
asfasf
use command `air`

### test api Hot reload

asfasf
run `air -c .air_test.toml`

## Run test api

run `go run cmd/test-api.go` in the container.
as
## Migrate

### create new migration

`migrate create -ext sql -dir database/migrations create_users_table`

### run migrate
- up
`migrate -database "mysql://${username}:${password}@tcp${MYSQL_URL}/${DB_NAME}" -path db/migrations up`


- down
`migrate -database "mysql://${username}:${password}@tcp${MYSQL_URL}/${DB_NAME}" -path db/migrations down ${BATCH}`


asf
asfasfaoeu