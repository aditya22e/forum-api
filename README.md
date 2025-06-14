Current Folder Structure

discussion-forum/
├── cmd/
│   └── discussion-forum/
│       └── main.go # entry point
├── api/
│   ├── handlers/
│   │   ├── topic.go
│   │   ├── user.go
│   │   └── comment.go
│   ├── middleware/
│   │   └── auth.go
│   └── routes/
│       └── routes.go
├── config/
│   └── db.go
├── models/
│   ├── user.go
│   ├── topic.go
│   ├── comment.go
│   └── subscription.go
├── services/
│   ├── topic.go
│   ├── user.go
│   ├── email.go
│   └── scheduler.go
├── utils/
│   └── jwt.go
├── tests/
│   └── topic_test.go
├── docs/
│   └── swagger.yaml
├── scripts/
│   └── init.sql
├── .env
├── .gitignore
├── go.mod
├── go.sum
└── README.md 
