package book

type Book struct {
	Name        string
	Price       float64
	Description string
}

type SetResponse struct {
	Status     string      `json:"status"`
	Data       interface{} `json:"data,omitempty"`
	AccessTime string      `json:"accessTime"`
}
