package grpcapp

import (
	"context"
	"log"
	"time"

	pb "github.com/jnst/x-go/grpcapp/internal/gen/lol"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func StartClient() {
	const addr = "localhost:50051"
	conn, err := grpc.Dial(addr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("failed to connect: %v", err)
	}
	defer func(conn *grpc.ClientConn) {
		err := conn.Close()
		if err != nil {
			log.Fatalf("failed to close connection: %v", err)
		}
	}(conn)
	c := pb.NewLeagueOfLegendsClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	res, err := c.ListChampions(ctx, &pb.ListChampionsRequest{})
	if err != nil {
		log.Fatalf("failed to ListChampions: %v", err)
	}

	log.Printf("ListChampions: %+v", res.GetChampions())
}
