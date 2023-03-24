# gql-todo

## Introduction

This is a demo project implementing several graphQL-related techniques in go. 
And also to be served as a resource for my future referencing.  

## Get Started

### How to set up environment 

Run **MySQL** in docker
```bash
docker run --name gql-todo-mysql \
-p 3306:3306 \
-e MYSQL_ROOT_PASSWORD=myrootsecretpassword \
-e MYSQL_USER=user \
-e MYSQL_PASSWORD=mysecretpassword \
-e MYSQL_DATABASE=gql-todo \
-d mysql:8.0.32
```

Create gql-todo database
```bash
docker exec -it gql-todo-mysql bash

mysql -u root -p
#myrootsecretpassword

CREATE DATABASE gql-todo
```

Run Application
```bash
go run server.go
```

## Feature Implementations

- [x] Custom Data Validation
- [x] Cursor-based pagination - Connections
- [ ] Optimizing N+1 database queries

### Related Resources

- [Gqlgen Custom Data Validation](https://david-yappeter.medium.com/gqlgen-custom-data-validation-part-1-7de8ef92de4c)
- [Explaining GraphQL Connections](https://www.apollographql.com/blog/graphql/explaining-graphql-connections/)
- [Dataloaders](https://gqlgen.com/reference/dataloaders/)