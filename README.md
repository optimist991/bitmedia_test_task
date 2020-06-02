# bitmedia_test_task

## 1: Prepare project env
#### Create app database:
####  Install MongoDb:
```bash
sudo apt-get install -y mongodb-org

sudo systemctl start mongod
```
#### Or follow the installation instructions for MongoDB from the official website:
https://docs.mongodb.com/manual/tutorial/install-mongodb-on-ubuntu/
#### Check installation:
```bash
sudo systemctl status mongod
```

## 2: Getting Started

#### To install all dependencies run:
```bash
dep ensure -v
```
#### To add users from json:
```bash
cd users_migration

go run upload_users.go
```
#### To run app:
```bash
bee run -downdoc=true -gendoc=true
```
#### Swagger docs:
```bash
http://127.0.0.1:8080/swagger/
```
