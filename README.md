# studygroup-backend

## Developer Guild
### go version
18.6

### env config example
```
export DB_RECONNECT_TIMES=5
export DB_RECONNECT_BOUNCE_SEC=1
export DB_HOST=localhost
export DB_PORT=5432
export DB_USER=postgres
export DB_PASSWORD=password
export DB_NAME=postgres
export DB_DRIVER=postgres

export SERVER_ADDR=0.0.0.0:8080

export REDIS_CONNECT_SIZE=10
export REDIS_NETWORK=tcp
export REDIS_ADDR=localhost:6379
export REDIS_PASSWORD=

```

### Swagger
This server use swagger to generate api docs
Developor sholud remember to update the swagger docs by using `swag` cli of golang

You can reference to this github [swaggo/gin-swagger](https://github.com/swaggo/gin-swagger) for more details

#### Install 
```
go install -v github.com/swaggo/swag/cmd/swag
```

setup path, you can consider to add this into your bashrc or zshrc
```
export PATH=$(go env GOPATH)/bin:$PATH
```

#### Usage
* Run this command after you finished your work
  This command can be used to generate/update the swagger files
  ```
    swag init -g <path to main.go>
  ```
