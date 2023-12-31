# Сервис Загрузки Статей в IPFS

## Введение

Сервис Загрузки Статей в IPFS предоставляет возможность добавлять статьи в базу данных MongoDB и загружать их содержание на IPFS (InterPlanetary File System). Каждая статья хранится в базе данных и имеет уникальный CID (Content Identifier), который представляет собой хэш-сумму содержания статьи в сети IPFS.

## Запуск Сервиса

Для запуска сервиса, выполните следующие шаги:

1. **Установка Docker**: Убедитесь, что у вас установлен Docker на вашем компьютере.

2. **Отредактируйте переменные окружения**:

   ```env
   MONGO_URI=mongodb://root:qwerty@mongodb:27017
   MONGO_DB_NAME=papers
   MONGO_COLL_NAME=articles
   HTTP_SERVER_PORT=8080
   PINATA_API_KEY=your_pinata_api_key
   PINATA_API_SECRET=your_pinata_api_secret
   ```

   Замените `your_pinata_api_key` и `your_pinata_api_secret` на ваши собственные ключи Pinata API.

3. **Запуск Контейнеров**: Выполните команду `docker-compose up` в терминале, находясь в корневой директории проекта.

4. **Проверка Сервиса**: После успешного запуска контейнеров, сервис будет доступен по адресу `http://localhost:8080`.

## Использование Сервиса

### Добавление Статьи

Для добавления новой статьи в базу данных и её загрузки на IPFS, отправьте POST-запрос на следующий адрес:

```
http://localhost:8080/add-paper
```

Запрос должен иметь тело в формате JSON с полем `content`, содержащим текст статьи. Пример запроса с использованием `curl`:

```bash
curl -X POST -H "Content-Type: application/json" -d '{"_id":"5", "content":"Текст новой статьи"}' http://localhost:8080/add-paper
```

Если операция успешно выполнена, вы получите ответ с кодом состояния 200 OK и сообщением, содержащим ID.

### Получение CID Статьи

Чтобы получить CID статьи по её уникальному идентификатору (ID), отправьте GET-запрос на следующий адрес:

```
http://localhost:8080/paper/{id}
```

Замените `{id}` на фактический ID статьи, который вы хотите найти. Если статья с указанным ID найдена, вы получите CID статьи на IPFS в ответе.

## Замена Переменных

Перед развертыванием сервиса в продакшн, необходимо заменить следующие переменные:

- `PINATA_API_KEY`: Замените на ваш собственный API ключ от Pinata.

- `PINATA_API_SECRET`: Замените на ваш собственный секретный API ключ от Pinata.

## Работа Сервиса

Сервис выполняет следующие шаги при запросах:

1. Получает запрос на добавление статьи в формате JSON с текстом статьи.

2. Добавляет статью в базу данных MongoDB.

3. Загружает текст статьи на IPFS, получая CID.

4. Возвращает CID в ответ на запрос.

5. При запросе по ID статьи, сервис ищет статью в базе данных MongoDB и возвращает её CID на IPFS.
---
