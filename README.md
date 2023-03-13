# Summary 

## 1. Local Development Guideline
### Prerequisites
 - Docker & Docker Compose
 - Go (v1.18 or above), Echo Framework

### Setup Local Development Environment**
1. Clone the project to local machine and go to the folder
```
git clone https://github.com/thanhtuan472k/expense-tracker-server
cd expense-tracker-server 
```
2. Run ``make setup`` to install dependencies
3. Run ``make mongo-seed`` to setup the local DB (Run migration and seeding)
4. Run ``make run-app`` to start server API app
5. Run ``make run-admin`` to start server API admin

**Note**: The local DB will use port 27017. If the port is being used, please change it to a difference port in Zookeeper.md and set port again



