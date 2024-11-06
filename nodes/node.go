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
	ownPort    string
	nextPort   string
	hasToken   bool
}

func (node *Node) SendTokenToNextCLient(ctx context.Context, req *proto.TokenSendRequest) (*proto.TokenSendResponse, error) {
	pls, ok := node.activeNodes[node.nextPort]
	if !ok {
		log.Fatalf("Could not find next node")
		return nil, nil
	}

	if req.Token == 1 {
		node.hasToken = true
	}

	//node.useToken()

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
