package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/tren03/logster/handlers"
)

//	t := time.Now().Unix()
//	sample_event := global.Event{EventName: "login"}
//	sample_log := global.EventLog{UnixTimeStamp: t, EventName: sample_event}
//	fmt.Println(sample_log)

// THIS IS TO LEARN READ BUFFERS
//	file, err := os.Open("./logs.txt")
//	if err != nil {
//		log.Println("error opening the file")
//	}
//	b := make([]byte, 10)
//	for {
//		bytesRead, err := file.Read(b)
//		if err != nil {
//			if err.Error() == "EOF" {
//				log.Println("end of file reached")
//				break
//			} else {
//
//				log.Println("issue in buffer")
//			}
//		}
//        fmt.Println("return of read : ",bytesRead)
//        fmt.Println(string(b[:bytesRead]))
//        fmt.Println("content of buffer : ",b)
//	}
//
//	defer file.Close()

func main() {
	fmt.Println("server started at port 8080")
	http.HandleFunc("POST /log", handlers.HandleLog)
	log.Println("testing buffer")

//	file, err := os.Open("./logs.txt")
//	if err != nil {
//		log.Println("error opening the file")
//	}
//
//    stream := strings.NewReader("this is the data stream to simulate streanming of data yay")
//	b := make([]byte, 5)
//
//    
//	defer file.Close()

	log.Fatal(http.ListenAndServe(":8080", nil))
}
