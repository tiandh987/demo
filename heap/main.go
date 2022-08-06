package main

import (
	"fmt"
	"log"
	"net/http"
	_ "net/http/pprof"
	"time"
	"unsafe"
)

func main() {
	times := 100

	go func() {
		log.Println(http.ListenAndServe("0.0.0.0:8099", nil))
	}()

	time.Sleep(time.Second * 30)

	for i := 0; i < times; i++ {
		//temp := make(map[int]bool)
		temp := make([]byte, 1)
		fmt.Printf("i:%d, size: %d\n", i, unsafe.Sizeof(temp))

		time.Sleep(time.Second * 5)
	}

	select {}
}
