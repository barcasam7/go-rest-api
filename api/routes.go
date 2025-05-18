package main

func (app *App) HandleRoutes() {
	app.Router.HandleFunc("/adverts", app.getAdverts).Methods("GET")
	app.Router.HandleFunc("/adverts/{id}", app.getAdvert).Methods("GET")
}
