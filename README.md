Current Folder Structure


└───forum-api-main
    │   go.mod
    │   go.sum
    │   
    ├───api
    │   ├───handlers
    │   │       comment.go
    │   │       topic.go
    │   │       user.go
    │   │       
    │   ├───middleware
    │   │       auth.go
    │   │
    │   └───routes
    │           routes.go
    │
    ├───cmd
    │   └───forum
    │           main.go
    │
    ├───config
    │       db.go
    │
    ├───models
    │       comment.go
    │       subscription.go
    │       topic.go
    │       user.go
    │
    ├───scripts
    │       init.sql
    │
    ├───services
    │       email.go
    │       scheduler.go
    │       topic.go
    │       user.go
    │
    ├───tests
    │       topic_tests.go
    │
    └───utils
            jwt.go
