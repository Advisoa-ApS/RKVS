package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"net"
	"net/http"

	pb "github.com/Advisoa-ApS/rkvs/proto/gen"
	badger "github.com/dgraph-io/badger/v4"
	"google.golang.org/grpc"
)

const (
	protoFile = "/etc/rkvs/proto/rkvs.proto"
)

var (
	serverIp   = flag.String("ip", "127.0.0.1", "IP address for the gRPC server to listen on")
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

func (s *RkvsServer) Get(ctx context.Context, k *pb.Key) (*pb.Item, error) {
	var key []byte
	var val []byte
	err := s.db.View(func(txn *badger.Txn) error {
		item, err := txn.Get([]byte(k.Key))
		if err == badger.ErrKeyNotFound {
			return nil // No error, just an empty value
		}
		if err != nil {
			return err
		}
		key = item.KeyCopy(nil)
		val, err = item.ValueCopy(nil)
		return err
	})
	if err != nil {
		return nil, err
	}

	return &pb.Item{
		Key:   string(key),
		Value: string(val),
	}, nil
}

func (s *RkvsServer) GetAll(ctx context.Context, p *pb.Prefix) (*pb.Items, error) {
	items := make(map[string]string)

	err := s.db.View(func(txn *badger.Txn) error {
		opts := badger.DefaultIteratorOptions
		opts.PrefetchValues = true
		it := txn.NewIterator(opts)
		defer it.Close()

		prefix := []byte(p.Prefix)
		for it.Seek(prefix); it.ValidForPrefix(prefix); it.Next() {
			item := it.Item()
			val, err := item.ValueCopy(nil)
			if err != nil {
				return err
			}
			items[string(item.KeyCopy(nil))] = string(val)
		}
		return nil
	})
	if err != nil {
		return nil, err
	}

	return &pb.Items{Items: items}, nil
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
	opts := badger.DefaultOptions("/var/lib/rkvs/data")
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
