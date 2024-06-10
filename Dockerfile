# 使用官方的Go语言镜像作为构建环境
FROM golang:1.22 AS build

# 设置工作目录
WORKDIR /app

# 复制所有文件到工作目录
COPY . .

# 下载依赖
RUN go mod download

# 编译Go应用
RUN go build -o deeplx

# 使用一个Ubuntu镜像作为运行环境
FROM ubuntu:latest

# 安装必要的包
RUN apt-get update && apt-get install -y tzdata ca-certificates

# 设置时区
ENV TZ=Etc/UTC

# 设置工作目录
WORKDIR /app

# 从构建环境复制编译好的二进制文件
COPY --from=build /app/deeplx .

# 复制配置文件
COPY config.yaml .

# 暴露应用运行的端口
EXPOSE 1188

# 运行应用
CMD ["./deeplx"]
