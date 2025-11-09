## Database
### Create migrations
```bash
atlas migrate diff test --env gorm
```

### Migrate
```bash
atlas schema apply --env gorm -u "postgresql://<user>:<password>@<host>:<port>/<database_name>?search_path=public<&sslmode=disable>"
```

## Swagger
### Build docs from code
```bash
swag init
```