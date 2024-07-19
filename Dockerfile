# 设置默认环境变量为pro，可以通过 --build-arg 更改
ARG APP_ENV=pro
# 使用 Golang 官方镜像作为构建环境
FROM golang:1.22 as builder
LABEL authors="nan"
# 设置工作目录
WORKDIR /app

# 将代码复制到容器中
COPY . .

# 设置环境变量，使用ARG中的值
ENV APP_ENV=${APP_ENV}

# 下载依赖
RUN go mod tidy

# 编译应用程序
RUN go build -o main ./cmd/main.go
# RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o main ./cmd/main.go

# 使用 scratch 作为最终运行镜像，这是一个空白的镜像，用于创建尽可能小的容器
FROM alpine:latest

WORKDIR /root/

# 从构建器镜像中复制二进制文件到最终镜像
COPY --from=builder /app/main .

COPY config-pro.yaml .
COPY config-dev.yaml .

# 声明运行时容器提供服务的端口
EXPOSE 8990

# 运行二进制
CMD ["./main"]