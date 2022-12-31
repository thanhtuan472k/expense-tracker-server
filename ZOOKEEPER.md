- External data
```shell
create /expense_tracker
```

```shell
create /expense_tracker/admin
create /expense_tracker/admin/server expense_admin
create /expense_tracker/admin/port:
```

```shell
create /expense_tracker/app
create /expense_tracker/app/server expense_app
create /expense_tracker/admin/port :8001
```

```shell
create /expense_tracker/mongodb
create /expense_tracker/mongodb/expense
create /expense_tracker/mongodb/expense/uri "mongodb://localhost:27017"
create /expense_tracker/mongodb/expense/db_name "expenses"
```

```shell
create /expense_tracker/redis/expense
create /expense_tracker/redis/uri 127.0.0.1:6379
create /expense_tracker/redis/expense/auth
create /expense_tracker/redis/expense/auth
```