FROM postgres:15.2-alpine

ENV POSTGRES_USER postgres
ENV POSTGRES_PASSWORD postgres
ENV POSTGRES_DB linkcabinet_db

COPY ./docker/postgreSQL/initdb/init.sql /docker-entrypoint-initdb.d/