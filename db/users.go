package db

import (
	types "github.com/kjais1720/grpc-go-server/types"
)
var Users = []types.UserDetails{
  {
    Uid: 1,
    Name: "Krituraj Anand",
		Email: "krituraj@krituraj.com",
		Password:"Krituraj@123",
		Events : []types.Event{
			{
				Title:"Eat",
				Time:"12AM",
				Id:"swded",
				Date:"08/17/2022",
			},
			{
				Title:"Sleep",
				Time:"11AM",
				Id:"asdfg",
				Date:"08/17/2022",
			},
			{
				Title:"Code",
				Time:"11AM",
				Id:"qwerty",
				Date:"08/17/2022",
			},

		},
  },
	{
    Uid: 2,
    Name: "Satyam Raj",
		Email: "satyam@satyam.com",
		Password:"satyam@123",
		Events : []types.Event{
			{
				Title:"Trek",
				Time:"12AM",
				Id:"swded",
				Date:"08/17/2022",
			},
			{
				Title:"Dive",
				Time:"11AM",
				Id:"asdfg",
				Date:"08/17/2022",
			},
			{
				Title:"Climb",
				Time:"7PM",
				Id:"qwerty",
				Date:"08/17/2022",
			},

		},
  },
}
