# NATS Streaming with Go

`ДЛЯ РАБОТЫ НЕОБХОДИМ ЗАПУЩЕННЫЙ NATS STREAMING SERVER`

Инструкция:

```bash
git clone github.com/odysseymorphey/httpServer.git
cd httpServer
make
make run 
```

Сводка по Makefile:
```text
make - сбилдит проект
make run - запустит сервер
make migrate - создаст таблицу в бд
make pub - отправит сообщение в NATS Streaming(если запущен сервер, то данные запишутся в бд)
```
