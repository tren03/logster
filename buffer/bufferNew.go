package buffer

import (
	"encoding/json"
	"fmt"
	"log"

	"github.com/tren03/logster/azureblob"
	"github.com/tren03/logster/global"
)

// CHAN SIZE IS APPROX 10MB

var B = make(chan string, 177725)

func PutData(logRecieved global.EventLog) {
    fmt.Println("logRecieved in put data")
	temp_json, err := json.Marshal(logRecieved)
	if err != nil {
		log.Println("issue in doing json", err)
	}
	logData := string(temp_json)
	B <- logData
}

// upload to db
func SendData() {
    for {
        var final_string string
        for i := 0; i < 177725; i++ {  // Process a batch 
            select {
            case str := <-B:
                final_string += str + "\n"
                log.Println("size of one entry,",len(final_string))
            default:
                // If no data is immediately available, break out
                break
            }
        }
        if final_string != "" {
            azureblob.UploadToBlob(final_string)
        }
    }
}

// spins up go routing to recieve data from the channel which i call in main 
func StartSender(){
    go SendData()
}

func CloseChan(){
    close(B)
}
