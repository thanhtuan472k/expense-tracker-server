# Summary **🚀**

## 1. Functionalities

- Admin
    - Login with account Root
    - Management information
    - Management categories (CRUD)
    - Management sub-categories (CRUD)
- User
    - Register account
    - Login/logout, using JWT for authentication
    - Expense-tracking:
        - Classify money (Income, Expenses)
        - Add expenses with the following details: date, category, description, and amount.
        - Edit expense
    - View statistic money
        - Users should be able to view their expenses in a list or a chart, filter expenses by category or date range.

## 2. **Local Development Guideline**

### **Prerequisites**

- Docker & Docker Compose
- Go (v1.18 or above), Echo Framework

### **Setup Local Development Environment**

1. Clone the project to local machine and go to the folder
    
    ```jsx
    git clone https://github.com/thanhtuan472k/expense-tracker-server
    cd expense-tracker-server
    ```
    
2. Clone sub module to local machine
    
    ```jsx
    git clone https://github.com/thanhtuan472k/expense-tracker-server
    cd expense-tracker-server
    ```
    
3. Run `make setup` to install package dependencies
4. Find the file `/expense-tracker-server/Zookeeper.md`. If you don’t have it, create one following the template in `/expense-tracker-server/Zookeeper.md`
5. Update the value in that file
    
    ```jsx
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
    ```
    
6. This file should **NEVER** be committed to source control because it contains sensitive information
7. Run `make mongo-seed` to set up the local DB (Run migration and seeding)
8. Run `make run-app` to start the server API app
9. Run `make-run-admin` to start the server API admin

**Notes**: The local DB will use port 27017. If the port is being used, please change it to a different port in `expense-tracker-server/Zookeeper.md` and set the value to Zookeeper Env

## 3. Other Notes

### What can be improved

1. More test cases for back-end
2. Write some end-to-end tests
3. Build and deploy to the cloud (AWS, EC2)
4. Upload image