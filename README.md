# user_service

Сервис пользователя для приложения LinkUp

Install go-swagger:
* `go install github.com/swaggo/swag/cmd/swag@latest`

Generate swagger:
* `swag init`

Start migration:
* `go run ./api/base/db/migrations/start_migrations.go`

Generate business constants from table business_constants
* `go run ./api/base/db/migrations/generate_business_constants/generate.go`