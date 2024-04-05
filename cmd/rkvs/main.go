package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"net"
	"net/http"

	pb "github.com/Advisoa-ApS/rkvs/proto/gen"
	badger "github.com/dgraph-io/badger/v3"
	"google.golang.org/grpc"
)

const (
	protoFile = "../../proto/rkvs.proto"
)

var (
	serverIp   = flag.String("ip", "192.168.56.1", "IP address for the gRPC server to listen on")
	grpcPort   = flag.String("port", "9090", "Port for the gRPC server")
	httpPort   = flag.String("http_port", "9091", "Port for the HTTP server to serve the proto file")
	daemonMode = flag.Bool("daemon", false, "Run in daemon mode")
)

type RkvsServer struct {
	pb.UnimplementedRkvsServer
	db *badger.DB
}

func serveProtoFile(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, protoFile)
}

func NewRkvsServer(db *badger.DB) *RkvsServer {
	return &RkvsServer{db: db}
}

func (s *RkvsServer) Get(ctx context.Context, k *pb.Key) (*pb.Value, error) {
	var val []byte
	err := s.db.View(func(txn *badger.Txn) error {
		item, err := txn.Get([]byte(k.Key))
		if err != nil {
			return err
		}
		val, err = item.ValueCopy(nil)
		return err
	})
	if err != nil {
		return nil, nil
	}
	return &pb.Value{Value: string(val)}, nil
}

func (s *RkvsServer) ExecuteTransaction(ctx context.Context, req *pb.TransactionRequest) (*pb.Ack, error) {
	err := s.db.Update(func(txn *badger.Txn) error {
		for _, operation := range req.Operations {
			switch op := operation.OperationType.(type) {
			case *pb.Operation_Set:
				if err := txn.Set([]byte(op.Set.Key), []byte(op.Set.Value)); err != nil {
					return err
				}
			case *pb.Operation_Delete:
				if err := txn.Delete([]byte(op.Delete.Key)); err != nil {
					return err
				}
			}
		}
		return nil
	})

	return &pb.Ack{Success: err == nil}, err
}

func main() {

	flag.Parse()

	// Init BadgerDB
	opts := badger.DefaultOptions("../data")
	db, err := badger.Open(opts)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	// Start the gRPC server in a goroutine
	go func() {
		lis, err := net.Listen("tcp", fmt.Sprintf("%s:%s", *serverIp, *grpcPort))
		if err != nil {
			log.Fatalf("failed to listen: %v", err)
		}
		s := grpc.NewServer()
		pb.RegisterRkvsServer(s, NewRkvsServer(db))
		log.Printf("Serving gRPC on %s:%s", *serverIp, *grpcPort)
		if err := s.Serve(lis); err != nil {
			log.Fatalf("failed to serve: %v", err)
		}
	}()

	// The HTTP server for the .proto file remains unchanged
	http.HandleFunc("/", serveProtoFile)
	log.Printf("Serving proto file on HTTP %s:%s", *serverIp, *httpPort)
	if err := http.ListenAndServe(fmt.Sprintf("%s:%s", *serverIp, *httpPort), nil); err != nil {
		log.Fatalf("failed to serve HTTP: %v", err)
	}
}
