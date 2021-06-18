package main

import "time"

func main() {

	go server()
	time.Sleep(2 * time.Second)
	go client()
	select {}
}
