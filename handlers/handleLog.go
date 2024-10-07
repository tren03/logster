package handlers

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"time"

	"github.com/tren03/logster/azureblob"
	"github.com/tren03/logster/buffer"
	"github.com/tren03/logster/global"
)

//	t := time.Now().Unix()
//	sample_event := global.Event{EventName: "login"}
//	sample_log := global.EventLog{UnixTimeStamp: t, EventName: sample_event}
//	fmt.Println(sample_log)

func HandleLog(w http.ResponseWriter, r *http.Request) {
	t := time.Now().Unix()
	var event global.Event
	bodyBytes, err := io.ReadAll(r.Body)
	if err != nil {
		log.Println("error reading req")
	}
	err = json.Unmarshal(bodyBytes, &event)
	if err != nil {
		log.Println("error in unmarshalling the json request")
	}

	eventLog := global.EventLog{UnixTimeStamp: t, EventName: event}
	fmt.Println(eventLog)
   // buffer.EncodeBigData(eventLog)
    fmt.Println("sending data to buffer")
    buffer.PutData(eventLog)
}


func HandleRoot(w http.ResponseWriter, r *http.Request) {
    azureblob.GetBlobInfo()
} 

func HandleUpload(w http.ResponseWriter, r *http.Request) {
    //buffer.UploadCleanup()
} 


