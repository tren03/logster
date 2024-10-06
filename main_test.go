package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"testing"

	"github.com/tren03/logster/global"
)

func Test25(t *testing.T) {
    URL := fmt.Sprintf("http://localhost:8080/log")

	for i := 0; i < 25; i++ {
		log := global.Event{EventName: "login"}
		log_json, err := json.Marshal(log)
		if err != nil {
           t.Fail() 
			fmt.Println("err conv to json", log_json)
		}
		temp_buf := bytes.NewReader(log_json)
		_, err = http.Post(URL, "application/json", temp_buf)
        if err!=nil{
            t.Fail()
            fmt.Println("err sending req",err)
        }
	}
    fmt.Println("Test complete, sent 25 requests.")

}
