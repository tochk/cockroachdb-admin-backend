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


### Get tables "/api/tables/"
Request:
```
{
    "token": "sample token",
    "db": "db_name"
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

### Create table "/api/tables/create/"

Request:
```
{
    "token": "sample token",
    "db": "db_name",
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
    "db": "db_name",
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

error 500 - fatal error
