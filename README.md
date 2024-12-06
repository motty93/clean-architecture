## Clean Architecture
### 環境
- Golang 1.22.5
### 起動
```bash
go mod tidy

air
```
### 概要

#### ディレクトリ構成
```
.
├── cmd
│   └── server
│       └── main.go
├── go.mod
├── go.sum
├── internal
│   ├── domain
│   │   ├── model
│   │   │   └── user.go
│   │   └── service
│   │       ├── user_service.go
│   │       └── user_service_test.go
│   ├── infrastructure
│   │   ├── database.go
│   │   ├── slack_service.go
│   │   ├── slack_service_test.go
│   │   └── supabase_repository.go
│   │   └── supabase_repository_test.go
│   ├── interface
│   │   ├── cli
│   │   │   ├── user_command.go
│   │   │   └── user_command_test.go
│   │   └── handler
│   │       ├── user_handler.go
│   │       └── user_handler_test.go
│   ├── repository
│   │   ├── user_repository.go
│   │   └── user_repository_test.go
│   ├── scheduler
│   │   └── job_scheduler.go
│   │   └── job_scheduler_test.go
│   └── usecase
│       ├── user_usecase.go
│       └── user_usecase_test.go
```
