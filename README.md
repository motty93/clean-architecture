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

#### レイヤー構成
- ドメイン層(`internal/domain`)
    - モデル（エンティティ)
        - ファイル: `internal/domain/model/user.go`
    - ドメインサービス
        - ビジネスロジックをエンティティの操作に閉じ込める
        - ファイル: `internal/domain/service/user_service.go`
- インフラ(ストラクチャ)層(`internal/infrastructure`)
    - 外部サービスとの通信を行う
        - Supabaseの実装
            - ファイル: `internal/infrastructure/supabase_repository.go`
        - Slack通知の実装
            - ファイル: `internal/infrastructure/slack_service.go`
- プレゼンテーション層(`internal/interface`)
    - `http`, `api`, `presentation`というディレクトリ名でもＯＫ。今回はHTTPとCLIを同時に使用する想定で作成しているため、`interface`というディレクトリ名を使用している。
        - `http`: HTTPリクエストに特化している場合
        - `api`: RESTful APIやgRPCなど、特に外部APIに注力した設計の場合
        - `presentation`: Clean Architectureでの「プレゼンテーション層」という名前をそのまま使用する場合
    - ハンドラ
        - リクエストを受け取り、ユースケースを呼び出す
        - ファイル: `internal/interface/handler/user_handler.go`
    - CLI
        - コマンドラインインターフェース
        - ファイル: `internal/interface/cli/user_command.go`
- リポジトリ層(`internal/repository`)
    - データベースとの通信を行う
        - ファイル: `internal/repository/user_repository.go`
- ユースケース層(`internal/usecase`)
    -  アプリケーションの操作（例: ユーザー取得と通知）を提供する
        -  ファイル: `internal/usecase/user_usecase.go`
- エントリポイント層(`cmd/server`)
    - エントリポイントで依存関係を構築し、ユースケースを呼び出す
        - サーバー起動
        - ファイル: `cmd/server/main.go`
