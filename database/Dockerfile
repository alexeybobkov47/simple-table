FROM postgres:13.1
COPY ./sql/*.sql /docker-entrypoint-initdb.d/
ENV POSTGRES_USER=postgres
ENV POSTGRES_PASSWORD mypassword 
ENV POSTGRES_DB simpleTable
ENV PGDATA /var/lib/pgdata