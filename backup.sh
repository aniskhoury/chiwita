mysqldump -u chiwita -p --no-data chiwita > database.sql
mariadb -u chiwita -p < database.sql