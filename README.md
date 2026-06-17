# gophkeeper

![Go](https://img.shields.io/badge/Go-1.24.2-00ADD8?logo=go)
![PostgreSQL](https://img.shields.io/badge/PostgreSQL-16.8-336791?logo=postgresql)
![gRPC](https://img.shields.io/badge/gRPC-1.73.0-00ACD7?logo=gRPC)
![JWT](https://img.shields.io/badge/JWT-4.5.2-000000?logo=JSON%20web%20tokens)
![Docker](https://img.shields.io/badge/Docker-26.1.0-2496ED?logo=docker)

**gophkeeper** — это клиент-серверное приложение для безопасного хранения конфиденциальной информации, включая пароли, банковские карточки и бинарные данные.

## 📋 Особенности

- **🔐 Защита паролей**: Хранение логинов и паролей с шифрованием
- **💳 Управление картами**: Безопасное хранение данных банковских карт (номер, срок действия, CVV)
- **📂 Бинарные данные**: Хранение любых бинарных файлов с шифрованием
- **🔑 Аутентификация**: JWT-токены для безопасной аутентификации
- **🔒 Шифрование**: AES-GCM для конфиденциальных данных, bcrypt для паролей
- **📡 gRPC**: Высокопроизводительные RPC вызовы между клиентом и сервером
- ** PostgreSQL**: Надежное хранение данных в реляционной СУБД

## 🏗️ Архитектура

```
┌─────────────────┐          gRPC          ┌─────────────────┐
│                 │  ──────────────────▶     │                 │
│   Client CLI    │                          │   Server API    │
│                 │  ◀──────────────────     │                 │
└─────────────────┘                          └─────────────────┘
                                                      │
                                                      │
                                                      ▼
                                              ┌───────────────┐
                                              │  PostgreSQL   │
                                              │   Database    │
                                              └───────────────┘
```

## 🛠️ Стек технологий

### Backend
- **Go 1.24.2** — основной язык программирования
- **gRPC** — коммуникация между клиентом и сервером
- **Protocol Buffers** — определение gRPC API
- **PostgreSQL 16.8** — реляционная база данных
- **bcrypt** — хеширование паролей
- **AES-GCM** — симметричное шифрование данных

### Инструменты
- **Docker & Docker Compose** — контейнеризация
- **Cobra** — создание CLI приложения
- **Zap** — логирование
- **JWT** — токеновая аутентификация

## 🚀 Запуск проекта

### 1. Клонирование репозитория

```bash
git clone <repository-url>
cd gophkeeper
```

### 2. Запуск с помощью Docker Compose

```bash
# Запуск всех сервисов (PostgreSQL)
docker-compose up -d

# Проверка статуса контейнера
docker-compose ps
```

### 3. Запуск сервера

```bash
# Переход в директорию сервера
cd cmd/server

# Запуск сервера (по умолчанию на localhost:5050)
go run main.go
```

**Переменные окружения для сервера:**
- `DATABASE_DSN` — строка подключения к PostgreSQL (по умолчанию: `postgresql://postgres:postgres@localhost:5432/gophkeeper`)
- `DATABASE_TYPE` — тип базы данных (`postgres`)
- `JWT_SECRET` — секретный ключ для JWT токенов
- `JWT_EXPIRATION` — срок действия токена в часах (по умолчанию: 3)
- `CRYPTO_SECRET` — секретный ключ для AES шифрования (по умолчанию: `3a7d4e1f9c02b58e7d9a2f3e8b01c9d7`)
- `GRPC_SERVER_ADDRESS` — адрес gRPC сервера (по умолчанию: `127.0.0.1:5050`)

### 4. Запуск клиента

```bash
# Переход в директорию клиента
cd cmd/client

# Запуск клиента
go run main.go
```

**Использование CLI:**

```bash
# Регистрация пользователя
gothkeeper user register --login <login> --password <password>

# Вход в систему
gothkeeper user login --login <login> --password <password>

# Добавление пароля
gothkeeper password add --title <title> --login <login> --password <password>

# Получение пароля
gothkeeper password get --title <title>

# Обновление пароля
gothkeeper password update --title <title> --login <login> --password <password>

# Удаление пароля
gothkeeper password remove --title <title>

# Добавление банковской карточки
gothkeeper card add --title <title> --bank <bank> --number <number> --dataEnd <date> --secretCode <cvv>

# Получение карточки
gothkeeper card get --title <title>

# Добавление бинарных данных
gothkeeper binary add --title <title> --data <file>

# Получение бинарных данных
gothkeeper binary get --title <title>
```

### Компиляция бинарников

```bash
# Сборка сервера
go build -o gophkeeper-server ./cmd/server

# Сборка клиента
go build -o gophkeeper-client ./cmd/client
```

### Генерация gRPC кода

```bash
# Перегенерация протобуферов
./gen_proto.sh
```

## 📁 Структура проекта

```
gophkeeper/
├── cmd/                          # Точки входа в приложение
│   ├── client/                   # Клиентское приложение (CLI)
│   └── server/                   # Серверное приложение
├── internal/                     # Внутренние пакеты
│   ├── client/                   # Клиентская логика
│   │   ├── app/                  # gRPC клиент
│   │   ├── cli/                  # CLI команды
│   │   └── config/               # Конфигурация клиента
│   ├── logger/                   # Логирование
│   └── server/                   # Серверная логика
│       ├── adapters/             # Адаптеры (DB)
│       │   └── db/psql/          # PostgreSQL реализация
│       ├── app/                  # GPRC сервер и обработчики
│       ├── auth/                 # Аутентификация (JWT)
│       ├── crypto/               # Криптография (AES, bcrypt)
│       ├── interfaces/           # Интерфейсы
│       ├── models/               # Модели данных
│       └── services/             # Бизнес-логика
├── proto/                        # Protocol Buffers файлы
├── docker-compose.yml            # Конфигурация Docker
├── gen_proto.sh                  # Скрипт генерации gRPC кода
└── go.mod                        # Зависимости Go
```

## 🔐 Безопасность

- Пароли хешируются с помощью bcrypt (cost=10)
- Конфиденциальные данные шифруются AES-GCM
- JWT токены имеют ограниченное время жизни
- Все данные передаются по защищенному каналу gRPC

## 📝 Логирование

Приложение использует структурированное логирование через Zap. Логи выводятся в консоль с цветовой разметкой и детальной информацией о событиях.
