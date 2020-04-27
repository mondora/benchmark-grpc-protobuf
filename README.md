### gRPC+Protobuf or JSON+HTTP?

This repository contains 2 equal APIs: gRPC using Protobuf and JSON over HTTP. The goal is to run benchmarks for 2 approaches and compare them. APIs have 1 endpoint to create user, containing validation of request. Request, validation and response are the same in 2 packages, so we're benchmarking only mechanism itself. Benchmarks also include response parsing.

### Requirements

 - Go 1.14.2

### Run tests

Run benchmarks:
```
GO111MODULE=on go test -bench=. -benchmem
go test -bench=.
```

### Results

```
goos: darwin
goarch: amd64
BenchmarkGRPCProtobuf-8                    11379            103889 ns/op            9467 B/op        192 allocs/op
BenchmarkGzippedGRPCProtobuf-8              5190            228779 ns/op           19147 B/op        209 allocs/op
BenchmarkHTTPJSON-8                        10000            107691 ns/op            8757 B/op        115 allocs/op

PASS
ok      benchmark-grpc-protobuf-vs-http-json  7.370s
```

They are almost the same, HTTP+JSON is a bit faster and has less allocs/op.

### CPU usage comparison

This will create an executable `benchmark-grpc-protobuf-vs-http-json.test` and the profile information will be stored in `grpcprotobuf.cpu` and `httpjson.cpu`:

```
GO111MODULE=on go test -bench=BenchmarkGRPCProtobuf -cpuprofile=grpcprotobuf.cpu
GO111MODULE=on go test -bench=BenchmarkHTTPJSON -cpuprofile=httpjson.cpu
```

Check CPU usage per approach using:

```
go tool pprof grpcprotobuf.cpu
go tool pprof httpjson.cpu
```

My results show that Protobuf consumes less ressources, around **30% less**.

### gRPC definition

 - Install [Go](https://golang.org/dl/)
 - Install [Protocol Buffers](https://github.com/google/protobuf/releases)
 - Install protoc plugin: `go get github.com/golang/protobuf/proto github.com/golang/protobuf/protoc-gen-go`

```
protoc --go_out=plugins=grpc:. proto/api.proto
```
