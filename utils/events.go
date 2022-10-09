package utils
import (
	pb "github.com/kjais1720/grpc-go-server/protobuf"
	types "github.com/kjais1720/grpc-go-server/types"
)


func ToEvent(data types.Event) *pb.Event {
  return &pb.Event {
    Id: data.Id,
		Date: data.Date,
		Title: data.Title,
		Time: data.Time,
	}
}

func FromEvent(event *pb.Event) types.Event {
  return types.Event {
		Id: event.GetId(),
		Date: event.GetDate(),
		Title: event.GetTitle(),
		Time: event.GetTime(),
  }
}

func FromEventsArray(eventsArr []*pb.Event) []types.Event{
	newEventsArray := []types.Event{};
	for _, event := range eventsArr {
		newEventsArray = append(newEventsArray, FromEvent(event))
	}
	return newEventsArray
}

func ToEventsArray(eventsArr []types.Event) []*pb.Event{
	newEventsArray := []*pb.Event{};
	for _, event := range eventsArr {
		newEventsArray = append(newEventsArray, ToEvent(event))
	}
	return newEventsArray
}
