# Quote API

Простое REST API для работы с цитатами. Позволяет получать случайные цитаты, добавлять новые и  удалять существующие

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