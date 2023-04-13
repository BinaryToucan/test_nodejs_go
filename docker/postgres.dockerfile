FROM postgres:15.1-alpine

LABEL author="Emiliya Melnikova"
LABEL description="Postgres Image for test"
LABEL version="1.0"

COPY *.sql /docker-entrypoint-initdb.d/