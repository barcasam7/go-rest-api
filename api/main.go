package main

func main() {
	// Initialize the application
	app := App{}
	err := app.Initialise()
	if err != nil {
		panic(err)
	}

	// Start the server
	app.Run("localhost:10000")
}
