package main

import (
	proto "HW4/grpc"
	"context"
	"fmt"
	"log"
	"math/rand/v2"
	"net"
	"sync"
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
}

var hasToken int = 0
var requestToken bool = false

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

	log.Printf("Node started")

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
	log.Printf("Node is connecting to the next node")

	conn, err := grpc.Dial(node.nextPort, grpc.WithInsecure(), grpc.WithBlock())

	if err != nil {
		log.Fatalf("Failed to dial server: %v", err)
	}
	nextNode = proto.NewHomeworkFourServiceClient(conn)
}

func (node *Node) SendIDToNextClient(ctx context.Context, req *proto.IDSendRequest) (*proto.IDSendResponse, error) {
	if node.ownPort == ":5050" {
		node.NodeID = 0
		node.NextNodeID = 1
	} else if node.nextPort == ":5050" {
		node.NodeID = req.SenderID + 1
		node.NextNodeID = 0
	} else {
		node.NodeID = req.SenderID + 1
		node.NextNodeID = req.SenderID + 2
	}

	return &proto.IDSendResponse{
		Success: true,
	}, nil
}

func main() {
	var WaitGroup sync.WaitGroup
	// Create the nodes
	fmt.Println("Insert port here:")
	var thisPort string
	fmt.Scanln(&thisPort)

	var targetPort string
	fmt.Println("Insert target port here:")
	fmt.Scanln(&targetPort)

	var thisNode *Node

	if thisPort == "5050" {
		thisNode = &Node{NodeID: 0, NextNodeID: 1, ownPort: ":5050", nextPort: ":" + targetPort}

		hasToken = 1

		log.Println("Node created")
		go thisNode.startNode()
		thisNode.ConnectToNextNode()

		WaitGroup.Add(1)
		go func() {
			defer WaitGroup.Done()
			nextNode.SendIDToNextClient(context.Background(), &proto.IDSendRequest{
				SenderID: 0,
			})
		}()

	} else {
		thisNode = &Node{NodeID: 0, NextNodeID: 1, ownPort: ":" + thisPort, nextPort: ":" + targetPort}

		hasToken = 1

		log.Println("Node created")
		go thisNode.startNode()
		thisNode.ConnectToNextNode()

		WaitGroup.Add(1)
		go func() {
			defer WaitGroup.Done()
			nextNode.SendIDToNextClient(context.Background(), &proto.IDSendRequest{
				SenderID: thisNode.NodeID,
			})
		}()
	}

	// Wait for the ID to be updated
	WaitGroup.Wait()

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
					Token:    0,
					SenderID: thisNode.NodeID,
				})
				if err != nil {
					log.Fatalf("Failed to send token to next node: %v", err)
				}

				time.Sleep(2 * time.Second)
			}
		}
	}
}
