package buffer

import (
	"bytes"
	"encoding/gob"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"os"
	"time"

	"github.com/tren03/logster/global"
)

var Buf bytes.Buffer

// everytime we create a instace of this encoded, type defenition is put in, which causes duplication during decode, so we only initialize it 1
var enc = gob.NewEncoder(&Buf)

// Writes bytes to the Buffer
func EncodeData(logRecieved global.EventLog) {
	err := enc.Encode(logRecieved)
	if err != nil {
		log.Println("error in encoding struct to gob", err)
	}
	fmt.Println("encoded buf ", Buf)

	global.DATA_SENT += len(Buf.Bytes())

	if Buf.Len() > 200 {
        // file write
        UploadData()
	}

}

func UploadData() {
	file, err := os.OpenFile("./logs.txt", os.O_RDWR|os.O_APPEND, 0644)
	defer file.Close()
	if err != nil {
		log.Println("error opening file", err)
	}
	decodedData := DecodeData()
	time.Sleep(3 * time.Second)
	_, err = file.Write([]byte(decodedData))
	if err != nil {
		log.Println("error writing to file", err)
	}
	Buf.Reset()
	enc = gob.NewEncoder(&Buf)
}

// Decodes bytes from the buffer
func DecodeData() string {
	var temp global.EventLog
	var logArray string

	// to reset the buff pointer to the start of buffer to read properly
	dec := gob.NewDecoder(bytes.NewReader(Buf.Bytes()))
	for {
		err := dec.Decode(&temp)
		if err != nil {
			if err == io.EOF {
				log.Println("reached end of file", err)
				break
			}
			log.Println("error in decoding ", err)

		}
		temp_json, err := json.Marshal(temp)
		if err != nil {
			log.Println("issue in doing json", err)
		}

		logArray = logArray + string(temp_json) + "\n"

	}

	return logArray
}
