package buffer

import (
	"encoding/json"
	"fmt"
	"log"
	"sync"

	"github.com/tren03/logster/azureblob"
	"github.com/tren03/logster/global"
)

// var Buf bytes.Buffer
//
// // everytime we create a instace of this encoded, type defenition is put in, which causes duplication during decode, so we only initialize it 1
// var enc = gob.NewEncoder(&Buf)
// var typeInfo bytes.Buffer
// var flag = false
//
// // Writes bytes to the Buffer
//
//	func EncodeData(logRecieved global.EventLog) {
//		if flag == false {
//			encType := gob.NewEncoder(&typeInfo)
//			err := encType.Encode(logRecieved)
//			if err != nil {
//				log.Println("error in encoding struct to gob", err)
//			}
//			flag = true
//		} else {
//			err := enc.Encode(logRecieved)
//			if err != nil {
//				log.Println("error in encoding struct to gob", err)
//			}
//			fmt.Println("encoded buf ", Buf)
//
//			global.DATA_SENT += len(Buf.Bytes())
//
//			if Buf.Len() > 200 {
//				// file write
//				UploadData()
//			}
//
//		}
//	}
//
//	func UploadData() {
//		//	file, err := os.OpenFile("./logs.txt", os.O_RDWR|os.O_APPEND, 0644)
//		//	defer file.Close()
//		//	if err != nil {
//		//		log.Println("error opening file", err)
//		//	}
//		//	decodedData := DecodeData()
//		//	time.Sleep(3 * time.Second)
//		//	_, err = file.Write([]byte(decodedData))
//		//	if err != nil {
//		//		log.Println("error writing to file", err)
//		//	}
//		decodedData := DecodeData()
//		azureblob.UploadToBlob(decodedData)
//	   typeInfo.Reset()
//		Buf.Reset()
//
// }
//
// // Decodes bytes from the buffer
//
//	func DecodeData() string {
//		var temp global.EventLog
//		var logArray string
//
//		// to reset the buff pointer to the start of buffer to read properly
//	   combinedBuf := append(typeInfo.Bytes(),Buf.Bytes()...)
//		dec := gob.NewDecoder(bytes.NewReader(combinedBuf))
//		for {
//			err := dec.Decode(&temp)
//			if err != nil {
//				if err == io.EOF {
//					log.Println("reached end of file", err)
//					break
//				}
//				log.Println("error in decoding ", err)
//
//			}
//			temp_json, err := json.Marshal(temp)
//			if err != nil {
//				log.Println("issue in doing json", err)
//			}
//
//			logArray = logArray + string(temp_json) + "\n"
//
//		}
//
//		return logArray
//	}

//type Buf interface {
//	Lock()
//	Unlock()
//	ResetBuffer()
//	GetBuffer() []string
//	SignalFree()
//}

type ProducerBuffer struct {
	mu   sync.Mutex
	Buff []string
	cond *sync.Cond
}
type ConsumerBuffer struct {
	mu   sync.Mutex
	Buff []string
	cond *sync.Cond
}

var Prod = ProducerBuffer{cond: sync.NewCond(&sync.Mutex{})}
var Cons = ConsumerBuffer{cond: sync.NewCond(&sync.Mutex{})}

//type Buf1 struct {
//	mu   sync.Mutex
//	Buf1 []string
//	cond *sync.Cond // Condition variable to signal when buffer is empty
//}
//type Buf2 struct {
//	mu   sync.Mutex
//	Buf2 []string
//	flag bool
//	cond *sync.Cond // Condition variable to signal when buffer is empty
//}
//
//// Implement the methods for Buf1
//func (b *Buf1) Lock() {
//	b.mu.Lock()
//}
//
//func (b *Buf1) Unlock() {
//	b.mu.Unlock()
//}
//
//func (b *Buf1) GetBuffer() []string {
//	return b.Buf1
//}
//
//func (b *Buf1) ResetBuffer() {
//	b.Buf1 = []string{}
//}
//
//// Implement the methods for Buf2
//func (b *Buf2) Lock() {
//	b.mu.Lock()
//}
//
//func (b *Buf2) Unlock() {
//	b.mu.Unlock()
//}
//func (b *Buf2) SignalFree() {
//	b.cond.L.Lock()
//	b.cond.Signal()
//	b.cond.L.Unlock()
//}
//func (b *Buf1) SignalFree() {
//	b.cond.L.Lock()
//	b.cond.Signal()
//	b.cond.L.Unlock()
//}
//
//func (b *Buf2) GetBuffer() []string {
//	return b.Buf2
//}
//
//func (b *Buf2) ResetBuffer() {
//	b.Buf2 = []string{}
//}

var MAX_BUF = 10 // 1MB
//var B1 = Buf1{cond: sync.NewCond(&sync.Mutex{})}
//var B2 = Buf2{cond: sync.NewCond(&sync.Mutex{})}
//var activebuf = 1
//var activeBufMutex sync.Mutex // Mutex to protect access to `activebuf`

func EncodeBigData(logRecieved global.EventLog) {
	temp_json, err := json.Marshal(logRecieved)
	if err != nil {
		log.Println("issue in doing json", err)
	}
	logData := string(temp_json)

	Prod.mu.Lock()
	defer Prod.mu.Unlock()

	Prod.Buff = append(Prod.Buff, logData)
	log.Println("Added data to producer buffer:", len(Prod.Buff))

	if len(Prod.Buff) >= MAX_BUF {
		Prod.cond.Signal()
	}

}

func Consumer() {
	for {
		Prod.mu.Lock()
		if len(Prod.Buff) == 0 {
			Prod.cond.Wait()
		}

		Cons.mu.Lock()
		Cons.Buff = append(Cons.Buff, Prod.Buff...)
		Prod.Buff = []string{}
		log.Println("Moved data to consumer buffer for processing.")
		Cons.mu.Unlock()
		Prod.mu.Unlock()
		UploadData(&Cons)
	}

}

//func EncodeData(logRecieved global.EventLog) {
//
//	temp_json, err := json.Marshal(logRecieved)
//	if err != nil {
//		log.Println("issue in doing json", err)
//	}
//	logData := string(temp_json)
//	activeBufMutex.Lock()
//	a_buf := activebuf
//	activeBufMutex.Unlock()
//
//	if a_buf == 1 {
//		B1.Buf1 = append(B1.Buf1, logData)
//		log.Println("wrote to buf1")
//
//		size := 0
//		for _, str := range B1.Buf1 {
//			size += len(str)
//		}
//		log.Println("SIZE OF BUF1 : ", size)
//		log.Println("ITEMS IN BUF1 : ", len(B1.Buf1))
//		if size > MAX_BUF {
//
//			activeBufMutex.Lock()
//
//			B2.cond.L.Lock()
//			for len(B2.GetBuffer()) != 0 {
//				B2.cond.Wait() // Wait until B2 becomes empty
//			}
//			B2.cond.L.Unlock()
//
//			activebuf = 2 // Switch to Buf1 after uploading
//			activeBufMutex.Unlock()
//			go func() {
//				B1.Lock()         // Lock buffer before uploading
//				defer B1.Unlock() // Unlock after uploading is done
//				UploadData(&B1)
//			}()
//			//	activebuf = 2
//			//	UploadData(&B1)
//		}
//	} else {
//		B2.Buf2 = append(B2.Buf2, logData)
//		log.Println("wrote to buf2")
//		size := 0
//		for _, str := range B2.Buf2 {
//			size += len(str)
//		}
//		log.Println("SIZE OF BUF2: ", size)
//		log.Println("ITEMS IN BUF2 : ", len(B2.Buf2))
//		if size > MAX_BUF {
//			activeBufMutex.Lock()
//
//			B1.cond.L.Lock()
//			for len(B1.GetBuffer()) != 0 {
//				B1.cond.Wait() // Wait until B2 becomes empty
//			}
//			B1.cond.L.Unlock()
//
//			activebuf = 1 // Switch to Buf1 after uploading
//			activeBufMutex.Unlock()
//			go func() {
//				B2.Lock()         // Lock buffer before uploading
//				defer B2.Unlock() // Unlock after uploading is done
//				UploadData(&B2)
//			}()
//			//	activebuf = 1
//			//	UploadData(&B2)
//		}
//	}
//
//}

func UploadData(buf *ConsumerBuffer) {
	fmt.Println("CONSUMING DATA")
	buf.mu.Lock()
	defer buf.mu.Unlock()
	var final_string string
	for _, str := range buf.Buff {
		final_string += str + "\n"
	}
	azureblob.UploadToBlob(final_string)
	buf.Buff = []string{}
	log.Println("Consumer buffer processed and cleared.")

}

func StartConsumer() {
	go Consumer()
}

// Cleanup operation when server shuts down
func UploadCleanup() {
	Cons.mu.Lock()
	defer Cons.mu.Unlock()

	if len(Cons.Buff) > 0 {
		UploadData(&Cons)
	}
}

//func UploadCleanup() {
//	if len(B1.GetBuffer()) != 0 {
//		UploadData(&B1)
//	}
//	if len(B2.GetBuffer()) != 0 {
//		UploadData(&B2)
//	}
//}
