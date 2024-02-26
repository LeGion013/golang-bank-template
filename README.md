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

* How to run the docker image in DEBUG mode: 
```
docker run --name banktemplate -p 8080:8080 banktemplate:latest
```

* How to run the docker image in RELEASE MODE: 
```
docker run --name banktemplate -p 8080:8080 -e GIN_MODE=release banktemplate:latest
```

* How to run the docker image in RELEASE MODE + additional environment variable: 
```
docker run --name banktemplate -p 8080:8080 -e GIN_MODE=release -e DB_SOURCE="postgresql://root:secret@172.17.0.2:5432/bank_template?sslmode=disable" banktemplate:latest
```

* How to run the docker image in RELEASE MODE + additional environment variable + network: 
```
docker run --name banktemplate --network bank-network -p 8080:8080 -e GIN_MODE=release -e DB_SOURCE="postgresql://root:secret@postgres15:5432/bank_template?sslmode=disable" banktemplate:latest
```


* How to get the docker image settings/information: 
```
docker container inspect <container_id>
```

* How to get the docker networks: 
```
docker network ls 
```

* How to get the docker networks settings/information: 
```
docker network inspect <network_name>
```

* How to create new network for docker: 
```
docker network create <network_name> 
```

* How to add the docker image to newly created network: 
```
docker network connect <network_name> <container_id>
```

### Generate random

* How to generate random string: 
```
openssl rand -hex 64
```

* How to generate random string and take first 32: 
```
openssl rand -hex 64 | head -c 32
```


### Install aws cli

* Use curl:
```
curl "https://awscli.amazonaws.com/awscli-exe-linux-x86_64.zip" -o "awscliv2.zip"
unzip awscliv2.zip
sudo ./aws/install
```

* aws configure


```
ls -l ~/.aws
```

* get secrets:
```
aws secretsmanager get-secret-value --secret-id name_of_the_secret
```

* get secrets with filter:
```
aws secretsmanager get-secret-value --secret-id simple_bank --query SecretString
```

* get secrets with filter and in text fromat:
```
aws secretsmanager get-secret-value --secret-id simple_bank --query SecretString --output text
```

* get secrets with filter and in text fromat and write them to the app.env:
```
aws secretsmanager get-secret-value --secret-id simple_bank --query SecretString --output text | jq -r 'to_entries|map("\(.key)=\(.value)")|.[]' > app.env
```

export PATH=$PATH:/usr/local/go/bin