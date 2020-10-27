
## Graph*i*QL

## Queries

###
```
query listBooks ($limit: Int) {
  list(limit: $limit) {
    name
    price
    description
  }
}
```

#### Variables

```json
{
  "limit": 10
}
```

### get single book by name
```
query getBook ($name: String!) {
  book (name: $name) {
    name
    price
    description
  }
}
```

#### Variables

```json
{
  "name": "history"
}
```

## Mutations

## Create
```
mutation createCalculusBook {
  create(name: "calculus", price: 9.99, description: "this is a calculus amazing book") {
    name
    price
    description
  }
}
```

```
mutation createHistoryBook {
  create(name: "history", price: 9.99, description: "this is a history fascinating book") {
    name
    price
    description
  }
}
```

## Update
```
mutation updateBook ($name: String!) {
  update(name: $name, price: 19.99) {
    name
    price
    description
  }
}
```

### Variables

```json
{
  "name": "history"
}
```

## Delete
```
mutation deleteBook ($name: String!) {
  delete(name: $name) {
    name
  }
}
```

### Variables

```json
{
  "name": "history"
}
```
