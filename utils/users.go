package utils
import (
	pb "github.com/kjais1720/grpc-go-server/protobuf"
	types "github.com/kjais1720/grpc-go-server/types"
)

func ToUser(data types.UserDetails) *pb.User {
  return &pb.User {
    Uid: data.Uid,
    Name: data.Name,
		Email: data.Email,
		Events: ToEventsArray(data.Events),
  }
}

func FromUser(user *pb.User) types.UserDetails {
  return types.UserDetails {
    Uid: user.GetUid(),
    Name: user.GetName(),
		Email: user.GetEmail(),
		Events: FromEventsArray(user.GetEvents()),
		Password: user.GetPassword(),
  }
}