package types

import (
	pb "github.com/kjais1720/grpc-go-server/protobuf"
)

type Server struct {
  pb.UserServiceServer
}


type Event struct {
	Id string
	Date string
	Title string
	Time string
}

type UserDetails struct {
  Uid int32
  Name string
	Email string
	Events []Event
	Password string
}
