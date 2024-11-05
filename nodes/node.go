package main

import (
	proto "HW4/grpc"
	"context"
	"fmt"
	"log"
	"net"

	"google.golang.org/grpc"
)

type Node struct {
	proto.UnimplementedHomeworkFourServiceServer
	activeNodes map[string]proto.HomeworkFourServiceClient
	//idToPort    map[int32]string
	//Need to figure out how to check if id is correct
	NodeID     int32
	NextNodeID int32
	port       string
	hasToken   bool
}

func (node *Node) SendTokenToNextCLient(ctx context.Context, req *proto.TokenSendRequest) (*proto.TokenSendResponse, error) {
	pls, ok := node.activeNodes[node.port]
	if !ok {
		log.Fatalf("Could not find next node")
		return nil, nil
	}

	if req.Token == 1 {
		node.hasToken = true
	}

	node.useToken()

	tokenMeassage := &proto.TokenSendRequest{
		Token:    req.Token,
		SenderID: node.NodeID,
	}

	//THIS IS WRONG

	pls.SendTokenToNextCLient(ctx, tokenMeassage)

	return nil, nil
}

func (node *Node) startNode() error { //maybe rename to StartServer for more clarity
	listener, err := net.Listen("tcp", node.port)

	if err != nil {
		return err
	}

	grpcServer := grpc.NewServer() //maybe rename

	proto.RegisterHomeworkFourServiceServer(grpcServer, node)

	serve := grpcServer.Serve(listener)
	return serve
}

func (node *Node) connectToNode(forwardPort string) error {
	conn, err := grpc.Dial(forwardPort, grpc.WithInsecure())
	if err != nil {
		return nil
	}

	clientNode := proto.NewHomeworkFourServiceClient(conn)

	//dobbelcheck if this is correct
	node.activeNodes[forwardPort] = clientNode
	//node.idToPort[node.NodeID] = port NEED TO ADD SOMETHING LIKE THIS
	return nil
}

func (node *Node) useToken() error {
	if node.hasToken {
		log.Print(fmt.Sprintf("Node: %d has accessed and used the critical section", node.NodeID))
	}
	return nil
}

/*
type Node struct {
	proto.UnimplementedHomeworkFourServiceServer
	currentNodes map[int32]proto.HomeworkFourServiceClient

	NodeID     int32
	NextNodeID int32
}

var

func main() {

	//Maybe use some logic like this to create the nodes
	//var nodeCount int = 3


		for i := 0; i < nodeCount; i++ {
			// Create a new node

			if i == nodeCount-1 {
				node := &Node{
					currentNodes: make(map[int32]proto.TokenSend),
					NodeID:       int32(i),
					NextNodeID:   int32(0),
				}
			} else {

				node := &Node{
					currentNodes: make(map[int32]proto.TokenSend),
					NodeID:       int32(i),
					NextNodeID:   int32(i + 1),
				}
			}
		}


	node := &Node{
		//currentNodes: make(map[int32]proto.TokenSend),
		//NodeID:       int32(0),
		//NextNodeID:   int32(1),
	}

	node.start_node()

}

func (n *Node) SendTokenToNextCLient(ctx context.Context, stream *proto.TokenSend) error {

	// Save the client stream
	n.currentNodes[n.NextNodeID].SendTokenToNextCLient(ctx, stream.Send())

	// Send the token to the next node
	n.currentNodes[n.NextNodeID].Send(&proto.TokenSend{Token: 1})
	return nil
}

func (node *Node) start_node() {
	// Start the nodes
	node.currentNodes = make(map[int32]proto.HomeworkFourServiceClient)

	listener, err := net.Listen("tcp", "localhost:5050")

	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	grpcServer := grpc.NewServer() // n is for serving purpose

	proto.RegisterHomeworkFourServiceServer(grpcServer, node)

	err = grpcServer.Serve(listener)

	if err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}

}*/
