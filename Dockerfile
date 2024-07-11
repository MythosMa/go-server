# 使用官方 Go 语言镜像作为基础镜像
FROM golang:1.22.3 AS builder

# 设置工作目录
WORKDIR /app

# 将 go.mod 和 go.sum 文件复制到工作目录
COPY go.mod go.sum ./

# 下载依赖
RUN go mod download

# 将项目的所有文件复制到工作目录
COPY . .

# 编译 Go 应用程序
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o /game-server cmd/server/main.go

# 使用一个更小的基础镜像
FROM alpine:latest

# 设置工作目录
WORKDIR /root/

# 将编译好的二进制文件从构建阶段复制到这个镜像
COPY --from=builder /game-server .

# 声明容器监听的端口
EXPOSE 9000

# 设置容器启动时执行的命令
CMD ["./game-server"]
