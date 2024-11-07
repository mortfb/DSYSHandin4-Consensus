package main

import (
	proto "HW4/grpc"
	"context"
	"fmt"
	"log"
	"math/rand/v2"
	"net"
	"time"

	"google.golang.org/grpc"
)

var nextNode proto.HomeworkFourServiceClient

type Node struct {
	proto.UnimplementedHomeworkFourServiceServer
	NodeID     int32
	NextNodeID int32
	ownPort    string
	nextPort   string
	hasToken   int
}

var hasToken int = 0
var requestToken bool = false

//var nodes = []*Node{}

func (node *Node) SendTokenToNextCLient(ctx context.Context, req *proto.TokenSendRequest) (*proto.TokenSendResponse, error) {
	// Implement the method logic here

	if req.Token == 1 {
		hasToken = int(req.Token)
	}

	return &proto.TokenSendResponse{
		Success: true,
	}, nil
}

func (node *Node) startNode() error { //maybe rename to StartServer for more clarity
	listener, err := net.Listen("tcp", node.ownPort)

	log.Printf("Node %d started", node.NodeID)

	if err != nil {
		return err
	}

	grpcServer := grpc.NewServer() //maybe rename

	proto.RegisterHomeworkFourServiceServer(grpcServer, node)

	if err := grpcServer.Serve(listener); err != nil {
		return err
	}

	return nil
}

func (node *Node) ConnectToNextNode() {
	log.Printf("Node %d is connecting to the next node", node.NodeID)

	conn, err := grpc.Dial(node.nextPort, grpc.WithInsecure(), grpc.WithBlock())

	if err != nil {
		log.Fatalf("Failed to dial server: %v", err)
	}
	nextNode = proto.NewHomeworkFourServiceClient(conn)
}

func main() {
	// Create the nodes

	//maybe scan for ports, then make the nodes
	fmt.Println("Insert port here:")
	var thisPort string
	fmt.Scanln(&thisPort)

	var thisNode *Node

	if thisPort == "5050" {
		thisNode = &Node{NodeID: 0, NextNodeID: 1, ownPort: ":5050", nextPort: ":5051", hasToken: 1}

		hasToken = 1

		log.Println("Node 0 created")
		go thisNode.startNode()
		thisNode.ConnectToNextNode()

	}

	if thisPort == "5051" {
		thisNode = &Node{NodeID: 1, NextNodeID: 2, ownPort: ":5051", nextPort: ":5052", hasToken: 0}
		log.Println("Node 1 created")
		go thisNode.startNode()
		thisNode.ConnectToNextNode()

	}

	if thisPort == "5052" {
		thisNode = &Node{NodeID: 2, NextNodeID: 0, ownPort: ":5052", nextPort: ":5050", hasToken: 0}

		log.Println("Node 2 created")
		go thisNode.startNode()
		thisNode.ConnectToNextNode()

	}

	for {
		if !requestToken {
			var random = rand.IntN(4)
			if random == 1 {
				requestToken = true
				log.Printf("Node %d is requesting the token", thisNode.NodeID)
			}
		}
		if nextNode != nil {
			if hasToken == 1 {
				log.Printf("Node %d has the token", thisNode.NodeID)
				if requestToken {
					log.Printf("Node %d is using the toke to access the critical section", thisNode.NodeID)
					time.Sleep(2 * time.Second)
					requestToken = false
				}
				_, err := nextNode.SendTokenToNextCLient(context.Background(), &proto.TokenSendRequest{
					Message:  "now you have the token",
					Token:    1,
					SenderID: thisNode.NodeID,
				})
				hasToken = 0
				if err != nil {
					log.Fatalf("Failed to send token to next node: %v", err)
				}
				log.Printf("Node %d sends token to %d", thisNode.NodeID, thisNode.NextNodeID)

			} else {
				_, err := nextNode.SendTokenToNextCLient(context.Background(), &proto.TokenSendRequest{
					Message:  "you dont have the token",
					Token:    0,
					SenderID: thisNode.NodeID,
				})
				if err != nil {
					log.Fatalf("Failed to send token to next node: %v", err)
				}

				//log.Printf("Node %d sends message to %d", thisNode.NodeID, thisNode.NextNodeID)
				time.Sleep(2 * time.Second)
			}
		}
	}
}
