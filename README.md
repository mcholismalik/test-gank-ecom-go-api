# Gank Ecommerce
Gank Ecommerce is project for technical test from Gank Global that built with Golang Echo Framework

## Services 
This project consist of some services
- Auth 
- Product
- Order 

## Dependencies 
This project consist of some dependencies
- Postgres
- Elasticsearch

## Pre-requisite
``` bash
# Install docker desktop
https://www.docker.com/products/docker-desktop
```

## Run
``` bash
# move to deployments
$ cd deployments

# run dependencies, make sure it running well before run services 
$ ./run-dependencies.sh 

# run services
$ ./run-services.sh 
```

## Access
``` bash
# Access auth service 
$ localhost:3030/swagger/index.html

# Access product service 
$ localhost:3031/swagger/index.html

# Access order service 
$ localhost:3032/swagger/index.html
```

## Author
[Muhammad Cholis Malik](https://www.linkedin.com/in/mcholismalik/)