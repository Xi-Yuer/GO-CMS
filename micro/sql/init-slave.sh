#!/bin/bash
set -e

# Wait for the master to be fully up and running
until mysql -h mysql -u root -p${MYSQL_ROOT_PASSWORD} -e "select 1"; do
  >&2 echo "MySQL master is unavailable - sleeping"
  sleep 10
done

# Get master status
MASTER_STATUS=$(mysql -h mysql -u root -p${MYSQL_ROOT_PASSWORD} -e "SHOW MASTER STATUS\G")
CURRENT_LOG=$(echo "$MASTER_STATUS" | grep "File:" | awk '{print $2}')
CURRENT_POS=$(echo "$MASTER_STATUS" | grep "Position:" | awk '{print $2}')

# Configure replication
mysql -u root -p${MYSQL_ROOT_PASSWORD} <<EOF
STOP SLAVE;
RESET SLAVE;
CHANGE MASTER TO
  MASTER_HOST='mysql',
  MASTER_USER='replica',
  MASTER_PASSWORD='2214380963Wx!!',
  MASTER_LOG_FILE='$CURRENT_LOG',
  MASTER_LOG_POS=$CURRENT_POS;
FLUSH PRIVILEGES;
START SLAVE;
EOF
