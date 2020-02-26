# Docker Applications Manager
Docker Applications Manager (DAM) предоставляет пользователю управлять docker образами в системе,
как apt управляет deb пакетами в Ubuntu.   
Это поиск приложения в репозитории, его создание, установка, группировка приложений в продукты,
чистка системы от старых приложений. 

ВАЖНО!!!
На текущий момент реализована только одна основная функция - поиск образов докера и их версий 
в официальном и локальных docker registry (branch 0.0.x).
И сама работа с репозиториями docker.

## Требования к системе
### Сборка
- Go v1.13.4  или GNU Make 4.1
### Использование
- Дистрибутив Linux (Ubuntu 18.04.4 LTS)
- Docker Engine API version v1.40

## Сборка проекта

Изначально проект разрабатывался на версии Go v1.13.4.

Для сборки без Go команда `make build` создаст исполняемый файл `dam`.

## Пример работы

Из "из коробки" доступен поиск в оффициальном Docker Hub.
```
test@home-pc ~/go/src/dam $ ./dam search golang
accelbyte/golang-builder:1.10.3-alpine3.8, 1.11.1, 1.11.2, 1.11.5-alpine3.9, 1.11.5, 1.12.0-alpine3.9, 1.12.0, 1.13.0, latest
aksentyev/golang-glide:latest
alexkappa/golang-libgit2:latest
amd64/golang:1.0.0-alpine, 1.0.0-alpine3.5, ......и т.д.
```

Но можно настроить `dam` для своего локального registry:
```
test@home-pc ~/go/src/dam $ ./dam addrepo --name local --server localhost:5000 --default
```

Тогда поиск образов проходит в локальном registry:
```
test@home-pc ~/go/src/dam $ ./dam search ubunt
ubuntu:16.04
```

Чтобы переключиться обратно на официальный репозиторий, необходимо сделать его используемым по умолчанию:
```
test@home-pc ~/go/src/dam $ ./dam 1 official --default
```

Более подробно по командам см. хелп:
```
test@home-pc ~/go/src/dam $ ./dam help
/--/

Usage:
  dam [command]

Available Commands:
  addrepo     Add a registry to the system.
  help        Help about any command
  listrepos   List all defined repositories.
  modifyrepo  Modify properties of repositories specified.
  removerepo  Remove registry specified by name or number.
  search      Search application in remote registry.

Flags:
  -h, --help   help for dam

Use "dam [command] --help" for more information about a command.
```

## Подробнее для разработки
См. [DEVELOPMENT_ru.md](DEVELOPMENT_ru.md)

## License
DAM is under the Apache 2.0 license. See the LICENSE file for details.