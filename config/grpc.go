package config

import (
	"log"
	"os"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func GPTService() *grpc.ClientConn {
	host := os.Getenv("GPT_SERVICE_HOST")

	conn, err := grpc.Dial(host, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Println("could not connect to ", host, err)
	}

	return conn
}
