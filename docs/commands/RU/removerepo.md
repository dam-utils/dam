# dam removerepo

Команда удаляет выбранный репозиторий из настроек.
В качестве аргумента может быть, как id репозитория, так и его имя.

Дефолтный репозиторий можно удалить лишь добавив к команде флаг `--force`.
Без этого флага репозиторий можно удалить только сняв настройку default.

При удалении дефолтного репозитория, дефолтным становится репозиторий с id равным 1.

Репозиторий с id равному 1 официальный docker hub. Его удалить невозможно.