# blog

**config.yml file need to add as shown below**

```
  app:
    name: ""

  server:
    host: ""
    port: ""

  db:
    username: ""
    password: ""
    host: ""
    port: ""
    name: ""
```

**Usefull CLI commands**

```
go run main.go help
go run main.go migrate
go run main.go seed
go run main.go serve
```

**go run main.go help**
It's provide details of available commands

**go run main.go migrate**
It's create a table in DB using models

**go run main.go seed**
using for dummy entries in table

**go run main.go serve**
It's running server
