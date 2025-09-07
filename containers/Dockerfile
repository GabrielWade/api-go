from postgres:latest

copy create-database.sh /docker-entrypoint-initdb.d/

volume ["/etc/postgresql", "/var/log/postgresql", "/var/lib/postgresql"]

expose 5432