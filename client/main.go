package main

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"

	"github.com/machinebox/graphql"
)

const (
	serverEndpoint = "http://localhost:8080/query"

	getOrders = `
		query{
			orders{
				id
				orderAmount
				items {
					productName
				}
			}
		}
	`
	createOrder = `
mutation CreateOrder($newOrder: OrderInput!) {
  createOrder(input: $newOrder) {
    id
    customerName
    orderAmount
    items 
    	{
        id
        productCode
        productName
        quantity
      }
  }
}
`
	updateOrder = `
		{
			"orderId":1,
			"input": {
				"customerName": "Cristiano",
				"orderAmount": 9.99,
				"items": [
					{
						"productCode": "2323",
						"productName": "IPhone X",
						"quantity": 1
					}
				]
			}
		}
	`
	deleteOrder = `
		{
			"orderId": 3
		}
	`
)

func main() {
	useLibrary()

	//useHttpClient()
}

func useLibrary() {
	client := graphql.NewClient(serverEndpoint)

	//request := graphql.NewRequest(getOrders)
	request := graphql.NewRequest(createOrder)
	request.Header.Set("Cache-Control", "no-cache")

	var response interface{}
	ctx := context.Background()
	err := client.Run(ctx, request, &response)
	if err != nil {
		panic(err)
	}

	fmt.Println(response)
}

func useHttpClient() {
	jsonData := map[string]string{
		"query": `
            orders {
				id  
				customerName
				items {
					productName
					quantity
				}
			}
        `,
	}

	jsonValue, _ := json.Marshal(jsonData)

	request, err := http.NewRequest(http.MethodPost, serverEndpoint, bytes.NewBuffer(jsonValue))

	client := &http.Client{Timeout: time.Second * 10}

	response, err := client.Do(request)
	if err != nil {
		fmt.Printf("The HTTP request failed with error %s\n", err)
	}
	defer response.Body.Close()

	data, _ := ioutil.ReadAll(response.Body)
	fmt.Println(string(data))
}
