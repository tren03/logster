package buffer

import (
	"bytes"
	"encoding/gob"
	"fmt"
	"io"
	"log"

	"github.com/tren03/logster/global"
)

var Buf bytes.Buffer
// everytime we create a instace of this encoded, type defenition is put in, which causes duplication during decode, so we only initialize it 1 
var	enc = gob.NewEncoder(&Buf)

// Writes bytes to the Buffer
func EncodeData(logRecieved global.EventLog) {
	err := enc.Encode(logRecieved)
	if err != nil {
		log.Println("error in encoding struct to gob", err)
	}
	fmt.Println("encoded buf ", Buf)

}

// Decodes bytes from the buffer
func DecodeData() {
	var temp global.EventLog

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
		log.Println("decoded data ", temp)
	}
}
