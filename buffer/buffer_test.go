package buffer

import (
	"fmt"
	"sync"
	"testing"
	"time"
)

var buf = make(chan int,10)

func TestHello(t *testing.T) {
    var wg sync.WaitGroup
    wg.Add(2)

	fmt.Println("hello testing go routines and channels")

	go func() {
        defer wg.Done()
		for i := 1; i <=5; i++ {
			buf <- i
			fmt.Println("sent ", i)
		}
        close(buf)
        
	}()

	go func() {
        time.Sleep(2 *time.Second)
        defer wg.Done()
		fmt.Println("consuming the data")
        for data := range buf{
        time.Sleep(2 *time.Second)
		fmt.Println("data recieved",data)
        }
	}()


    wg.Wait()

}
