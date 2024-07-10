# 使用官方的 Golang 镜像作为基础镜像
FROM golang:1.21.1-alpine

# 设置工作目录
WORKDIR /app

# 将应用程序的二进制文件(auth) 和配置文件(etc) 复制到容器中
COPY role .
COPY etc ./etc

# 暴露服务的端口
EXPOSE 8882

# 定义容器启动时运行的命令
CMD ["./role", "-f", "./etc/role.yaml"]
