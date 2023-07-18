# golang_grpc_testovoe

## Установка

1. Установка protoc - https://grpc.io/docs/protoc-installation/
2. Установка protoc golang - https://grpc.io/docs/languages/go/quickstart/
3. Установка protoc dart - https://grpc.io/docs/languages/dart/quickstart/
4. Установка grpc-ui - https://github.com/fullstorydev/grpcui

## Задача
```
Написать асинхронный сервер и клиент, отдающий информацию о директории по заданному в запросе пути (список файлов/директорий и их размер) На сервере необходимо поддержать временный кэш (в оперативной памяти) ответов по недавно запрошенным путям.
Испорльзовать GRPC для общения между сервером и клиентом, реализовать работу с файловой системой встроенными средствами языка.
```
# Генерация proto
```shell
$ make build
```
# Настроки перед запуском сервера 
1. скопировать ``.env.example`` в ``.env``
2. для запуска redis  
```shell
docker-compose up -d
```
3. в файле client/client.go изменить путь до нужной папки либо восползоваться клиентом
## UI клиент если хотите протестировать через браузер
```shell
grpcui -open-browser -plaintext -import-path ./proto $(find ./proto -iname "*.proto" -printf ' -proto %h/%f') -port 8084 127.0.0.1:8080

```
