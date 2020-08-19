
# go-graphql

Sample project to understand GraphQL implementation in Golang

## Run

1. start PostgreSQL

    ```shell script
    make run-postgres
    ```

2. start go-graphql

    ```shell script
    cd custom
    make run
    ```

## Links

### Server

- https://graphql.org/
- https://gqlgen.com/
- https://github.com/99designs/gqlgen
- https://www.soberkoder.com/go-graphql-api-mysql-gorm/
- https://hub.docker.com/_/mysql

#### Issues

There is [an issue](https://github.com/99designs/gqlgen/issues/1283) with `gqlgen v0.12.x`, for now please use `v0.11.3`

### Client

- https://blog.machinebox.io/a-graphql-client-library-for-go-5bffd0455878
- https://www.thepolyglotdeveloper.com/2020/02/interacting-with-a-graphql-api-with-golang/
- https://github.com/machinebox/graphql
- https://altair.sirmuel.design/
