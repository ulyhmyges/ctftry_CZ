package main

import (
	"fmt"
	"sync"
)

func main() {
	OutputLog()

	//app := AppFiber()
	//var wc sync.WaitGroup
	//wc.Add(1)
	//go app.Listen(":3001")
	// starting server
	//log.Info("\nStarting server\nListen to the port 3001...")

	host := "10.49.122.144"
	ports := rangePort()
	//p := []string{"3001", "49789"}
	path := "/ping"
	//raw_connect(host, ports, path)

	var wg sync.WaitGroup
	wg.Add(1)
	Request(host, ports, path, &wg)
	wg.Wait()

	// Etape 1 : pong
	rightPort := <-channel
	fmt.Println("main - len: ", Ping(host, rightPort, path))
	path = "/signup"
	fmt.Println("Signup: ", Signup(host, rightPort, path))

}
