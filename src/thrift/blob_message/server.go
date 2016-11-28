package main

import (
	//"bytes"
	"fmt"
	"git.apache.org/thrift.git/lib/go/thrift"
	"thrift/blob_message/gen-go/blob"
)

func main() {
	//  fmt.Println("vim-go")

	var protocolFactory thrift.TProtocolFactory
	protocolFactory = thrift.NewTBinaryProtocolFactoryDefault()

	var transportFactory thrift.TTransportFactory
	transportFactory = thrift.NewTBufferedTransportFactory(8192)

	var addr string
	addr = "localhost:9000"

	if err := runServer(transportFactory, protocolFactory, addr); err != nil {
		fmt.Println("error running client:", err)
	}
}

type BlobHandler struct {
	log map[int]*blob.Message
}

func NewBlobHandler() *BlobHandler {
	return &BlobHandler{log: make(map[int]*blob.Message)}
}

func (p *BlobHandler) Request(msg *blob.Message) (ret string, err error) {
	fmt.Print(msg.PartitionID, ", ", msg.BlobID, "\n")
	ret_str := "ret_str"
	return ret_str, nil
}

func runServer(transportFactory thrift.TTransportFactory, protocolFactory thrift.TProtocolFactory, addr string) error {

	var transport thrift.TServerTransport
	var err error

	transport, err = thrift.NewTServerSocket(addr)

	if err != nil {
		return err
	}

	fmt.Printf("%T\n", transport)
	//handler := NewPingHandler()
	handler := NewBlobHandler()
	processor := blob.NewBlobSvcProcessor(handler)
	server := thrift.NewTSimpleServer4(processor, transport, transportFactory, protocolFactory)

	fmt.Println("Starting the simple server... on ", addr)
	return server.Serve()
}
