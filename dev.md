# Make your own yata-api

1. init the module: run `go mod init yata-api`.
2. install relevant modules:
    1. chi: `go get “github.com/go-chi/chi”`, http router
    2. renderer: `“github.com/thedevsaddam/renderer”`, response rendering package
    3. pq: `go get github.com/jackc/pgx/v5` postgres client for go.
