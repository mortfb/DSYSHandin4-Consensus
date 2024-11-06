package main

import (
	proto "HW4/grpc"
	"fmt"
	"net"

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

var nodes = []*Node{}
var hasToken bool

func (node *Node) ConnectToNextNode() error {
	conn, err := grpc.Dial(node.nextPort, grpc.WithInsecure())
	if err != nil {
		return err
	}

	nextNode = proto.NewHomeworkFourServiceClient(conn)
	fmt.Println(string(node.NodeID) + " Connected to next node " + string(node.NextNodeID))

	return nil
}

func (node *Node) SendTokenToNextCLient(stream proto.HomeworkFourService_SendTokenToNextCLientServer) error {

	// Implement the method logic here
	return nil
}

func (node *Node) startNode() error { //maybe rename to StartServer for more clarity
	listener, err := net.Listen("tcp", node.ownPort)

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

func main() {
	// Create the nodes
	if len(nodes) == 0 {
		nodeA := &Node{
			NodeID:     0,
			NextNodeID: 1,
			ownPort:    ":5050",
			nextPort:   ":5051",
		}

		nodes = append(nodes, nodeA)
		nodeA.startNode()

	} else if len(nodes) == 1 {
		nodeB := &Node{
			NodeID:     1,
			NextNodeID: 2,
			ownPort:    ":5051",
			nextPort:   ":5052",
		}

		nodes = append(nodes, nodeB)
		nodeB.startNode()

	} else if len(nodes) == 2 {
		nodeC := &Node{
			NodeID:     2,
			NextNodeID: 0,
			ownPort:    ":5052",
			nextPort:   ":5050",
		}
		nodes = append(nodes, nodeC)
		nodeC.startNode()
	}
}

/*
	// Create the nodes "automatically"
	for i := 0; i < numberNodes; i++ {
		if i == numberNodes-1 {
			node := &Node{
				NodeID:      int32(i),
				NextNodeID:  0,
				ownPort:     ":505" + string(i),
				nextPort:    ":5050",
			}
			nodes = append(nodes, node)
		} else {
			node := &Node{
				NodeID:      int32(i),
				NextNodeID:  int32(i + 1),
				ownPort:     ":505" + string(i),
				nextPort:    ":505" + string(i+1),
			}
			nodes = append(nodes, node)
		}
	}

	// Run the nodes
	for _, node := range nodes {
		go node.startNode()
	}
*/
