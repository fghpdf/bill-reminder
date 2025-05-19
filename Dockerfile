# 构建阶段（必须启用 CGO）
FROM golang:1.24-bullseye AS builder

WORKDIR /app

# 安装构建 SQLite 所需的依赖
RUN apt-get update && apt-get install -y gcc libsqlite3-dev

COPY go.mod go.sum ./
RUN go mod download

COPY . .

# ✅ 启用 CGO
ENV CGO_ENABLED=1
RUN go build -o /app/bill-reminder

# 运行阶段（可以选择更小的 debian-slim）
FROM debian:stable-slim

WORKDIR /app

# 运行时需要 libc + sqlite3 动态库
RUN apt-get update && apt-get install -y libsqlite3-0 ca-certificates && apt-get clean

COPY --from=builder /app/bill-reminder .

VOLUME ["/app/sqlite"]

EXPOSE 8080

CMD ["./bill-reminder"]