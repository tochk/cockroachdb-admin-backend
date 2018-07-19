# cockroachdb-admin-backend

## API reference

### Connect "/api/connect/"
Request:
```
{
    "login": "root",
    "password": "test"
}
```

Response:

```
{
    "token": "sample token"
}
```

### Get databases "/api/databases/"
Request:
```
{
    "token": "sample token"
}
```

Response:

```
[
    {
        "database": "db_name"
    },
    ...
]
```


### Create database "/api/databases/create/"

Request:
```
{
    "token": "sample token",
    "database": "db_name"
}
```

Response:

Databases list.

```
[
    {
        "database": "db_name"
    },
    ...
]
```

### Drop database "/api/databases/drop/"

Request:
```
{
    "token": "sample token",
    "database": "db_name"
}
```

Response:

Databases list.

```
[
    {
        "database": "db_name"
    },
    ...
]
```

### Get tables "/api/tables/"
Request:
```
{
    "token": "sample token",
    "database": "db_name"
}
```

Response:

```
[
    {
        "table": "table_name"
    },
    ...
]
```

### NOT COMPLETED! Create table "/api/tables/create/"

Request:
```
{
    "token": "sample token",
    "database": "db_name",
    "table": "table_name"
}
```

Response:

Tables list.

```
[
    {
        "table": "table_name"
    },
    ...
]
```

### Drop table "/api/tables/drop/"

Request:
```
{
    "token": "sample token",
    "database": "db_name",
    "table": "table_name"
}
```

Response:

Tables list.

```
[
    {
        "table": "table_name"
    },
    ...
]
```

### Get data "/api/data/"
Request:
```
{
    "token": "sample token",
    "database": "db_name",
    "table": "table_name",
    //optional
    "limit": 10,
    "offset": 10
}
```

Response:

```
[
    {
        "column": "data"
        ...
    },
    ...
]
```


### Get indexes "/api/indexes/"

Request:
```
{
    "token": "sample token",
    "database": "db_name",
    "table": "table_name"
}
```

Response:

```
[
    {
        "table":"departments",
        "name":"primary",
        "unique":true,
        "seq":1,
        "column":"rowid",
        "direction":"ASC",
        "storing":false,
        "implicit":false
    }
    ...
]
```

### Get data "/api/query/"
Request:
```
{
    "token": "sample token",
    "database": "db_name",
    "query": "SELECT..."
}
```

Response:

if select statement
```
[
    {
        "column": "data"
        ...
    },
    ...
]
```

else 
```
{ 
    "status": 200
}

```


### Errors
Any kind of error sends json reply like this:

```
{
    "code": 500,
    "human": "Fatal error",
    "err": "idk what is happening"
}
```


## Errors

error 1 - json parse error

error 2 - database connection error

error 3 - get databases error

error 4 - invalid token

error 5 - get tables error

error 6 - get data error

error 7 - create table error

error 8 - drop table error

error 9 - create database error

error 10 - drop database error

error 11 - get indexes error

error 12 - query error

error 500 - fatal error
