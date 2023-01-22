# golang-bank-template


### Migrate

* How to run the migrate program 
* Check migrate help:
```
migrate -help
```
* Create new migration script for db:
```
migrate create -ext sql -dir db/migration -seq add_users
```

### Docker build

* How to build the docker image: 
```
docker build -t banktemplate:latest .
```
