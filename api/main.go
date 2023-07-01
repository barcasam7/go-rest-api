package main

func Handler() {
	app := App{}
	app.Initialise()
	app.Run("localhost:10000")
}
