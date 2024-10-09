package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"github.com/tren03/logster/azureblob"
	"github.com/tren03/logster/buffer"
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
//
//	func genNos(stream chan<- []byte, count int) {
//		for i := 1; i <= count; i++ {
//	       strNum := strconv.Itoa(i)
//			stream <- []byte(strNum)
//		}
//
//		close(stream)
//
// }
//
//		file, err := os.OpenFile("./logs.txt", os.O_RDWR|os.O_APPEND, 0644)
//		if err != nil {
//			log.Println("error opening the file")
//		}
//
//		stream := make(chan []byte)
//		go genNos(stream, 1000)
//
//		b := make([]byte,0,100)
//
//		for num := range stream {
//			if len(b) == 100 {
//				// write to file with 1 second
//				time.Sleep(1 * time.Second)
//				_, err = file.Write(append(b,'\n'))
//				if err != nil {
//					log.Println("some error in write", err)
//				}
//	           b = b[:0]
//			} else {
//				b = append(b,num...)
//			}
//
//	       fmt.Println("finished one iteration")
//		}
//
//		if len(b) != 0 {
//			time.Sleep(1 * time.Second)
//			_, err = file.Write(b)
//			if err != nil {
//				log.Println("some error in write", err)
//			}
//		}
//
//		defer file.Close()
func cleanup() {
	fmt.Println("all done :)")
	buffer.CloseChan()

	//	ROOT_URL := fmt.Sprintf("http://localhost:8080/")
	//	fmt.Println("hitting root for verification of req")
	//	_, err := http.Get(ROOT_URL)
	//	if err != nil {
	//		fmt.Println("err sending get", err)
	//	}

}
func main() {
	sigChannel := make(chan os.Signal, 1)
	signal.Notify(sigChannel, syscall.SIGINT, syscall.SIGTERM)

	// Start a goroutine to handle cleanup when an OS signal is received
	go func() {
		sig := <-sigChannel
		fmt.Println("Received signal:", sig)
		cleanup()  // Call cleanup on termination
		os.Exit(0) // Exit the program after cleanup
	}()

	defer cleanup()
	fmt.Println("server started at port 8080")
	fmt.Println("creating container")
	azureblob.CreateContainer()
	http.HandleFunc("POST /log", handlers.HandleLog)
	//	http.HandleFunc("/up", handlers.HandleUpload)
	http.HandleFunc("/", handlers.HandleRoot)

	// starting 5 go routines to simulataneously run consume and upload data
    // lost some data when i had 10 go routines
	for i := 0; i < 8; i++ {
		buffer.StartSender(i)
	}

	log.Fatal(http.ListenAndServe(":8080", nil))
}
