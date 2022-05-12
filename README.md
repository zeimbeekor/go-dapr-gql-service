# GO-DAPR-QGL-SERVICE

## Overview

GraphQL servers in golang

## Pre-requisites

- [Latest version of Go](https://go.dev/dl/)
- [Docker Desktop](https://www.docker.com/products/docker-desktop/)
- [Dapr CLI and initialized environment](https://docs.dapr.io/getting-started)

## Quick start

Follow the steps below to use this service

1. Copy dapr bindings spec files

```console
$ cp ./dapr/bindings/components/jsonbinding.yaml $HOME/.dapr/components/jsonbinding.yaml
```
```console
$ cp ./dapr/bindings/components/graphql.yaml $HOME/.dapr/components/graphql.yaml
```

2. Prepare Data

```console
$ curl --location --request POST 'http://localhost:7000/query' \
--header 'Content-Type: application/json' \
--data-raw '{"query":"mutation {\n  createTodo {\n    id\n  }\n}","variables":{}}'
```

3. Run the API Graphql service alongside a Dapr sidecar

```console
$ make run
```

More help to get started [gqlgen](https://github.com/99designs/gqlgen)

## Reporting Issues

If you think you've found a bug, or something isn't behaving the way you think it should, please raise an [issue](https://github.com/zeimbeekor/go-dapr-gql-service/issues) on GitHub.

## Frequently asked questions

### Models

```graphql
type Post {
  id: Int!
  title: String!
  body: String!
  url: String!
  type: String!
  userId: Int!
  createdAt: String!
  comments: [Comment]
}
```

```graphql
type Comment {
  id: Int!
  name: String!
  email: String!
  body: String!
  postId: Int!
  createdAt: String!
}
```

```graphql
type Geo {
  lat: String!
  lng: String!
}
```

```graphql
type Address {
  street: String!
  suite: String!
  city: String!
  zipcode: String!
  geo: Geo!
}
```

```graphql
type Company {
  name: String!
  catchPhrase: String!
  bs: String!
}
```

```graphql
type User {
  id: Int!
  name: String!
  username: String!
  email: String!
  photo: String!
  website: String!
  createdAt: String!
  address: Address!
  company: Company!
  posts: [Post]
}
```

```graphql
type Todo {
  id: Int!
  title: String!
  completed: Boolean!
  user: User!
}
```

### Querys

- posts: `[Post!]!`
- comments: `[Comment!]!`
- todos: `[Todo!]!`
- users: `[User!]!`
- user(id: Int!): `User!`
- post(id: Int!): `Post!`
- comment(id: Int!): `Comment!`
- getInfoByUserId(id: Int!): `User!`
- getPostsByUserId(id: Int!, postType: String!): `[Post!]!`
- getCommentsByPostId(id: Int!): `[Comment!]!`

### Mutations

- createTodo: `[Todo!]!`

### Postman Collection

Examples: [misc/urth.postman_collection.json](https://github.com/zeimbeekor/go-dapr-gql-service/blob/master/misc/urth.postman_collection.json)

### How to get all `posts`?

```console
$ curl --location --request POST 'http://localhost:7000/query' \
--header 'Content-Type: application/json' \
--data-raw '{"query":"{\n  posts {\n    id\n    title\n    body\n    url\n    type\n    userId\n    createdAt\n  }\n}","variables":{}}'
```

### How to get all `comments`?

```console
$ curl --location --request POST 'http://localhost:7000/query' \
--header 'Content-Type: application/json' \
--data-raw '{"query":"{\n  comments {\n    id\n    name\n    email\n    body\n    postId\n    createdAt\n  }\n}","variables":{}}'
```

### How to get all `users`?

```console
$ curl --location --request POST 'http://localhost:7000/query' \
--header 'Content-Type: application/json' \
--data-raw '{"query":"{\n  users {\n    id\n    name\n    username\n    email\n    photo\n    website\n    address {\n      street\n      suite\n      city\n      zipcode\n      geo {\n        lat\n        lng\n      }\n    }\n    company {\n      name\n      catchPhrase\n      bs\n    }\n    createdAt\n  }\n}","variables":{}}'
```

## Contributing

* [Alvaro Vega Plata](https://github.com/zeimbeekor)

## License

Check the [LICENSE](LICENSE) file for details