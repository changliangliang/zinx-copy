package main

import "time"

func main() {

	go server()
	time.Sleep(2 * time.Second)
	for i := 0; i < 1000; i++ {
		go client()
	}

	select {}
}
