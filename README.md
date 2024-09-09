# Демо-проект: приложение по учету личных доходов и расходов "Family budget".
## Описание

Данный проект реализует приложение по учету личных доходов и расходов.
Проект все еще в работе, поэтому некоторого функционала может не хватать, тесты будут добавлены в ближайшее время. 
Составные части приложения:

- REST API
- Серверная часть приложения
- БД PostgreSQL

 ## Функции

- Ввод произвольного количества счетов
- Ввод произвольного количества категорий доходов и расходов. Возможность учета категорий в режиме многоуровневой иерархии.
- Ввод записей о доходах и расходах денежных средств в разрезе счетов, категорий

 ## Технологии
Для работы Family budget были использованы следующие пакеты и технологии:
- [cleanenv] - config reader
- [postgreSQL] - database
- [migrate] - database migrations
- [log/slog] - loger
- [net/http]- HTTP-server
- [chi] - http-service router 
- [middleware] - middleware handlers
- [render] - HTTP request/response manager
- [validator]- value validator for structs 


   [log/slog]: <https://pkg.go.dev/log/slogr>
   [chi]: <https://github.com/go-chi/chi>
   [cleanenv]: <https://github.com/ilyakaznacheev/cleanenv>
   [postgreSQL]: <https://www.postgresql.org> 
   [migrate]: <https://github.com/golang-migrate/migrate>
   [middleware]: <https://github.com/go-chi/chi>
   [render]: <https://github.com/go-chi/render>
   [validator]: <https://github.com/go-playground/validator>
   [net/http]: <https://pkg.go.dev/net/http>
  
