# golang-curd
Go언어를 이용해서 C(Create)U(Update)R(Read)D(Delete)를 할 수 있는 ToDo List를 만드는 프로젝트입니다.  


## Installation
```shell
$ go mod init github.com/dev-hyunsang/golang-todo-list
$ go get -u github.com/gofiber/fiber/v2
$ go get -u gorm.io/gorm
$ go get -u gorm.io/driver/sqlite
```
- Golang Package
    - [`fiber/v2`](https://gofiber.io/) / Back-End Framework
    - [`gorm`](https://gorm.io/ko_KR/docs/index.html) / DataBase ORM
- DataBase 
    [`sqlite`](https://www.sqlite.org/index.html)