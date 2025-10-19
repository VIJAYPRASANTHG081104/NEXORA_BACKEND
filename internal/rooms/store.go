package rooms

import (
	"context"
	"os"
	"time"

	"cloud.google.com/go/auth"
	livekit "github.com/livekit/protocol/livekit"
	lksdk "github.com/livekit/server-sdk-go"
)

func (h *RoomServiceStruct) CreateTokens(room, identity string) string{
	at := auth.NewAccessToken(os.Getenv("LIVEKIT_API_KEY"), os.Getenv("LIVEKIT_API_SECRET"))
	grant := &auth.VideoGrant{
		RoomJoin: true,
		Room: room,
	}
	at.AddGrant(grant).
		SetIdentity(identity).
		SetValidFor(time.Hour)
	token, _ := at.ToJWT()
	return token
}

const host = "https://my.livekit.host";
var roomClient = lksdk.NewRoomServiceClient(host, "api-key", "secret-key")

func (h *RoomServiceStruct) CreateRooms(roomName string,noOfParticipants int, time int) error{
	room, error := roomClient.CreateRoom(context.Background(), &livekit.CreateRoomRequest{
		Name:            roomName,
		EmptyTimeout:    10 * 60, // 10 minutes
		MaxParticipants: 20,
	})
	if error != nil{
		return error
	}
	_ = room
	return nil
}

func (h *RoomServiceStruct) DeleteRooms(roomName string) {	
	_, _ = roomClient.DeleteRoom(context.Background(), &livekit.DeleteRoomRequest{
	Room: roomName,
	})
}

func ListOfParticipants(roomName string) (*livekit.ListParticipantsResponse, error){
	return roomClient.ListParticipants(context.Background(), &livekit.ListParticipantsRequest{
		Room: roomName,
	})
}	

func DetailsOfParticipant(roomName string, participantSid string) (*livekit.ParticipantInfo, error){
	return roomClient.GetParticipant(context.Background(), &livekit.RoomParticipantIdentity{
		Room: roomName,
		Identity: participantSid,
	})
}	

func UpdateParticipantRole(roomName string, participantSid string, role livekit.ParticipantInfo) (*livekit.ParticipantInfo, error){
	return roomClient.UpdateParticipant(context.Background(), &livekit.UpdateParticipantRequest{
		Room: roomName,
		Identity: participantSid,
		Permission: &livekit.ParticipantPermission{
			CanSubscribe: true,
			CanPublish: true,
			CanPublishData: true,
		},
	})
}		

func RemoveParticipant(roomName string, identity string) (*livekit.RemoveParticipantResponse, error){
	return roomClient.RemoveParticipant(context.Background(), &livekit.RoomParticipantIdentity{
		Room: roomName,
		Identity: identity,
	})
}

func MuteParticipant(roomName string, identity string) (*livekit.MuteRoomTrackRequest, error){
	_, err := roomClient.MutePublishedTrack(context.Background(), &livekit.MuteRoomTrackRequest{
		Room:     roomName,
		Identity: identity,
		TrackSid: "track_sid",
		Muted:    true,
	})
	return nil, err
}	