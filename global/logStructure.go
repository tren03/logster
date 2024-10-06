package global

type Event struct {
	EventName string `json:"event_name"`
}

type EventLog struct{
    UnixTimeStamp int64
    EventName Event
}

var DATA_SENT int
var DATA_RECIEVED int
