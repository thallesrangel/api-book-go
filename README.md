# API Rest com GO!

## Tutorial

<a href="https://www.youtube.com/watch?v=socM14c9Vfk">
  <img src="https://img.shields.io/badge/-ASSSITA%20AQUI-black?style=for-the-badge&logo=youtube&color=red"></img>
</a>

## Docs API - Insomnia
- docs/Insomnia_books_api.json

## Run Database

```
$ docker-compose up --build
```

## Run Database

```
$ go run main.go
```

## Migrations

- Migrations is auto runned when start the database

## Seeds

- Run seeds is manual, you have to run

```
$ go run database/seeder/main.go
```
<hr>

## Setup - New Project
```
    go mod init github/thallesrangel/api-book-go
    go get github.com/gin-gonic/gin
    go get gorm.io/gorm
    go get gorm.io/driver/postgres
```