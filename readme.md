## Cтандартный веб сервер (Standard Web Server)

Базовый шаблон GO сервера с логгированием для построения дальнейшего REST API сервера.  


Пример передачи конфигурационного файла при запуске приложения:
```./api -path configs/api.toml```


Зависимости:
- Библиотека для логов: go get -u github.com/sirupsen/logrus
- Конфигурация http сервера: go get -u github.com/gorilla/mux
- Работа с toml файлом: go get -u github.com/BurntSushi/toml
- Для работы с postgres:  go get -u github.com/lib/pq 





