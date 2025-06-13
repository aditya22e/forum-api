

Folder Structure: 


forum-api-main/
│
├── config/                          # Configuration files
│   └── config.go                    # Loads environment variables
│
├── controllers/                    # HTTP Handlers (Controllers)
│   ├── auth.go                      # User register/login handlers
│   ├── comment.go                   # Comment creation handler
│   ├── discussion.go                # Topic CRUD handlers
│   └── subscription.go              # Subscription logic handler
│
├── middleware/                     # Middleware (e.g., auth)
│   └── jwt.go                       # JWT token validation
│
├── models/                         # Structs for database entities
│   ├── comment.go                   # Comment model
│   ├── discussion.go                # Topic model
│   ├── subscription.go             # Subscription model
│   └── user.go                      # User model
│
├── repository/                     # Database operations
│   ├── comment_repository.go        # DB methods for comments
│   ├── discussion_repository.go     # DB methods for topics
│   ├── subscription_repository.go   # DB methods for subscriptions
│   └── user_repository.go           # DB methods for users
│
├── routes/                         # Route grouping
│   └── routes.go                    # Initialize all routes
│
├── utils/                          # Utility functions
│   ├── auth.go                      # JWT generation/parsing
│   └── validator.go                 # Request validation logic
│
├── .env                             # Environment configuration
├── go.mod                           # Module file
├── go.sum                           # Dependency checksums
└── main.go                          # Application entry point
