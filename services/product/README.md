# Service Product
Service product is service that related to product activity   
- PORT : 3031
- PATH : /

## Feature 
This service have some feature
- Create product
- Read product
- Update product
- Delete product

## Installation
``` bash
# clone the repo
$ git clone 

# go into app's directory
$ cd my-project
```

## Build & Run
``` bash
# run in docker
$ docker compose up 

# run in development 
$ ENV=DEV go run main.go
$ ENV=DEV ./filego

# run in staging 
$ ENV=STAGING go run main.go
$ ENV=STAGING ./filego

# run in production 
$ ENV=PROD go run main.go
$ ENV=PROD ./filego
```

## Swagger Documentation
Install go swagger
``` bash
# get swagger package 
$ go get github.com/swaggo/swag

# move to swagger dir
$ cd $GOPATH/src/github.com/swaggo/swag

# install swagger cmd 
$ go install cmd/swag
```

Generate API documentation
``` bash
# generate swagger doc
$ swag init --propertyStrategy snakecase
```
to see the results, run app and access {{base_url}}/swagger/index.html

# Author
[Muhammad Cholis Malik](https://www.linkedin.com/in/mcholismalik/)