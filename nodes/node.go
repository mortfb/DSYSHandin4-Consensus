package main

import (
	proto "HW4/grpc"

	"google.golang.org/grpc"
)

var Stream proto.HomeworkFourService_SendTokenToNextCLientClient

type HomeworkFourServiceClient struct {
	proto.UnimplementedHomeworkFourServiceServer
	currentNodes map[int32]proto.HomeworkFourService_SendTokenToNextCLientClient

	NodeID     int32
	NextNodeID int32
}

func main() {
	server := &HomeworkFourServiceClient{
		currentNodes: make(map[int32]proto.HomeworkFourService_SendTokenToNextCLientClient),
	}
	proto.RegisterHomeworkFourServiceServer(grpc.NewServer(), server)

}
