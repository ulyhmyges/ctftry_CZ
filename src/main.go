package main

import (
	"fmt"
	"net/http"
	"sync"

	"github.com/gofiber/fiber/v2/log"
)

func main() {
	OutputLog()

	// CTF API
	host := "10.49.122.144"
	ports := rangePort()
	client := &http.Client{}

	var wg sync.WaitGroup
	wg.Add(1)
	Request(client, host, ports, "/ping", &wg)

	if len(channel) == 0 {
		wg.Done()
		log.Error("Aucun port ouvert, connexion non établie.")

	}

	wg.Wait()

	// Channel : finding the port
	rightPort := <-channel
	// Etape 1 : pong
	fmt.Println("main - len: ", Ping(client, host, rightPort, "/ping"))

	// Etape 2 : signup
	fmt.Println("Signup: ", Signup(client, host, rightPort, "/signup"))

	// Etape 3 : check
	fmt.Println("Check: ", Check(client, host, rightPort, "/check"))

	// Etape 4 : getUserSecret
	fmt.Println("Check: ", GetUserSecret(client, host, rightPort, "/getUserSecret"))

}
