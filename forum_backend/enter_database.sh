source .env
mysql -h 127.0.0.1 -P 3307 -u root -p${MYSQL_ROOT_PASSWORD} qq_forum --default-character-set=utf8mb4