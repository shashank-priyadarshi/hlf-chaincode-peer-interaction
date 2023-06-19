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
	"context"
	"encoding/json"
	"fmt"
	"net"
	"os"

	pb "mock-grpc-peer/pb"

	log "github.com/sirupsen/logrus"
	"google.golang.org/grpc"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/types/known/timestamppb"
)

type ChaincodeSupport struct {
	pb.ChaincodeSupportServer
}

var counter int = 0
var getCounter int = 0

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
	for {
		msg, err := x.Recv()
		switch msg.Type {
		case pb.ChaincodeMessage_REGISTER:
			fmt.Println("received message from chaincode: ")
			fmt.Println(msg, err)
			x.Send(&pb.ChaincodeMessage{
				Type: pb.ChaincodeMessage_REGISTERED,
			})
			fmt.Println("sent message to chaincode: ")
			fmt.Println(msg, err)
			x.Send(&pb.ChaincodeMessage{
				Type: pb.ChaincodeMessage_READY,
			})
			fmt.Println("sent message to chaincode: ")
			fmt.Println(msg, err)
			x.Send(&pb.ChaincodeMessage{
				Type: pb.ChaincodeMessage_INIT,
			})
		default:
			fmt.Println("received message from chaincode: ")
			fmt.Println(msg, err)
			// dummy, _ := json.Marshal(struct {
			// 	Args        [][]byte          `protobuf:"bytes,1,rep,name=args,proto3" json:"args,omitempty"`
			// 	Decorations map[string][]byte `protobuf:"bytes,2,rep,name=decorations,proto3" json:"decorations,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
			// 	// is_init is used for the application to signal that an invocation is to be routed
			// 	// to the legacy 'Init' function for compatibility with chaincodes which handled
			// 	// Init in the old way.  New applications should manage their initialized state
			// 	// themselves.
			// 	IsInit bool `protobuf:"varint,3,opt,name=is_init,json=isInit,proto3" json:"is_init,omitempty"`
			// }{
			// 	IsInit:      false,
			// 	Args:        [][]byte{[]byte("init"), []byte("a"), []byte("100"), []byte("b"), []byte("200")},
			// 	Decorations: map[string][]byte{"event": []byte("event")},
			// })
			var temp []byte
			var txID string
			ctx, _ := json.Marshal(context.Background())
			if counter < 10 {
				temp, _ = proto.Marshal(&pb.ChaincodeInput{
					Decorations: map[string][]byte{"event": []byte("event")},
					IsInit:      false,
					Args:        [][]byte{[]byte("PutState"), ctx, []byte(fmt.Sprintf("shashank%v", counter)), []byte(fmt.Sprintf("beautiful%v", counter))},
				})
				// channelID = fmt.Sprintf("mychannel%v", counter)
				txID = fmt.Sprintf("txid%v", counter)
			} else {
				temp, _ = proto.Marshal(&pb.ChaincodeInput{
					Decorations: map[string][]byte{"event": []byte("event")},
					IsInit:      false,
					Args:        [][]byte{[]byte("GetState"), ctx, []byte(fmt.Sprintf("shashank%v", getCounter))},
				})
				// channelID = fmt.Sprintf("mychannel%v", getCounter)
				txID = fmt.Sprintf("txid%v", getCounter)
				if getCounter == 10 {
					os.Exit(1)
				}
				getCounter++
			}
			x.Send(&pb.ChaincodeMessage{
				Type:           pb.ChaincodeMessage_TRANSACTION,
				ChannelId:      "mychannel",
				Txid:           txID,
				Payload:        temp,
				Proposal:       []byte{},
				ChaincodeEvent: &pb.ChaincodeEvent{},
				Timestamp:      &timestamppb.Timestamp{},
			})
			counter++
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
