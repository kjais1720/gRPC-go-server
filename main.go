package main

import (
	"context"
	"errors"
	"fmt"
	"log"
	"math"
	"net"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/keepalive"

	pb "github.com/kjais1720/grpc-go-server/protobuf"
	db "github.com/kjais1720/grpc-go-server/db"
	utils "github.com/kjais1720/grpc-go-server/utils"
)

type server struct {
  pb.UserServiceServer
}

func (s *server) Signup(ctx context.Context, req *pb.SignupRequest) (*pb.SignupResponse, error) {
  fmt.Println("Creating User")
  user := req.GetUser()

  data := utils.FromUser(user)

  db.Users = append(db.Users, data)

  return &pb.SignupResponse {
    User: utils.ToUser(data),
  }, nil
}

func (s *server) Login(ctx context.Context, req *pb.LoginRequest) (*pb.LoginResponse, error) {
  fmt.Println("Logging User in")
  email := req.GetEmail()
	password := req.GetPassword();
	for _, user := range db.Users{
		if user.Email == email && user.Password == password {
			return &pb.LoginResponse{
				User: utils.ToUser(user),
			},nil;
		}	
	}
  return nil, errors.New("invalid credentials")
}

func (s *server) CreateEvent(ctx context.Context, req *pb.CreateEventRequest) (*pb.CreateEventResponse, error) {
	fmt.Println("Creating Event");
	uid := req.GetUid();
	newEvent := req.GetEvent(); 
	for i, user := range db.Users {
		if user.Uid == uid {
			db.Users[i].Events = append(user.Events, utils.FromEvent(newEvent))
		}
	}
	return &pb.CreateEventResponse{
		Event :newEvent,
	}, nil
}


func (s *server ) DeleteEvent(ctx context.Context, req *pb.DeleteEventRequest) (*pb.DeleteEventResponse, error){
	eventId := req.GetId();
	uid := req.GetUid();
	for _, user := range db.Users {
		updatedEventsList := user.Events[:]
		if user.Uid == uid {
			for i, event := range user.Events {
				if(event.Id == eventId){
					updatedEventsList = append(updatedEventsList[:i], updatedEventsList[i+1:]...)
				}
			}
		}
	}
	return &pb.DeleteEventResponse{
		Uid: uid,
		Id: eventId,
	},nil

}

func main() {
  lis, err := net.Listen("tcp", "127.0.0.1:5000")
  if err != nil {
    log.Fatalf("Failed to listen: %v", err)
  }

  opts := []grpc.ServerOption{
    grpc.MaxRecvMsgSize(math.MaxInt64),
    grpc.KeepaliveParams(
      keepalive.ServerParameters{
        Timeout: 5 * time.Second,
      },
    ),
  }

  s := grpc.NewServer(opts...)
  pb.RegisterUserServiceServer(s, &server{})

  fmt.Println("Starting server...")
  fmt.Printf("Hosting server on: %s\n", lis.Addr().String())
  if err := s.Serve(lis); err != nil {
    log.Fatalf("Failed to serve: %v", err)
  }
}