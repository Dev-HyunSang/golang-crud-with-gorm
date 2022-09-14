# golang-todo-list
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

## ToDos:
- [X] 새로운 ToDO 항목 만들기
- [X] 생성되어 있는 ToDo 항목들 불러오기
- [X] UUID를 통해서 생성되어 있는 ToDo 항목 수정하기 
- [X] 생성되어 있는 ToDo 항목 삭제하기

## Docs
### POST `/create`
#### Request
```json
{
    "content": "안녕하세요!!"
}
```

### Response
```json
{
    "message": "새로운 할일 항목을 만들었어요!",
    "status": 200,
    "time": "2022-06-05T16:51:32.68058+09:00"
}
```

### POST `/read`
#### Request
```json
Null
```

#### Response
```json
{
    "datas": [
        {
            "todo_uuid": "bf8de572-cf5d-4f4b-8552-581015aae720",
            "content": "안녕하세요!!",
            "created_time": "2022-06-05T16:51:32.678999+09:00",
            "updated_time": "2022-06-05T16:51:32.678999+09:00"
        }
    ],
    "message": "성공적으로 생성되어 있는 항목들을 가지고 왔어요!",
    "status": 200,
    "time": "2022-06-05T16:52:45.074149+09:00"
}
```

### POST `/edit`

#### Request
```json
{
    "todo_uuid": "bf8de572-cf5d-4f4b-8552-581015aae720",
    "content": "반갑습니다~!"
}
```

#### Response
```json
{
    "message": "성공적으로 수정했어요!",
    "status": 200,
    "time": "2022-06-05T16:53:04.78207+09:00"
}
```

### Delete `/delete/:UUID`
#### Request
```
127.0.0.1:3000/delete/07265781-593b-4ef0-a25f-158a3460ebb7
```

#### Response
```json
{
    "message": "성공적으로 할일 항목을 삭제했어요.",
    "status": 200,
    "time": "2022-06-05T16:49:55.773332+09:00"
}
```
