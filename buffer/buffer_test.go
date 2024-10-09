package buffer

import (
	"fmt"
	"testing"
	"time"
)

var buf = make(chan int, 5)

func TestHello(t *testing.T) {

	// need a chan that has data and 3 consumers eat from that
	// spawning the 3 consumers
	go func() {
		for {
			for num := range buf {
				time.Sleep(1 * time.Second)
				fmt.Println("consuming from consumer 1 : ", num)
			}
		}
	}()

	go func() {
		for {
			for num := range buf {
				time.Sleep(1 * time.Second)
				fmt.Println("consuming from consumer 2 : ", num)
			}

		}
	}()
	go func() {
		for {
			for num := range buf {
				time.Sleep(1 * time.Second)
				fmt.Println("consuming from consumer 3 : ", num)
			}
		}
	}()

	// producer

	for i := 1; i <= 10; i++ {
		buf <- i
	}

    time.Sleep(10* time.Second)

}
