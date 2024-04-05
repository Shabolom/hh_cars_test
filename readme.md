
# Test_Cars.
___
## ***Подключение и настройка зависимостей.***
**Для автоматической докачки зависимостей:** используем команду
```go
go get ./..
```

**Подключение к db**
> Для подключения в базу происходит по средствам форматирования строки данными из env файла.

*Пример:*
```go
// создание строки подключения
connectionString := fmt.Sprintf("postgresql://%s:%s@%s:%s/%s?sslmode=disable",
    Env.DbUser,
    Env.DbPassword,
    Env.DbHost,
    Env.DbPort,
    Env.DbName,
)
// непостредственно подключение в базу
pool, err := pgxpool.Connect(context.Background(), connectionString)
if err != nil {
return err
}
```
**Инициализация squirrel и добавления плейсхолдеров.**

*Пример:*
```go
sqlBuilder := squirrel.StatementBuilder.PlaceholderFormat(squirrel.Dollar)
if err != nil {
    return err
}
```
**Запуск сервера**
> Для запуска сервера необходимо передать параметры хоста и порта, передача происходит с помощью env файла.

*Пример:*

```go
if err = r.Run(config.Env.Host + ":" + config.Env.Port); err != nil {
		log.WithField("component", "run").Fatal(err)
	}
}
```
**Создание файла миграции**
> Просто удобный способ создать файл миграции с правильныс расширением.
*Пример:*
```go
migrate create -ext sql -dir (dir) -seq fileName
```
**Инициализация миграции**
> Для создания миграции необходимо передать путь до файлов и указать строку подключения к базе данных.

*Пример:*
```go
migrate -source file://file_path -database postgresql://dbUser:dbPassword@dbHost:dbPort/dbName?sslmode=disable
```
___
## Api и его функционал.
___
### Api
```go
	r.GET("/docs/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	r.GET("/api/car", car.Get)
	r.GET("/api/car/:id", car.GetID)
	r.POST("/api/car", car.Post)
	r.DELETE("/api/car/:id", car.Delete)
	r.PUT("/api/car/:id", car.Update)
```

**Функционал**
>Ниже будет разобрана каждая ручка и ее функционал

  + ```go
    // Получение данных с фильтрацией по всем полям и пагинацией, данные передаются в query params.
    r.GET("/api/car", car.Get)
    ```
  + ```go
    // Получение данных получение данных по id машины для дальнейшего заполнения ручки Update для удобства пользователя.
    r.GET("/api/car/:id", car.GetID)
    ```
  + ```go
    // Добавления новых автомобилей в формате принятия массива номеров с последующим обращением на другой сервис и получения данных об этих машинах.
    r.POST("/api/car", car.Post)
    ```
  + ```go
    // Удаления по идентификатору.
    r.DELETE("/api/car/:id", car.Delete)
    ```
  + ```go
    // Изменение одного или нескольких полей по идентификатору.
    r.PUT("/api/car/:id", car.Update)
    ```
___

## Для кореектной работы Swagger.

> В папке cmd в файле main.go необходимо изменить тэе @host на используемый.

*Пример:*
```go
//	@host		localhost:8800
```
____
# Скриншоты структуры db и Swagger.
**Схема db:**
![схема_db](https://github.com/Shabolom/hh_cars_test/blob/finalv1.0/media/bd.png)

**Скриншот Swagger**
![swagger]
