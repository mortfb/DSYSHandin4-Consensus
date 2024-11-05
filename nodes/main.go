package main

import (
	proto "HW4/grpc"
	"context"
	"log"
	"math/rand"
	"sync"
)

var nodes = make(map[int32]*Node)

var waitGroup sync.WaitGroup

func main() {

	numberNodes := 3

	// Create the nodes
	for i := 0; i < numberNodes; i++ {
		if i == numberNodes-1 {
			node := &Node{
				activeNodes: make(map[string]proto.HomeworkFourServiceClient),
				NodeID:      int32(i),
				NextNodeID:  0,
				port:        ":505" + string(i),
			}
			nodes[node.NodeID] = node
		} else {
			node := &Node{
				activeNodes: make(map[string]proto.HomeworkFourServiceClient),
				NodeID:      int32(i),
				NextNodeID:  int32(i + 1),
				port:        ":505" + string(i),
			}
			nodes[node.NodeID] = node
		}
	}

	waitGroup.Add(numberNodes)

	// Run the nodes
	for _, node := range nodes {
		go node.runNode()
	}
	waitGroup.Wait()

}

func (node *Node) runNode() error {
	defer waitGroup.Done()

	if err := node.startNode(); err != nil {
		return err
	}

	if err := node.connectToNode(nodes[node.NextNodeID].port); err != nil {
		return err
	}

	var requestingToken bool = false

	for {
		//Code that checks if the node has the token
		var useToken int
		if !requestingToken {
			useToken = rand.Intn(10)
		}

		if useToken == 5 {
			requestingToken = true
		}

		if requestingToken {
			//Need to add some logic to request and handle the token
		}

		if node.hasToken {
			log.Print("Node: %d has accessed and used the critical section", node.NodeID)
			node.hasToken = false
			nodes[node.NextNodeID].SendTokenToNextCLient(context.Background(), &proto.TokenSendRequest{
				Token: 1,
			})

		}

	}

	return nil
}
