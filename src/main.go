package main

import (
	"fmt"
	"sync"
)

func main() {
	OutputLog()

	// CTF API
	host := "10.49.122.144"
	ports := rangePort()

	var wg sync.WaitGroup
	wg.Add(1)
	Request(host, ports, "/ping", &wg)
	fmt.Println("wait")
	//wg.Wait()

	rightPort := <-channel

	// Etape 1 : pong
	fmt.Println("main - len: ", Ping(host, rightPort, "/ping"))

	// Etape 2 : signup
	fmt.Println("Signup: ", Signup(host, rightPort, "/signup"))

	// Etape 3 : check
	fmt.Println("Check: ", Check(host, rightPort, "/check"))

	// Etape 4 : getUserSecret
	fmt.Println("Check: ", GetUserSecret(host, rightPort, "/getUserSecret"))

}
