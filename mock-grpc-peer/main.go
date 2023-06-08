package main

// picked up only a simple hlf chaincode from fabric contract api go/simple
// run it
// chaincode looking for peer on port 7051

// write a simple grpc server acting as mock peer
// chaincode looks for ChaincodeSupport service on mock peer
// added peer protos from fabric protos that provide services and payload
// chaincode and mock peer comminication is bidrectional grpc, payload is ChaincodeMessage
// chaincode is able to connect and register itself with peer

import (
	"fmt"
	"net"

	pb "mock-grpc-peer/pb"

	log "github.com/sirupsen/logrus"
	"google.golang.org/grpc"
)

type ChaincodeSupport struct {
	pb.ChaincodeSupportServer
}

func main() {
	// create listener
	// register rpc
	// start rpc server
	lis, err := net.Listen("tcp", ":7051")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	s := grpc.NewServer()
	pb.RegisterChaincodeSupportServer(s, &ChaincodeSupport{})
	s.Serve(lis)
}

// func (cs *ChaincodeSupport) ProcessProposal(ctx context.Context, in *pb.SignedProposal) (*pb.ProposalResponse, error) {
// 	return nil, nil
// }

func (cs *ChaincodeSupport) Register(x pb.ChaincodeSupport_RegisterServer) error {
	fmt.Println("hihi hahaha hahaha")

	for i := 0; i < 10; i++ {
		if msg, err := x.Recv(); msg != nil {
			fmt.Println("received message from chaincode: ")
			fmt.Println(msg, err)
			x.Send(&pb.ChaincodeMessage{
				Type: pb.ChaincodeMessage_REGISTERED,
			})
		} else {
			fmt.Println("received message from chaincode: ")
			fmt.Println(msg, err)
			x.Send(&pb.ChaincodeMessage{
				Type: pb.ChaincodeMessage_COMPLETED,
			})
		}
	}
	// x.Context().Done()
	return nil
}

type ChaincodeSupport_RegisterServer struct {
	pb.ChaincodeSupport_RegisterServer
}

func (x *ChaincodeSupport_RegisterServer) Send(m *pb.ChaincodeMessage) error {
	return nil
}

func (x *ChaincodeSupport_RegisterServer) Recv() (*pb.ChaincodeMessage, error) {
	return nil, nil
}
