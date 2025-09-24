FROM alpine/sqlite:3.49.2

WORKDIR /database

EXPOSE 3306

ENTRYPOINT ["sqlite3"]