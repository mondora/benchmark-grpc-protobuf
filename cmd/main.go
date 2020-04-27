package main

import (
	grpcprotobuf "benchmark-grpc-protobuf/grpc-protobuf"
	"log"
)

func main() {
	log.Println("start grpc-srv")
	grpcprotobuf.Start()
}
