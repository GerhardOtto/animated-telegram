package main

import "log"

func main() {
	app, err := NewApp()
	if err != nil {
		log.Fatalf("failed to initialise app: %v", err)
	}
	defer app.Close()

	if err := app.Run(); err != nil {
		log.Fatalf("app error: %v", err)
	}
}
