#!/bin/bash
set -e

# Create replication user
mysql -u root -p${MYSQL_ROOT_PASSWORD} <<EOF
CREATE USER 'replica'@'%' IDENTIFIED BY '${MYSQL_ROOT_PASSWORD}';
GRANT REPLICATION SLAVE ON *.* TO 'replica'@'%';
FLUSH PRIVILEGES;
ALTER USER 'replica'@'%' IDENTIFIED WITH mysql_native_password BY '${MYSQL_ROOT_PASSWORD}';
SHOW MASTER STATUS;
EOF
