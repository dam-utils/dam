# dam export

Команда предназначена для экспорта списка установленных приложений в текстовый файл.
В текстовом файле содержится информация об имени и версии приложения.
Обратите внимание, что информация о репозитории-источнике приложения хранится в метке DAM_APP_SERVERS и manifest.json
самих приложений, а не в файле экспрота.
Если при экспорте указан флаг `--all`, то результирующий файл будет являться архивом. 
В котором вместе с текстовым файлом создадутся docker образы приложений в gz формате.

# Процесс

- если флаг --all не указан, то список приложений (имя приложения и тэг без имени репозитория)
сохраняется в указанный текстовый файл
- если флаг указан, то текстовый файл сохраняется в tar-архиве.
В этот же архив сохраняются сжатые gz образы приложений.