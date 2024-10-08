package global

type Event struct {
	EventName string `json:"event_name"`
}

type EventLog struct{
    UnixTimeStamp int64
    Data Event
}

