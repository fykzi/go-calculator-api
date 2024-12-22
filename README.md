# Go-calculator-API

**Go-calculator-API** — сервис для решения математических выражений.

---

## Установка и запуск

Для запуска проекта выполните следующие шаги:

1. Склонируйте репозиторий:

```bash
git clone https://github.com/fykzi/go-calculator-api
```

2. Перейдите в корневую папку проекта:

```bash
cd go-calculator-api/
```
  
3. Установите Golang, если у вас он не установлен. Для этого перейдите по https://go.dev/dl.

4. Возможно могут возникнуть проблемы с зависимостями. В случае этой проблемы, если проект не может запуститься, введите следующую команду.

```bash
go mod tidy
```
5. Запустите API-сервер:

```bash
go run ./cmd/main.go
```

Сервер запустится на `127.0.0.1:8000` по умолчанию.

## Конфигурация проекта

Проект поддерживает конфигурацию. Вы можете изменить host, port, а также log_level.

По умолчанию конфиг проекта находиться в `go-calculator-api/config/config.yaml`.

Вы можете разместь свой конфиг в любом месте, но для работы проекта вам будет необходимо через флаг `--config` передать путь до файла, например:

```bash
go run ./cmd/main.go --config="path/to/config.yaml"
```

Конфиг по умолчанию выглядит следующим образом:

```yaml
host: "127.0.0.1"
port: 8000
log_level: "INFO"
```
### Значения в конфиге.
- host - иными словами это ip-address, на каком вы хотети запустить ваш проект.
- port - числовое значение в диапазоне от 0 до 65536. Будьте осторожны, если порт в вашей системе занят, проект не сможет запуститься.
- log_level - уровень логирования. Принимает следующие значения - `DEBUG`, `INFO`, `WARN` и `ERROR`. Подробнее можно прочитать на `https://pkg.go.dev/log/slog#Level`

---

## Использование API

### Эндпоинт

```
POST /api/v1/calculate
```

### Заголовки

- `Content-Type: application/json`

### Тело запроса

Пример:

```json
{
  "expression": "2 + 2 * 2"
}
```

### Ответы

1. **Успешный запрос**

   **Статус-код:** `200 OK`  
   **Пример ответа:**

   ```json
   {
     "result": "6"
   }
   ```

2. **Ошибка обработки выражения**

   **Статус-код:** `422 Unprocessable Entity`  
   **Пример ответа:**

   ```json
   {
     "error": "Expression is not valid"
   }
   ```

3. **Неподдерживаемый метод**

   **Статус-код:** `405 Method Not Allowed`  
   **Пример ответа:**

   ```json
   {
     "error": "Invalid request method"
   }
   ```

4. **Некорректное тело запроса**

   **Статус-код:** `400 Bad Request`  
   **Пример ответа:**

   ```json
   {
     "error": "Bad request"
   }
   ```
5. **Внутрення ошибка сервера**

   **Статус-код:** `500 Internal Server Error`  
   **Пример ответа:**

   ```json
   {
     "error": "Internal server error"
   }
   ```

---

## Примеры использования

1. **Успешный запрос**:

```bash
curl http://127.0.0.1:8000/api/v1/calculate \
--header 'Content-Type: application/json' \
--data '{
  "expression": "2 + 2 * 2"
}' -X POST
```

Ответ:

```json
{
  "result": "6"
}
```

2. **Ошибка: некорректное выражение**:

```bash
curl http://127.0.0.1:8000/api/v1/calculate \
--header 'Content-Type: application/json' \
--data '{
  "expression": "2 + 2 * 2)"
}' -X POST
```

Ответ:

```json
{
  "error": "Expression is not valid"
}
```

3. **Ошибка: неверный метод**:

```bash
curl http://127.0.0.1:8000/api/v1/calculate \
--header 'Content-Type: application/json' \
--data '{
  "expression": "2 + 2 * 2"
}' -X GET
```

Ответ:

```json
{
  "error": "Invalid request method"
}
```

4. **Ошибка: неверное тело запроса**:

```bash
curl http://127.0.0.1:8000/api/v1/calculate \
--header 'Content-Type: application/json' \
--data '{
  "expression": "2 + 2 * 2"
' -X POST
```

Ответ:

```json
{
  "error": "Bad request"
}
```
