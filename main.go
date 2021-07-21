package main

import (
	"go-rest-test/web"
)

type IndexClient struct {
	HttpClient *web.HttpClient
}

func NewIndexClient() *IndexClient {
	c := &IndexClient{
		HttpClient: web.NewHttpClient(),
	}
	return c
}

func main() {
	server := NewIndexClient()
	server.HttpClient.Run(":80")
}
