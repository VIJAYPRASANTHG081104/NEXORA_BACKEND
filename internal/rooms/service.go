package rooms

type RoomServiceInterface interface{   
    createTokens()
}

type RoomServiceStruct struct{}

func CreateRoomService() *RoomServiceStruct {
    return &RoomServiceStruct{} 
}



