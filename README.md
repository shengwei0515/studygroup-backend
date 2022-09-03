# studygroup-backend

## Developer Guild
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
    go init -g <path to main.go>
  ```
