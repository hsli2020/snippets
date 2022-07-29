
go install github.com/melkeydev/go-blueprint@latest
go-blueprint create
go-blueprint create --name my-project --framework gin --driver postgres
go-blueprint create -h

- [Chi](https://github.com/go-chi/chi)
- [Gin](https://github.com/gin-gonic/gin)
- [Fiber](https://github.com/gofiber/fiber)
- [HttpRouter](https://github.com/julienschmidt/httprouter)
- [Gorilla/mux](https://github.com/gorilla/mux)
- [Echo](https://github.com/labstack/echo)

- [Mysql](https://github.com/go-sql-driver/mysql)
- [Postgres](https://github.com/jackc/pgx/)
- [Sqlite](https://github.com/mattn/go-sqlite3)
- [Mongo](https://go.mongodb.org/mongo-driver)
- [Redis](https://github.com/redis/go-redis)

- [HTMX](https://htmx.org/) support using [Templ](https://templ.guide/)
- CI/CD workflow setup using [Github Actions](https://docs.github.com/en/actions)
- [Websocket](https://pkg.go.dev/nhooyr.io/websocket) sets up a websocket endpoint

Blueprint UI [go-blueprint.dev](https://go-blueprint.dev)

go-blueprint create --name my-project --framework gin --driver postgres
go-blueprint create --advanced
go-blueprint create --advanced --feature htmx
go-blueprint create --advanced --feature githubaction
go-blueprint create --advanced --feature websocket
go-blueprint create --name my-project --framework chi --driver mysql --advanced --feature htmx --feature githubaction --feature websocket
