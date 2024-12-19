package main

import (
	"dotdev/pocketbase"
	"log"
)

func main() {
	app := pocketbase.New()

	if err := app.Start(); err != nil {
		log.Fatal(err.Error())
	}
}
