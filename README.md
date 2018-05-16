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

error 500 - fatal error
