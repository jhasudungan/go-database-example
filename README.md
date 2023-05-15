# Go-Database-Example

An example project about connecting and querying database (RDBMS) in go program. I use Postgre and pq-lib standard library here.

## Using this Project

To use this project :

```bash
git pull {githubtothisproject}
```

create a **.env** in /.env and fill it with your database configuration. You could find the environtment variable in **env-example.txt**

```bash
DB_HOST=
DB_PORT=
DB_DATABASE_NAME=
DB_USER=
DB_PASSWORD=

LIMIT_PAGINATION_PER_PAGE=
```

Run the **master.sql** in this project. After that you could run all the test from terminal using:

```bash
go test --v -run=TestConnect
go test --v -run=TestCreateConnectionString
go test --v -run=TestFindAll
go test --v -run=TestFindWithPage
go test --v -run=TestInsert
go test --v -run=TestUpdate
go test --v -run=TestDelete
```
