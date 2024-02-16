## Base API server with simple configs


#### Work on this subject will continue


#### Some basic info

1. Библиотеки для работы с БД: 
database/sql, 
sqlx, 
gosql. 


2. Инициализация хранилища: 
storage/storage.go, 
go get -u github.com/lib/pq - для работы с постгрес, 
SQLDrivers - на github. 

3. Первичная миграция: 
golang/migrate 

3.1. Создание миграционного репозитория: \
В данном репозитории будут находиться up/down пары sql миграционных запросов к бд. \
```migrate create -ext sql dir migrations UsersCreationMigration```

3.2 Применить миграцию: \
```migrate -path migrations -database "postgres://localhost:5434/restapi?sslmode=disable&user=postgres&password=admin" up```

