# Summary **🚀**

## 1. Functionalities

- Admin
    - Login with account Root
    - Management infomation
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

3. Run ``make setup`` to install package dependencies
4. Run ``make mongo-seed`` to setup the local DB (Run migration and seeding)
5. Run ``make run-app`` to start server API app
6. Run ``make-run-admin`` to start server API admin

## 3. Other Notes
### What can be improved

1. More test case for back-end
2. Write some end-to-end tests
3. Build and deploy to cloud