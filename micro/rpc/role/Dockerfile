# 使用官方的 Go 语言镜像作为基础镜像
FROM golang:1.21.1-alpine

# 设置 Go 模块代理
ENV GOPROXY=https://goproxy.cn,direct

# 设置工作目录
WORKDIR /app

COPY . .

# 创建一个目录用于存放编译后的二进制文件和配置文件
RUN mkdir -p ./bin

# 安装依赖
RUN go mod download

# 编译 Go 程序
RUN go build -o ./bin/cmd ./rpc/role/role.go

COPY ./rpc/role/etc ./bin/etc

# 暴露服务的端口
EXPOSE 8888

# 定义容器启动时运行的命令
CMD ["./bin/cmd", "-f", "./bin/etc/role.yaml"]
