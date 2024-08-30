package grpcClient

import (
	"context"
	"fmt"
	pb "github.com/henrylamb/object-generation-golang/grpc"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
	"time"
)

// SendRequestToServer sends a request to the gRPC server with authorization headers
func (c *Client) GrpcGenerateObject(serverAddress string, prompt string, definition *pb.Definition, authToken string) (*pb.Response, error) {
	// Set up a connection to the gRPC server
	conn, err := grpc.Dial(serverAddress, grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		return nil, fmt.Errorf("failed to connect to server: %v", err)
	}
	defer func(conn *grpc.ClientConn) {
		err = conn.Close()
		if err != nil {

		}
	}(conn)

	// Create a new client from the gRPC service
	client := pb.NewJSONSchemaServiceClient(conn)

	// Create a context with a timeout
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()

	// Set up metadata with the authorization token
	md := metadata.New(map[string]string{"authorization": "Bearer " + authToken})
	ctx = metadata.NewOutgoingContext(ctx, md)

	// Create the request object
	request := &pb.RequestBody{
		Prompt:     prompt,
		Definition: definition,
	}

	// Call the gRPC method on the client
	response, err := client.GenerateObject(ctx, request) // Replace 'GenerateObject' with the actual RPC method name
	if err != nil {
		return nil, fmt.Errorf("failed to call GenerateObject: %v", err)
	}

	return response, nil
}
