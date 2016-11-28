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

	if err := runClient(transportFactory, protocolFactory, addr); err != nil {
		fmt.Println("error running client:", err)
	}
}

func handleClient(client *blob.BlobSvcClient) (err error) {

	//partition_id := "abcd2"
	//blob_id := "wxyz1"

	//	msg_struct := blob.Message{"asdf", "asdf"}
	msg_struct := blob.NewMessage()
	msg_struct.PartitionID = "asdf1" //partition_id
	msg_struct.BlobID = "wxyz1"      //blob_id

	server_msg_rsp, err := client.Request(msg_struct)
	//fmt.Println("server_response: %s", server_msg_rsp)
	fmt.Print("[SERVER_RSP] ", server_msg_rsp, "\n")

	return nil
}

func runClient(transportFactory thrift.TTransportFactory, protocolFactory thrift.TProtocolFactory, addr string) error {
	var transport thrift.TTransport
	var err error
	transport, err = thrift.NewTSocket(addr)

	if err != nil {
		fmt.Println("Error opening socket:", err)
		return err
	}

	if transport == nil {
		return fmt.Errorf("Error opening socket, got nil transport. Is server available?")
	}
	transport = transportFactory.GetTransport(transport)

	if transport == nil {
		return fmt.Errorf("Error from transportFactory.GetTransport() , got nil transport. Is server available?")
	}

	err = transport.Open()
	if err != nil {
		return err
	}
	defer transport.Close()

	//return handleClient(tutorial.NewCalculatorClientFactory(transport, protoclFactor))
	return handleClient(blob.NewBlobSvcClientFactory(transport, protocolFactory))
}
