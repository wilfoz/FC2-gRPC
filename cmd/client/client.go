package main

import (
	"fmt"
	"log"
	"io"
	"context"
	"time"
	"google.golang.org/grpc"
	"github.com/wilfoz/FC2-gRPC/pb"
)
 
func main () {
	connection, err := grpc.Dial("localhost:50051", grpc.WithInsecure())

	if err != nil {
		log.Fatalf("Could not connect to gRPC Server: %v", err)
	}

	defer connection.Close()

	client:= pb.NewUserServiceClient(connection)
	//AddUser(client)
	//AddUserVerbose(client)
	AddUsers(client)
}

func AddUser(client pb.UserServiceClient) {
	req := &pb.User {
		Id: "0",
		Name: "Carlos",
		Email: "carlos@email.com",
	}

	res, err := client.AddUser(context.Background(), req)
	if err != nil {
		log.Fatalf("Could not make gRPC request: %v", err)
	}

	fmt.Println(res)
}

func AddUserVerbose(client pb.UserServiceClient) {
	req := &pb.User {
		Id: "0",
		Name: "Carlos",
		Email: "carlos@email.com",
	}

	responseStream, err := client.AddUserVerbose(context.Background(), req)
	if err != nil {
		log.Fatalf("Could not make gRPC request: %v", err)
	}

	for {
		stream, err := responseStream.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatalf("Could not receive the msg: %v", err)
		}
		fmt.Println("Status:", stream.Status)
	}
}

func AddUsers(client pb.UserServiceClient) {
	reqs := []*pb.User{
		&pb.User{
			Id: "2",
			Name: "Marcos",
			Email: "m@email.com",
		},
		&pb.User{
			Id: "3",
			Name: "Lucas",
			Email: "l@email.com",
		},
		&pb.User{
			Id: "4",
			Name: "Carlos",
			Email: "c@email.com",
		},
		&pb.User{
			Id: "5",
			Name: "Michel",
			Email: "m@email.com",
		},
	}

	stream, err := client.AddUsers(context.Background())
	if err != nil {
		log.Fatalf("Error creating request: %v", err)
	}

	for _, req := range  reqs {
		stream.Send(req)
		time.Sleep(time.Second * 3)
	}

	res, err := stream.CloseAndRecv()
	if err != nil {
		log.Fatalf("Error receiving response: %v", err)
	}

	fmt.Println(res)
}