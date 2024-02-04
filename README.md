## Scan repo!!!!

### DB Migration

To add migration:
```shell
docker run --rm -it -v $PWD/migration:/migration migrate/migrate:v4.16.2 create -ext sql -dir /migration -seq insert_data
```

Migrate up:
```shell
docker run --rm -it -v $PWD/migration:/migration migrate/migrate:v4.16.2 -path=/migration/ -database "postgres://USER:PASS@10.0.0.16:5432/DB?sslmode=disable" up [N]
```

Migrate down:
```shell
docker run --rm -it -v $PWD/migration:/migration migrate/migrate:v4.16.2 -path=/migration/ -database "postgres://USER:PASS@10.0.0.16:5432/DB?sslmode=disable" down [N]
```