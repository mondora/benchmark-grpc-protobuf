### gRPC+Protobuf or JSON+HTTP?

This repository contains 2 equal APIs: gRPC using Protobuf and JSON over HTTP.
The goal is to run benchmarks for 2 approaches and compare them.
APIs have 1 endpoint to create user, containing validation of request.
Request, validation and response are the same in 2 packages, so we're benchmarking only mechanism itself.
Benchmarks also include response parsing.

### Requirements
git remote add origin https://github.com/mondora/benchmark-grpc-protobuf.git
 - Go 1.14.2

### Run tests

Run benchmarks:
```
go test -bench=.
go test -bench=. -benchmem
```

### Results

```
goos: darwin
goarch: amd64
BenchmarkGRPCProtobuf-8                    11379            103889 ns/op            9467 B/op        192 allocs/op
BenchmarkGzippedGRPCProtobuf-8              5190            228779 ns/op           19147 B/op        209 allocs/op
BenchmarkHTTPJSON-8                        10000            107691 ns/op            8757 B/op        115 allocs/op

PASS
ok      benchmark-grpc-protobuf  7.370s
```

They are almost the same (gRPC and HTTP+JSON).
gZipped gRPC request is slower.

### CPU usage comparison

This will create an executable `benchmark-grpc-protobuf.test` and the profile information will be stored in `grpcprotobuf.cpu` (+gzipped) and `httpjson.cpu`:

```
go test -bench=BenchmarkGRPCProtobuf -cpuprofile=grpcprotobuf.cpu
go test -bench=BenchmarkGzippedGRPCProtobuf -cpuprofile=grpcgzippedprotobuf.cpu
go test -bench=BenchmarkHTTPJSON -cpuprofile=httpjson.cpu
```

Check CPU usage per approach using:

```
go tool pprof grpcprotobuf.cpu
go tool pprof grpcgzippedprotobuf.cpu
go tool pprof httpjson.cpu
```

### gRPC definition

 - Install [Protocol Buffers](https://github.com/google/protobuf/releases)
 - Install protoc plugin: `go get github.com/golang/protobuf/proto github.com/golang/protobuf/protoc-gen-go`

```
protoc --go_out=plugins=grpc:. proto/api.proto
```
