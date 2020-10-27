package main

import (
	"context"
	"fmt"

	"github.com/machinebox/graphql"
)

const (
	serverEndpoint = "http://localhost:8080/query"

	listBooks = `
query ($limit: Int) {
  list(limit: $limit) {
    name
    price
    description
  }
}
	`

	getBook = `
query ($name: String!) {
  book (name: $name) {
    name
    price
    description
  }
}
`

	createBook = `
mutation ($name: String!, $price: Float!, $description: String!) {
  create(name: $name, price: $price, description: $description) {
    name
    price
    description
  }
}
`

	updateBook = `
mutation ($name: String!, $price: Float, $description: String) {
  update(name: $name, price: $price, description: $description) {
    name
    price
    description
  }
}
`
	deleteBook = `
mutation ($name: String!) {
  delete(name: $name) {
    name
  }
}
`
)

func main() {
	// list books
	listReq := graphql.NewRequest(listBooks)
	listReq.Var("limit", 10)
	doRequest(listReq)

	// create
	createReq := graphql.NewRequest(createBook)
	createReq.Var("name", "history")
	createReq.Var("price", 9.99)
	createReq.Var("description", "this is an amazing historical book")
	doRequest(createReq)

	// get book by name
	getReq := graphql.NewRequest(getBook)
	getReq.Var("name", "history")
	doRequest(getReq)

	// uopdate
	updReq := graphql.NewRequest(updateBook)
	updReq.Var("name", "history")
	updReq.Var("price", 19.99)
	doRequest(updReq)

	// get book by name
	doRequest(getReq)

	// delete
	delReq := graphql.NewRequest(deleteBook)
	delReq.Var("name", "history")
	doRequest(delReq)

	// list books
	doRequest(listReq)
}

func doRequest(request *graphql.Request) {
	client := graphql.NewClient(serverEndpoint)

	request.Header.Set("Cache-Control", "no-cache")

	var response interface{}
	ctx := context.Background()
	err := client.Run(ctx, request, &response)
	if err != nil {
		fmt.Printf("ERROR: %s \n", err.Error())
	} else {
		fmt.Printf("Response: %+v \n", response)
	}
}
