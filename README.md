# MessageProcessing
Микросервис обработки сообщений

## Содержание
- [TalkBoard](#TalkBoard)
- [содержание](#содержание)
- [Архитектура](#архитектура)
- [Локальный запуск](#локальный-запуск)

## Архитектура

Проект основан на "чистой" архитектуре и разработан на языке Go. Реализованная архитектура включает три основных слоя: handler, service и repository. В слое handler обрабатываются HTTP запросы, слой service содержит бизнес-логику, а в repository выполняются операции с базой данных.

Для упрощения развертывания и изоляции окружений используется Docker с docker-compose для сборки и запуска контейнеров. Это позволяет легко воспроизводить среду разработки и обеспечивать консистентность окружений между разработкой и продуктивной средой.

## Локальный запуск
Для развертывания сервиса необходимы [Docker](https://docs.docker.com/engine/install/) и [Docker Compose](https://docs.docker.com/compose/).

**Важно:** Перед запуском убедитесь, что вы установили пароли. Для этого создайте файл .env в корневой директории и добавьте туда:\
`DB_PASSWORD=...`

1. `git clone --recurse-submodules https://github.com/klausfun/MessageProcessing`
2. `cd MessageProcessing`
3. `docker pull postgres` - для скачивания образа postgres
4. `docker-compose up --build`
5. `migrate -path ./schema_db -database 'postgres://postgres:qwerty@localhost:5436/postgres?sslmode=disable' up` (вместо 'qwerty' необходимо указать Ваш пароль от БД)