// iniatilize the mod and sum file
package main

import (
	"context"
	"log"

	"gRPCtest/pb"

	"google.golang.org/grpc"
)

func main() {
	addr := "localhost:8080"
	conn, err := grpc.Dial(addr, grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()
	client := pb.NewWineClassifierServiceClient(conn)
	req := pb.WineReviewRequest{
		Review: "Dry with a fruity aroma",
	}

	resp, err := client.GetWineVariety(context.Background(), &req)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Predicted wine variety: %v", resp.Variety)
}

//protoc --go_out=plugins=grpc:/home/arijit/go/src/gRPCtest/pb --go_opt=paths=source_relative --proto_path=/home/arijit/go/src/gRPCtest /home/arijit/go/src/gRPCtest/wine_varieties.proto
