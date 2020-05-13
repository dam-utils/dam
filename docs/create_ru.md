# dam create
Нет смысла усложнять логику, утилита должна быть простой.
Лучше добавить больше валидаций и обучать пользователя ворнингами и ошибками.
Чем усложнять бизнес логику, делать гибкость.

Структура каталога проекта:
Dockerfile
meta/
    DESCRIPTION (не обязателен)
    install.exp (не обязателен)
    install
    uninstall.exp (не обязателен)
    uninstall
    ENVIRONMENT (не обязателен)

В файлах .exp обязательно должны быть заменены переменные окружения и они превратятся в файлы без .exp
В install может быть сформировано имя образа, чтобы во внешней системе знать, что устанавливать.
В файлах .exp заменяются все переменные в скобках ${} на переменные из:
- файла ENVIRONMENT
- Dockerfile
- переменных окружения, начинающихся с DAM_PWD="./dam"
Те переменные, что не найдены при замене в файле, выводятся при сборке ворнинги, заменяются на пустое.
(Обычные переменные окружения можно указать просто $VAR)
Можно подумать и реализовать переменные по умолчанию.
При создании ищутся файлы .exp в каталоге meta и заменяются.
Приоритет переменных окружения означает, что при каждой новой проверке, переменные окружения заменяются на более приоритетные из списка приоритетов.

Имя приложения формируется из (по приоритету):
- "unknown"
- имя каталога проекта, за вычетом лишних символов (если не пустое)
- $DAM_APP_NAME

Версия приложения формируется из (по приоритету):
- SNAPSHOT
- $DAM_APP_VERSION

При создании приложения:
- dam проверяет, указан ли проект. Есть ли в нем Dockerfile и meta/, с файлами install (.exp) и uninstall (.exp)
- формирует список переменных окружения
- проверяет, копируется (нет строчек с метой ADD или COPY) ли meta в Dockerfile - иначе ошибка
- проверяет, Есть ли family в Dockerfile
- создает образ приложения с тэгом <имя дефолтного репозитория>/<имя приложения>:<версия приложения>
- пишет, что создано приложение <тэг>

Переменные $DAM_APP_NAME, $DAM_APP_VERSION, $DAM_APP_FAMILY получаются зарезервированы.

## ENVIRONMENT файл
ENVIRONMENT файл проверяется в директории проекта

Составлены критерии формата ENVIRONMENT файла:
- Пример `FOO=foo`
- Не игнорируются пробелы и табуляции
- комментарии с символа '#'
- нельзя переменную в несколько строчек
- нельзя комментарии в той же строке, что и переменная
- строчки без '=' игнорируются

## Dockerfile 
Чтобы извлечь переменные, необходимо придерживаться:
- Пример `ENV FOO=foo`
- Не игнорируются пробелы и табуляции
- нельзя переменную в несколько строчек
- нельзя комментарии в той же строке, что и переменная
- разделителем имени переменной могут быть символы "=" и " "
- переменные без значения игнорируются

## Meta
Генерирую файлы из шаблона .exp в той же директории meta.

# TODO
- Разобраться с созданием файла и маской
- в мапках лучше проверять, что поля ок
- перед копированием файла проверять, что не существует уже файл
- будет флаг, указывающий на проект, при сборке делается cd в проект
- будет валидация, что проект готов для создания приложения (описать в документации)
- следует создать флаг для указания docker build opts
- добавить информацию в вывод