# GraphQL CRUD API with Golang

Simple CRUD GraphQL API with Golang and MySQL.

## Run Local

```bash
make dev
```

Open browser: http://localhost:8080/graphql

## Test API

### Create User

```bash
mutation {
  createUser(name: "John Doe", age: 30) {
    id
    name
    age
  }
}
```

### Update User By ID

```bash
mutation {
  updateUser(id: 1, name: "Jane Doe", age: 31) {
    id
    name
    age
  }
}
```

### Delete User By ID

```bash
mutation {
  deleteUser(id: 1)
}
```

### Get All Users

```bash
query {
  users {
    id
    name
  }
}
```

### Get User By ID

```bash
query {
  user(id: 1) {
    id
    name
  }
}
```

## Make a request in Postman or Curl

```bash
curl -X POST -H "Content-Type: application/json" \
-d '{"query": "{ users { id name age } }"}' \
http://localhost:8080/graphql
```
