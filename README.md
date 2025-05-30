# Quote API

Простое REST API приложение для работы с цитатами. Позволяет получать случайные цитаты, добавлять новые и  удалять существующие

##  Установка и запуск

1. Клонируйте репозиторий:
   ```bash
   git clone https://github.com/datslim/quote-api.git
   cd quote-api
   ```
2. Запустите сервер: 
    ```bash 
    go run cmd/main.go
    ```


Сервер будет доступен по адресу: http://localhost:8080/quotes

## Использование

### Запустить unit-тесты

Для запуска unit-тестов используйте команду: 
```bash
    go test -v ./unit-tests/
```

### Проверить отправку запросов
Для проверки отправки запросов предусмотрена схема для `Postman`.
Для ее использования, после запуска сервера, нажмите Import внутри приложения Postman и выберите `postman_collection.json`
После этого вы можете попробовать отправить тестовые запросы.  
Если же в вашем случае этого сделать не удается, вы всегда можете воспользоваться `curl` с помощью приведенных ниже запросов.  

### Добавить новую цитату
```bash 
    curl -X POST http://localhost:8080/quotes \ -H "Content-Type: application/json" \ -d '{"author":"Confucius", "quote":"Life is simple, but we insist on making it complicated."}'
```
### Получить все цитаты
```bash
    curl http://localhost:8080/quotes 
```

### Получить случайную цитату
```bash
    curl http://localhost:8080/quotes/random 
```

### Получить все цитаты автора
```bash
    curl http://localhost:8080/quotes?author=Confucius
```
### Удалить цитату по ID
```bash
    curl -X DELETE http://localhost:8080/quotes/1
```


## Структура проекта
```
    cmd/main.go                    — запуск сервера
    internal/handlers/quotes.go    — HTTP-хендлеры
    internal/storage/memory.go     — in-memory хранилище
    internal/controller/routes.go  — описание всех endpoint'ов
    internal/model/quote.go        — модель Quote
    unit-tests/quotes_test.go      — unit-тесты
    postman_collection.json        — json-схема для Postman
```