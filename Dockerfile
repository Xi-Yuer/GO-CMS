# 基于 Node.js 镜像创建一个 Docker 镜像
FROM node:18

# 进入前端目录
RUN cd /web

# 将当前工作目录设置为/app
WORKDIR /app

# 将 package.json 和 package-lock.json 复制到 /app 目录下
COPY package*.json ./

# 运行 npm install 安装依赖
RUN npm install

# 将源代码复制到 /app 目录下
COPY . .

# 打包构建
RUN npm run build

# 将构建后的代码复制到 nginx 镜像中
FROM nginx:latest
COPY --from=0 /app/dist /usr/share/nginx/html

# 暴露容器的 8080 端口，此处其实只是一个声明作用，不写的话也可以，后面运行容器的
# docker run --name container_name -p <host_port>:<container_port>命令中container_port可以覆盖此处的声明，不写就默认80端口
EXPOSE 8080

# 启动 nginx 服务
CMD ["nginx", "-g", "daemon off;"]
