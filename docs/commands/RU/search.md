# dam search

Команда осуществляет поиск приложений в репозиториях по маске.
Если маска не указана, то выводится список всех приложений.
В выводе указывается имя приложения и список доступных для него версий.
В маске указывается только часть имени приложения, но не версии.

Поиск осуществляется в следующем порядке:
* определяется дефолтный репозиторий
* если дефолтный репозиторий официальный, поиск осуществляется в стандартном порядке
* если дефолтный репозиторий не официальный, происходит проверка необходима ли авторизация.
Для этого проверяется настройки пользователя, указан ли его логин.
Логин указан, происходит авторизация
* первоначально авторизация идет по протоколу https.
В случае неудачи происходит попытка по протоколу http. Иначе выводится ошибка
* далее делается запрос на список приложений в репозитории
* список фильтруется по маске
* для каждого приложения делается запрос версий и выводится пользователю