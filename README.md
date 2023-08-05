# Task Manager

## Install Libraries

Server side libraries

```shell
$ go mod tidy
```

Client side libraries

```shell
$ cd frontend
$ npm i
```

## Prepare database

Launch

```shell
$ docker compose up
```

Migration

```shell
$ GO_ENV=dev go run migrate/migrate.go
```

## Launch

### Server

develop

```shell
$ GO_ENV=dev go run main.go
```

### Client

develop

```shell
$ cd frontend
$ npm run dev
```

## Environment variables to set

```env
PORT={server port}
POSTGRES_USER={db user}
POSTGRES_PW={db port}
POSTGRES_DB={db name}
POSTGRES_PORT={db port}
POSTGRES_HOST={db host}
SECRET={secret key}
GO_ENV={go env}
API_DOMAIN={api domain}
FE_URL={front end url}
```
