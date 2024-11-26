# 基础镜像
FROM golang:1.23.2-alpine AS builder

# 设置工作目录
WORKDIR /app

# 将代码复制到镜像中
COPY . .

# 下载依赖并编译应用
RUN go mod init simple-go-app && \
    go build -o app main.go

# 创建运行时镜像
FROM alpine:latest
WORKDIR /root/
COPY --from=builder /app/app .

# 暴露服务端口
EXPOSE 8932

# 启动应用
CMD ["./app"]
