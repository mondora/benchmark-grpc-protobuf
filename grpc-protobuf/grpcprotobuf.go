package grpcprotobuf

import (
	"benchmark-grpc-protobuf/grpc-protobuf/usertest"
	"errors"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"io"
	"log"
	"net"
	"net/mail"
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

// CreateUsers handler
func (s *Server) CreateUsers(userStream usertest.API_CreateUsersServer) error {
	var count uint64 = 0
	for {
		user, err := userStream.Recv()
		if err == io.EOF {
			return userStream.SendAndClose(&usertest.ResponseManyUsers{
				Message: "OK",
				Code:    200,
				Count:   count,
			})
		}
		if err != nil {
			return err
		}
		err = validate(user)
		if err != nil {
			return userStream.SendAndClose(&usertest.ResponseManyUsers{
				Code:    500,
				Message: err.Error(),
				Count:   count,
			})
		}
		count++
	}
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
