# Subscriptions API
Simple subscriptions service

You can run this app with:
```
docker compose up
```
You need to have .env file in the working directory with these vars in it:
```
# These are set automatically when using docker compose
DB_HOST=db
DB_PORT=5432
# When using docker compose these values will be used
# to initialize the PostgreSQL database
DB_NAME=db
DB_USER=admin
DB_PASSWORD=admin
# The port at which the API will be available
HTTP_PORT=8080
```

Swagger documentation can be accessed on /swagger/index.html
