# GREEN-API
Проект запущен и доступен по адресу [http://109.172.114.84](http://109.172.114.84).

## Backend
Я решил начать с инкапсуляции методов для запросов к нужному API в небольшую библиотеку, расположенную в файле `/backend/greenapi/greenapi.go`. В этой библиотеке используются структуры, которые хранят URL для запроса, идентификатор инстанса и токен.

Далее был написан основной файл `/backend/main.go`, который реализует логику API сервера. Этот файл отвечает за обработку запросов и вызывает методы библиотеки, передавая необходимые параметры для работы с API.

## Frontend
Фронтенд находится в папке `/frontend/app/index.html` и представляет собой стандартную HTML-страницу с прописанными в ней стилями и JS. Страница минимально валидирует данные и позволяет взаимодействовать с ранее созданным API. Статика обслуживается через Nginx, конфигурация для которого хранится в папке с фронтендом. Так же Nginx используется для проксирования запросов

## Deploy
Для удобства развертывания каждого компонента предусмотрен `Dockerfile`, который описывает, как собрать и запустить сервисы в контейнерах. Вся система собирается и запускается с помощью Docker Compose

## Тестовые данные
- **IdInstance**: `1103157166`
- **ApiTokenInstance**: `cecd74e7d35849efa82c5a46f1fc543d618525b63cb84abda2`
- **ChatId**: `79687019003`
- **UrlFile**: [Image URL](https://sun9-37.userapi.com/impg/rWmGseJlnzKZjgbL9qNstDQyYpw7lo5S80Scgg/saRRNI0Crho.jpg?size=1170x1143&quality=95&sign=1b5e80e430c6338c90c9a4e06d2775b7&type=album)

