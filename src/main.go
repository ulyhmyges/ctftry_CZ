package main

import "github.com/gofiber/fiber/v2/log"

func main() {
	OutputLog()
	app := AppFiber()

	// starting server
	log.Info("\nStarting server\nListen to the port 3001...")
	app.Listen(":3001")

}
