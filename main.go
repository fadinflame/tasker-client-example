package main

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"log"
	"net/http"
	taskerClients "tasker-client-example/clients"
	"tasker-client-example/consts"
	"tasker-client-example/models"
	"tasker-client-example/pb"
)

type TestClient interface {
	TestMethods(task models.Task) (bool, error)
}

func MakeTests(clientName string, tc TestClient) {
	initialTask := models.Task{Title: "Test Task", Text: "Buy some stuff"}
	success, err := tc.TestMethods(initialTask)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%s client tests success: %t\n\n", clientName, success)
}

func main() {
	var opts []grpc.DialOption
	opts = append(opts, grpc.WithTransportCredentials(insecure.NewCredentials()))

	conn, err := grpc.Dial(consts.GRPCConnAddr, opts...)
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	grpcClient := pb.NewTaskServiceClient(conn)
	httpClient := http.Client{}

	grpcTester := taskerClients.GrpcTest{GrpcClient: grpcClient, Ctx: context.Background()}
	httpTester := taskerClients.HttpClient{Client: httpClient, BaseURL: consts.HTTPBaseUrl}

	MakeTests("gRPC", &grpcTester)
	MakeTests("HTTP", &httpTester)
}
