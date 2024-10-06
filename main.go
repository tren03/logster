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

// THIS IS TO LEARN READ BUFFERS - code reads data into buffer and displays it
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

// TO READ FROM STRING STREAM AND WRITE TO FILE
//	file, err := os.OpenFile("./logs.txt",os.O_RDWR|os.O_APPEND,0644)
//	if err != nil {
//		log.Println("error opening the file")
//	}
//
//	stream := strings.NewReader("this is the data stream to simulate streaming of data yay")
//	b := make([]byte, 5)
//
//	for {
//		bytesRead, err := stream.Read(b)
//		if err != nil {
//			if err == io.EOF {
//				fmt.Println("We reached end bro")
//                break
//			}
//		}
//
//
//        time.Sleep(1 * time.Second)
//        _,err = file.Write(b[:bytesRead])
//        if err!=nil{
//           log.Println("some write issue",err)
//
//        }
//        fmt.Println("one operation done")
//
//	}
//
//	defer file.Close()



// CREATES A STREAM OF DATA AND WRITES IT TO A FILE WITH SIMULATED DELAY
//func genNos(stream chan<- []byte, count int) {
//	for i := 1; i <= count; i++ {
//        strNum := strconv.Itoa(i)
//		stream <- []byte(strNum)
//	}
//
//	close(stream)
//
//}
//	file, err := os.OpenFile("./logs.txt", os.O_RDWR|os.O_APPEND, 0644)
//	if err != nil {
//		log.Println("error opening the file")
//	}
//
//	stream := make(chan []byte)
//	go genNos(stream, 1000)
//
//	b := make([]byte,0,100)
//
//	for num := range stream {
//		if len(b) == 100 {
//			// write to file with 1 second
//			time.Sleep(1 * time.Second)
//			_, err = file.Write(append(b,'\n'))
//			if err != nil {
//				log.Println("some error in write", err)
//			}
//            b = b[:0]
//		} else {
//			b = append(b,num...)
//		}
//
//        fmt.Println("finished one iteration")
//	}
//
//	if len(b) != 0 {
//		time.Sleep(1 * time.Second)
//		_, err = file.Write(b)
//		if err != nil {
//			log.Println("some error in write", err)
//		}
//	}
//
//	defer file.Close()
func main() {
	fmt.Println("server started at port 8080")
	http.HandleFunc("POST /log", handlers.HandleLog)
    http.HandleFunc("/", handlers.HandleRoot)
	log.Println("testing buffer")

//	file, err := os.OpenFile("./logs.txt", os.O_RDWR|os.O_APPEND, 0644)
//	if err != nil {
//		log.Println("error opening the file")
//	}
//	defer file.Close()

	log.Fatal(http.ListenAndServe(":8080", nil))
}
