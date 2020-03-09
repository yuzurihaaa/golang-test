# Golang-test

1. A project for me to learn Golang
2. This codebase is not a good place for reference because I just started learning the language

## Dev environment
1. Postgres database (more in [docker-compose](docker-compose.yml))
2. Docker
3. A correct golang setup
4. ORM is using [sqlboiler](https://github.com/volatiletech/sqlboiler)
5. Config for sqlboiler is in [toml](sqlboiler.toml) but for docker is in [.env](database.env)
6. Model is auto created, run `go generate` (in `main.go`)

## Running the project
1. Run database with `docker-compose up`
2. Run the project with `go run main.go`
3. Start using the API.
