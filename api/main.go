package main

import "net/http"

func Handler(w http.ResponseWriter, r *http.Request) {
	app := App{}
	app.Initialise()
	app.Run("localhost:10000")
}
