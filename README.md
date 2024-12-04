# Тестовое задание для Ingry.tech

## Затраченное время на выполнение: 3ч 40м

## Для запуска
1) git clone https://github.com/ZiganshinDev/local-deploy.git
2) docker compose up postgres-dev
3) git clone https://gitlab.com/zigan/ingrytech.git
4) make up в корневой директории

## Далее описание тестового задания с пометками

#### Тестовое задание 1: Разработка REST API - выполнено, но данные хранятся в бд (пропустил шаг с сохранением in-memory)

Описание:
Создайте простое REST API для управления библиотекой книг. API должен поддерживать следующие операции:
- Получить список всех книг (GET /books)
- Получить информацию о конкретной книге по ID (GET /books/{id})
- Добавить новую книгу (POST /books)
- Обновить информацию о книге по ID (PUT /books/{id})
- Удалить книгу по ID (DELETE /books/{id})

Требования:
- Использовать фреймворк ECHO.
- Данные книг должны храниться в памяти (можно использовать слайс или мапу).
- Книги должны содержать следующие поля: ID, название, автор, год издания.

---

#### Тестовое задание 2: Использование GORM и работа с базами данных - выполнено (выбрана PostgreSQL)

Описание:
Создайте приложение, использующее GORM для работы с базой данных. Реализуйте ту же функциональность, как в первом задании, но теперь данные должны храниться в базе данных MySQL или PostgreSQL.

Требования:
- Настроить подключение к базе данных с использованием GORM.
- Реализовать миграции для создания таблицы книг.
- Обработать все возможные ошибки, возникающие при взаимодействии с базой данных.
- Предоставить скрипт для создания базы данных и необходимые настройки для её использования.

---

#### Тестовое задание 3: Обработка ошибок и Middleware - выполнено

Описание:
Расширьте API для управления библиотекой книг, добавив middleware для обработки ошибок. Все ошибки должны возвращаться клиенту в формате JSON с соответствующим статусом HTTP.

Требования:
- Создайте middleware для глобальной обработки ошибок.
- Если происходит ошибка, возвращайте ответ с статусом 400 или 500 в зависимости от характера ошибки.
- Все сообщения об ошибках должны содержать понятное описание.

---

#### Тестовое задание 4: Написание тестов - не успел

Описание:
Напишите тесты для вашего REST API, используя библиотеку testing и net/http.

Требования:
- Использовать Go testing package.
- Написать тесты для всех ваших эндпоинтов (GET, POST, PUT, DELETE).
- Убедитесь, что тесты покрывают успешные сценарии и обработку ошибок.
- Добавить команду для запуска тестов в README.

---

#### Тестовое задание 5: Документация и CI/CD - быстро набрасал CI/CD, но не успел добавить генерацию спеки OpenAPI/Swagger

Описание:
Создайте документацию для вашего API с использованием Swagger или аналогичного инструмента.

Требования:
- Оснастить свой проект документацией, которая показывает, как использовать каждый эндпоинт.
- Настроить базовую CI/CD конфигурацию в GitLab для автоматической сборки и тестирования вашего приложения при каждом пуше в репозиторий.