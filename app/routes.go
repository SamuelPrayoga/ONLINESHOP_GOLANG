package app

import "github.com/SamuelPrayoga/gotoko/app/controllers"

func (server *Server) initializeRoutes() {
	server.Router.HandleFunc("/", controllers.Home).Methods("GET")
}
