# Overview

GraphQL servers in golang

## Quick start

Follow the steps below to use this service

1. Prepare Data

       curl --location --request POST 'http://localhost:7000/query' \
       --header 'Content-Type: application/json' \
       --data-raw '{"query":"mutation {\n  createPosts {\n    id\n  }\n  createComments {\n    id\n  }\n  createUsers {\n    id\n  }\n}","variables":{}}'

2. Start the graphql server

       make run

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
}
```

### Querys

- posts: [Post!]!
- comments: [Comment!]!
- users: [User!]!

### How to get all posts?

       curl --location --request POST 'http://localhost:7000/query' \
       --header 'Content-Type: application/json' \
       --data-raw '{"query":"{\n  posts {\n    id\n    title\n    body\n    url\n    type\n    userId\n    createdAt\n  }\n}","variables":{}}'

### How to get all comments?

       curl --location --request POST 'http://localhost:7000/query' \
       --header 'Content-Type: application/json' \
       --data-raw '{"query":"{\n  comments {\n    id\n    name\n    email\n    body\n    postId\n    createdAt\n  }\n}","variables":{}}'

### How to get all users?

       curl --location --request POST 'http://localhost:7000/query' \
       --header 'Content-Type: application/json' \
       --data-raw '{"query":"{\n  users {\n    id\n    name\n    username\n    email\n    photo\n    website\n    address {\n      street\n      suite\n      city\n      zipcode\n      geo {\n        lat\n        lng\n      }\n    }\n    company {\n      name\n      catchPhrase\n      bs\n    }\n    createdAt\n  }\n}","variables":{}}'