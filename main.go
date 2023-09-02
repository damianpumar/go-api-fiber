package main

import (
	"go-api-fiber-starterkit/server"
)

func main() {
	server.New().Listen(":3000")
}
