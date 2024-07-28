# Auth API

### Fix golang interpreter path
```
export PATH=$PATH:/usr/local/go/bin
```

### Libraries
- [Request Handler](https://echo.labstack.com/docs/quick-start)
  - go get github.com/labstack/echo/v4
- [Gorm](https://gorm.io/)
  - go get -u gorm.io/gorm
- [Gorm PostgreSQL](https://github.com/go-gorm/postgres)
	- go get github.com/go-gorm/postgres
- [Redis](https://github.com/redis/go-redis)
  - go get github.com/redis/go-redis/v9

### Database configuration
```sh
cp -R data/database.sql dataPgsql/db/database.sql

./dockermg.sh postgres bash

psql -U postgres -d gopos -f /var/lib/postgresql/data/database.sql
```

### Redis
```sh
./dockermg.sh redis bash

redis-cli -a yourpassword
```

### Requests
```sh
curl --request POST localhost:8082/auth/login -H "Content-Type: application/json" -d '{"email":"admin@example.com", "password":"user_pass"}'

curl --request POST localhost:8082/auth/logout -H "Content-Type: application/json" -d '{"email":"admin@example.com"}'

curl --request GET localhost:8082/users\?page\=1\&limit\=10 -H "Authorization: Bearer $JWT"

curl --request POST localhost:8082/users -H "Authorization: Bearer $JWT" \
  -H "Content-Type: application/json" \
  -d '{"name": "Client1", "email":"client1@example.com", "password":"user_pass", "birthday":"1980-01-01"}'
```
