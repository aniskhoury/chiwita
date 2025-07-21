mysqldump -u chiwita -p --no-data chiwita > basededatos.sql
mariadb -u chiwita -p < basededatos.sql