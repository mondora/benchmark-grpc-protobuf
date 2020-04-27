package benchmarks

import (
	grpcprotobuf "benchmark-grpc-protobuf/grpc-protobuf"
	"benchmark-grpc-protobuf/grpc-protobuf/usertest"
	"google.golang.org/grpc/encoding/gzip"
	"testing"
	"time"

	"golang.org/x/net/context"
	g "google.golang.org/grpc"
)

func init() {
	go grpcprotobuf.Start()
	time.Sleep(time.Second)
}

func BenchmarkGRPCProtobuf(b *testing.B) {
	conn, err := g.Dial("127.0.0.1:60000", g.WithInsecure())
	if err != nil {
		b.Fatalf("grpc connection failed: %v", err)
	}

	client := usertest.NewAPIClient(conn)

	for n := 0; n < b.N; n++ {
		doGRPC(client, b)
	}
}

func BenchmarkGzippedGRPCProtobuf(b *testing.B) {
	conn, err := g.Dial("127.0.0.1:60000", g.WithInsecure())
	if err != nil {
		b.Fatalf("grpc connection failed: %v", err)
	}

	client := usertest.NewAPIClient(conn)

	for n := 0; n < b.N; n++ {
		doGzippedGRPC(client, b)
	}
}

func doGRPC(client usertest.APIClient, b *testing.B) {
	resp, err := client.CreateUser(context.Background(), &usertest.User{
		Email:    "foo@bar.com",
		Name:     "Bench",
		Password: "bench",
	})

	if err != nil {
		b.Fatalf("grpc request failed: %v", err)
	}

	if resp == nil || resp.Code != 200 || resp.User == nil || resp.User.Id != "1000000" {
		b.Fatalf("grpc response is wrong: %v", resp)
	}
}

func doGzippedGRPC(client usertest.APIClient, b *testing.B) {
	resp, err := client.CreateUser(context.Background(), &usertest.User{
		Email:    "foo@bar.com",
		Name:     "Bench",
		Password: "bench",
	},
	g.UseCompressor(gzip.Name))

	if err != nil {
		b.Fatalf("grpc request failed: %v", err)
	}

	if resp == nil || resp.Code != 200 || resp.User == nil || resp.User.Id != "1000000" {
		b.Fatalf("grpc response is wrong: %v", resp)
	}
}
