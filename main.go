package main

import (
	"fmt"
	"net/http"
	"time"

	"github.com/hashicorp/vault/api"
)

const (
	url = "http://127.0.0.1:8200"

	token = "STATIC_TOKEN"
)

var httpClient = &http.Client{
	Timeout: 10 * time.Second,
}

func main() {

	client, err := api.NewClient(&api.Config{Address: url, HttpClient: httpClient})
	if err != nil {
		panic(err)
	}

	client.SetToken(token)
	results, err := client.Logical().Read("kv/data/customer-service")
	if err != nil {
		fmt.Println(err)
		return
	}

	secrets, _ := results.Data["data"].(map[string]interface{})

	fmt.Println(secrets)

}
