- External data
```shell
create /expense_tracker
```

```shell
create /expense_tracker/admin
create /expense_tracker/admin/server expense_admin
create /expense_tracker/admin/port :8000
```

```shell
create /expense_tracker/app
create /expense_tracker/app/server expense_app
create /expense_tracker/app/port :8001
```

```shell
create /expense_tracker/mongodb
create /expense_tracker/mongodb/expense
create /expense_tracker/mongodb/expense/uri "mongodb://localhost:27017"
create /expense_tracker/mongodb/expense/db_name "expenses"
```

```shell
create /expense_tracker/redis
create /expense_tracker/redis/expense
create /expense_tracker/redis/uri 127.0.0.1:6379
create /expense_tracker/redis/expense/auth
```

```shell
create /expense_tracker/admin/authentication
create /expense_tracker/admin/authentication/secret_key authentication
create /expense_tracker/admin/authentication/time_expired_token 15552000 // 6 months
```

```shell
create /expense_tracker/app/authentication
create /expense_tracker/app/authentication/secret_key authentication
create /expense_tracker/app/authentication/time_expired_token 15552000 // 6 months
```

