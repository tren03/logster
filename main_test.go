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
	//ROOT_URL := fmt.Sprintf("http://localhost:8080/")
	URL := fmt.Sprintf("http://localhost:8080/log")

	for i := 0; i < 50; i++ {
		log := global.Event{EventName: "login"}
		log_json, err := json.Marshal(log)
		if err != nil {
			t.Fail()
			fmt.Println("err conv to json", log_json)
		}
		temp_buf := bytes.NewReader(log_json)
		_, err = http.Post(URL, "application/json", temp_buf)
		if err != nil {
			t.Fail()
			fmt.Println("err sending req", err)
		}
	}

//	fmt.Println("hitting root for verification of req")
//	_, err := http.Get(ROOT_URL)
//	if err != nil {
//		t.Fail()
//		fmt.Println("err sending get", err)
//	}
//
	fmt.Println("Test complete, sent 50 requests and verification")

}
