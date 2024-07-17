# MySQL 基于 Docker 设置主从复制

## 1. 创建 MySQL 主从复制容器(使用dockerCompose的方式快速运行容器)
```bash
services:
  mysql:                                                                          # 数据库(主库)
    image: mysql:8.0
    container_name: mysql
    environment:
      MYSQL_DATABASE: micro_cms                                                   # 数据库(该容器会自动创建改数据库)
      MYSQL_ROOT_PASSWORD: 2214380963Wx!!                                         # 添加 MySQL 的 root 密码
    ports:
      - "3306:3306"                                                               # 将容器的 3306 端口映射到宿主机的 3306 端口
    networks:
      - microservices-network                                                     # 确保容器在同一个网络环境中
    volumes:
      - type: bind
        source: ./sql/master.cnf                                                  # 挂载配置文件(配置文件所在位置)
        target: /etc/mysql/conf.d/master.cnf                                      # 挂载配置文件(配置文件被挂载到容器中的位置)
        read_only: true                                                           # 只读
        consistency: consistent                                                   # 一致性(保了卷被挂载为具有一致性)
      - type: bind
        source: ./sql/master_init.sql                                             # 挂载初始化 SQL 文件(初始化 SQL 文件所在位置)
        target: /docker-entrypoint-initdb.d/master_init.sql                       # 挂载初始化 SQL 文件(配置文件被挂载到容器中的位置)
        read_only: true                                                           # 只读
        consistency: consistent                                                   # 一致性(保了卷被挂载为具有一致性)

  mysql-slave:
    image: mysql:8.0
    container_name: mysql-slave
    environment:
      MYSQL_DATABASE: micro_cms
      MYSQL_ROOT_PASSWORD: 2214380963Wx!!
    ports:
      - "3307:3306"
    volumes:
      - ./sql/slave.cnf:/etc/mysql/conf.d/slave.cnf:ro,consistent
      - ./sql/master_init.sql:/docker-entrypoint-initdb.d/master_init.sql:ro,consistent
    depends_on:
      - mysql
    command:
      - --default-authentication-plugin=mysql_native_password
      - --server-id=2
      - --relay-log=relay-log
      - --log-bin=mysql-bin
      - --read-only=1
      - --skip-slave-start
    networks:
      - microservices-network
```

## 配置主从库
在主库配置文件中，需要添加以下配置：
```
server-id=1
log-bin=mysql-bin
binlog-do-db=micro_cms
```
在从库的配置文件中，需要添加以下配置：
```
server-id=2
relay-log=relay-log
log-bin=mysql-bin
read-only=1
binlog-do-db=micro_cms
```

## 初始化数据库
确保主库和从库都初始化了数据库，这里我使用了 `master_init.sql` 文件，并配置在 `docker-entrypoint-initdb.d` 目录下，这样在容器启动时就会自动执行该文件中的 SQL 语句

## 启动容器

## 主从库设置
在主库上执行以下命令：
```sql
# 创建一个用户，用于从库复制主库的 Binlog 日志文件
# 创建了一个名为replica的用户，可以从任意主机连接到MySQL服务器（因为@'%'表示任意主机）。用户的密码为2214380963Wx!!
CREATE USER 'replica'@'%' IDENTIFIED BY '2214380963Wx!!';
# 授予用户replica对所有数据库的复制权限。这是MySQL主从复制所必须的权限。
GRANT REPLICATION SLAVE ON *.* TO 'replica'@'%';
# 刷新权限
FLUSH PRIVILEGES;

# 修改用户replica的认证插件为mysql_native_password，并重新设置用户的密码为2214380963Wx!!
ALTER USER 'replica'@'%' IDENTIFIED WITH mysql_native_password BY '2214380963Wx!!';

# 查看主库的状态（获取主库的 MASTER_LOG_POS 和 MASTER_LOG_FILE 从库需要用到）
SHOW MASTER STATUS;
```
在从库上执行以下命令：
```sql
# 设置主库信息
CHANGE MASTER TO MASTER_HOST='mysql', MASTER_USER='replica', MASTER_PASSWORD='2214380963Wx!!', MASTER_LOG_FILE='binlog.000002', MASTER_LOG_POS=1175;
# 启动从库
START SLAVE;
# 查看从库状态
SHOW SLAVE STATUS;
# 停止从库
STOP SLAVE;
# 重置从库
RESET SLAVE;
```

## 测试
在主库上插入一条数据，然后在从库上查询，如果能够查询到，说明主从库配置成功

## 注意
1. 主库和从库都需要初始化数据库以及数据表


