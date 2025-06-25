# Сервис отслеживания I/O задач

Высокопроизводительный Go сервис для отслеживания и мониторинга I/O задач через RESTful HTTP API. Создан с учетом масштабируемости и расширяемости, этот сервис предоставляет информацию о ваших I/O операциях в режиме реального времени с минимальными накладными расходами.

## 🚀 Возможности

- **Мониторинг I/O задач в реальном времени** - Отслеживание запросов к базам данных, файловых операций, сетевых запросов и многого другого
- **RESTful HTTP API** - Чистые, интуитивно понятные эндпоинты для управления задачами и мониторинга
- **Расширяемая архитектура** - Легкая замена баз данных, веб-фреймворков и других зависимостей
- **Поддержка Docker** - Готовая к продакшену контейнеризация с многоэтапной сборкой

## 📋 API Эндпоинты

```
POST /create - Создать новую задачу
GET  /view   - Получить детали конкретной задачи
POST /update - Обновить статус задачи
POST /delete - Удалить задачу из отслеживания
```

## 🐳 Использование Docker

### Сборка Docker образа

```bash
# Собрать образ
docker build -t io-task-tracker .
```

### Запуск контейнера

```bash
# Запуск с конфигурацией по умолчанию
docker run -p 8080:8080 io-task-tracker
```

## 🏗️ Архитектура и расширяемость

Сервис построен с модульной архитектурой, поддерживающей легкое расширение и замену компонентов:

### Слой базы данных

```go
type TaskDB interface {
    Create(name, desc string) (uint, error)
    View(taskID uint) (Task, error)
    UpdateStatus(taskID uint, status string) (Task, error)
    UpdateResult(taskID uint, result string) (Task, error)
    Delete(taskID uint) error
}
```

**Поддерживаемые реализации:**
- In-memory (по умолчанию)
- PostgreSQL 
- MySQL
- MongoDB
- Redis

### Слой веб-фреймворка

```go
type Server interface {
    http.Handler
    GET(pattern string, handler HandlerFunc)
    POST(pattern string, handler HandlerFunc)
}
```

**Поддерживаемые фреймворки:**
- Standard library (по умолчанию)
- Gin 
- Gorilla Mux
- и другие...

> **ПРИМЕЧАНИЕ:** Модульный дизайн позволяет вам менять реализации без изменения бизнес-логики. Просто реализуйте необходимые интерфейсы и внедрите ваши предпочтительные зависимости.

### Тестирование API

Для комплексного тестирования API я рекомендую использовать [ghostman](https://github.com/bigelle/ghostman) - инструмент, созданный мной (кстати).

```bash
# Примеры тестов API
# Пока нет бинарника :(
# Из директории git репозитория ghostman:
go run main.go localhost:8080/create -M POST --print-out --data '{"name": "foo", "description": "literally foo"}'
```

## 🚀 Быстрый старт

1. **Клонируйте репозиторий:**
   ```bash
   git clone https://github.com/bigelle/taskservice
   cd taskservice
   ```

2. **Запустите с Docker:**
   ```bash
   docker run -p 8080:8080 io-task-tracker
   ```

3. **Протестируйте API:**
   ```bash
   # Используя curl (или можете использовать любой другой инструмент на ваш выбор *подмигивание*)
   curl http://localhost:8080/create --data '{"name": "foo", "description": "literally foo"}'
   ```

## 📝 Разработка

### Требования

- Go 1.24+
- Docker
