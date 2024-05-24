package main

import (
	"fmt"
	"keystone_auth/rpc"
	"keystone_auth/src"
	"net/http"
)

func main() {
	server := &src.Server{}
	grpcHandler := rpc.NewAuthServiceServer(server)

	err := http.ListenAndServe(":8080", grpcHandler)
	if err != nil {
		fmt.Println("Error starting server:", err)
		return
	}
}
