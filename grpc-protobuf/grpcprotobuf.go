package grpcprotobuf

import (
	"errors"
	"log"
	"net"
	"net/mail"

	"benchmark-grpc-protobuf/grpc-protobuf/usertest"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
)

// Start entrypoint
func Start() {
	lis, _ := net.Listen("tcp", ":60000")

	srv := grpc.NewServer()
	usertest.RegisterAPIServer(srv, &Server{})
	log.Println(srv.Serve(lis))
}

// Server type
type Server struct{}

// CreateUser handler
func (s *Server) CreateUser(ctx context.Context, in *usertest.User) (*usertest.Response, error) {
	validationErr := validate(in)
	if validationErr != nil {
		return &usertest.Response{
			Code:    500,
			Message: validationErr.Error(),
		}, validationErr
	}

	in.Id = "1000000"
	return &usertest.Response{
		Code:    200,
		Message: "OK",
		User:    in,
	}, nil
}

func validate(in *usertest.User) error {
	_, err := mail.ParseAddress(in.Email)
	if err != nil {
		return err
	}

	if len(in.Name) < 4 {
		return errors.New("Name is too short")
	}

	if len(in.Password) < 4 {
		return errors.New("Password is too weak")
	}

	return nil
}
